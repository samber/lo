#!/usr/bin/env node

const fs = require('fs');
const path = require('path');
const readline = require('readline');
const { listMarkdownFiles, parseFrontmatter } = require('./utils');

const repoRoot = path.resolve(__dirname, '..', '..');
const dataDir = path.resolve(__dirname, '..', 'data');

function readFile(filePath) {
  return fs.readFileSync(filePath, 'utf8');
}

function* walkGoFiles(dir, excludeDirs = new Set()) {
  const entries = fs.readdirSync(dir, { withFileTypes: true });
  for (const entry of entries) {
    if (entry.name.startsWith('.')) continue;
    const abs = path.join(dir, entry.name);
    const rel = path.relative(repoRoot, abs);
    if (entry.isDirectory()) {
      if (excludeDirs.has(entry.name)) continue;
      yield* walkGoFiles(abs, excludeDirs);
    } else if (entry.isFile() && entry.name.endsWith('.go')) {
      // skip tests
      if (entry.name.endsWith('_test.go')) continue;
      // skip docs/ directory
      if (rel.split(path.sep)[0] === 'docs') continue;
      yield abs;
    }
  }
}

function buildFunctionRegex(name) {
  // Matches: func Name[...]( or func Name(
  const escaped = name.replace(/[-\/\\^$*+?.()|[\]{}]/g, '\\$&');
  return new RegExp('^func\\s+' + escaped + '(?:\\(|\\[)', '');
}

async function findFunctionDeclarations(name, preferredPathHint) {
  const fnRegex = buildFunctionRegex(name);
  const hits = [];

  // Prefer hinted file if provided
  if (preferredPathHint) {
    const hintedAbs = path.resolve(repoRoot, preferredPathHint);
    if (fs.existsSync(hintedAbs)) {
      const hit = await scanFileForSignature(hintedAbs, fnRegex);
      if (hit) hits.push(hit);
    }
  }

  for (const abs of walkGoFiles(repoRoot)) {
    const hit = await scanFileForSignature(abs, fnRegex);
    if (hit) hits.push(hit);
  }
  return hits;
}

function stripBOM(s) {
  return s.charCodeAt(0) === 0xfeff ? s.slice(1) : s;
}

async function scanFileForSignature(absPath, fnRegex) {
  const rl = readline.createInterface({
    input: fs.createReadStream(absPath, { encoding: 'utf8' }),
    crlfDelay: Infinity,
  });
  let lineNo = 0;
  for await (const rawLine of rl) {
    lineNo++;
    const line = stripBOM(rawLine);
    if (fnRegex.test(line)) {
      // Normalize multiple spaces and tabs minimally: keep original line
      let signature = line.trim();
      // Remove end-of-line comments before processing trailing bracket
      signature = signature
        // remove line comments
        .replace(/\/\/.*$/, '')
        // remove trailing block comment
        .replace(/\/\*.*?\*\/\s*$/, '')
        .trimEnd();

      // Remove spaces before and after trailing opening brace, then drop it
      signature = signature.replace(/\s*\{\s*$/, '');
      const rel = path.relative(repoRoot, absPath).replace(/\\/g, '/');
      return { file: rel, line: lineNo, signature };
    }
  }
  return null;
}

function parseSourceRefFile(sourceRef) {
  if (!sourceRef) return null;
  const idx = sourceRef.indexOf('#');
  if (idx === -1) return sourceRef;
  return sourceRef.slice(0, idx);
}

function normalizeSignature(signature) {
  // Collapse all whitespace and remove spaces before '(' or '['
  return signature
    .replace(/\s+/g, ' ')
    .replace(/\s+(\(|\[)/g, '$1')
    .trim();
}

async function main() {
  const args = new Set(process.argv.slice(2));
  const files = listMarkdownFiles(dataDir);
  let issues = 0;

  for (const absPath of files) {
    const content = readFile(absPath);
    const fm = parseFrontmatter(content) || {};
    const name = fm.name;
    if (!name) continue;

    const hintFile = parseSourceRefFile(fm.sourceRef);
    const hits = await findFunctionDeclarations(name, hintFile);
    const relMd = path.relative(repoRoot, absPath).replace(/\\/g, '/');

    if (!hits || hits.length === 0) {
      // eslint-disable-next-line no-console
      console.warn(`[missing-helper] ${relMd} -> name="${name}"`);
      issues++;
      continue;
    }

    // Sort hits by file path to ensure consistent order
    hits.sort((a, b) => a.file.localeCompare(b.file));

    // Deduplicate signatures from hits (preserve first encountered formatting)
    const seenFromHits = new Set();
    const uniqueHitSignatures = [];
    for (const hit of hits) {
      const norm = normalizeSignature(hit.signature);
      if (!seenFromHits.has(norm)) {
        seenFromHits.add(norm);
        uniqueHitSignatures.push(hit.signature);
      }
    }

    // Existing frontmatter signatures
    const existing = Array.isArray(fm.signatures) ? fm.signatures : [];
    const existingNorm = existing.map(normalizeSignature);

    // Report duplicate signatures within frontmatter (second and further occurrences)
    const seenExisting = new Set();
    for (let i = 0; i < existing.length; i++) {
      const norm = existingNorm[i];
      if (seenExisting.has(norm)) {
        // eslint-disable-next-line no-console
        console.warn(`[duplicate-signature] ${relMd} -> "${existing[i]}"`);
        issues++;
      } else {
        seenExisting.add(norm);
      }
    }

    // Unknown signatures (exist in frontmatter but not in code)
    const hitsNormalized = new Set(uniqueHitSignatures.map(normalizeSignature));
    for (const sig of existing) {
      const norm = normalizeSignature(sig);
      if (!hitsNormalized.has(norm)) {
        // eslint-disable-next-line no-console
        console.warn(`[unknown-signature] ${relMd} -> "${sig}"`);
        issues++;
      }
    }

    // Missing signatures (found in code but not listed in frontmatter)
    const existingNormalizedSet = new Set(existingNorm);
    for (const sig of uniqueHitSignatures) {
      const norm = normalizeSignature(sig);
      if (!existingNormalizedSet.has(norm)) {
        // eslint-disable-next-line no-console
        console.warn(`[missing-signature] ${relMd} -> "${sig}"`);
        issues++;
      }
    }

    // SourceRef verification
    const expectedSourceRef = `${hits[0].file}#L${hits[0].line}`;
    if (fm.sourceRef !== expectedSourceRef) {
      // eslint-disable-next-line no-console
      console.warn(`[sourceRef-outdated] ${relMd} -> expected=${expectedSourceRef} actual=${fm.sourceRef || '""'}`);
      issues++;
    }
  }

  if (args.has('--check') && issues > 0) {
    process.exitCode = 1;
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});


