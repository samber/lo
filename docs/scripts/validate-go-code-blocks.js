#!/usr/bin/env node
//
// Extracts every ```go fenced code block from docs/data/*.md and
// docs/docs/**/*.md, wraps each into a standalone program, and compiles
// all of them in a single `go build` pass against the LOCAL source tree
// (via a go.mod replace directive) -- so a helper whose signature changed
// but whose doc example didn't gets caught before a reader hits it.
//
// This exists because getting-started.md shipped three non-compiling
// examples (wrong stdlib call, wrong signature, wrong argument count) that
// went unnoticed for a while -- nothing in CI actually compiled the Go
// embedded in the docs.
//
// Known limitations, in decreasing order of how often they actually bite:
//
// 1. Each ```go block is checked independently. A handful of doc files
//    (e.g. core-attempt.md) spread one narrative across several sequential
//    fenced blocks, reassigning (`=`) variables `:=`-declared in an earlier
//    block. Checked in isolation, the later blocks report "undefined: x".
//    This is a real signal too (the example isn't copy-paste runnable on
//    its own), just not the same kind of bug as a wrong signature.
// 2. Doc examples almost never call fmt.Println on their result (the
//    convention is "call + `// expected output` comment"), which would
//    make a naive wrap fail with "declared and not used" on every single
//    block. Worked around by blank-discarding every identifier declared
//    with `:=` at the OUTERMOST scope of the snippet (tracked via brace
//    depth) -- a snippet that declares an unused variable inside a nested
//    for/if/func-literal block can still produce a false "declared and not
//    used", but doc examples are written flat, so this is rare.
// 3. A block that is intentionally illustrative pseudocode (e.g. a function
//    signature with a literal `{ ... }` body, or a fence mislabeled ```go
//    around a shell snippet) will correctly fail to parse. That's not a
//    false positive -- it means the block should either be fixed to be
//    real Go or given a different fence language.
const fs = require('fs');
const os = require('os');
const path = require('path');
const { execFileSync } = require('child_process');

const REPO_ROOT = path.resolve(__dirname, '..', '..');
const DATA_DIR = path.join(REPO_ROOT, 'docs', 'data');
const DOCS_DIR = path.join(REPO_ROOT, 'docs', 'docs');

function listMarkdownFilesRecursive(dir) {
  const out = [];
  if (!fs.existsSync(dir)) return out;
  for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) out.push(...listMarkdownFilesRecursive(full));
    else if (entry.name.endsWith('.md') || entry.name.endsWith('.mdx')) out.push(full);
  }
  return out;
}

// Splits a markdown file into its ```go fenced blocks, along with the
// 1-indexed line of the first line of code (for error reporting).
function extractGoBlocks(filePath) {
  const lines = fs.readFileSync(filePath, 'utf8').split('\n');
  const blocks = [];
  let inBlock = false;
  let current = [];
  let startLine = 0;

  lines.forEach((line, i) => {
    if (!inBlock && /^```go\s*$/.test(line.trim())) {
      inBlock = true;
      current = [];
      startLine = i + 2;
      return;
    }
    if (inBlock && line.trim() === '```') {
      inBlock = false;
      blocks.push({ code: current.join('\n'), startLine });
      return;
    }
    if (inBlock) current.push(line);
  });

  return blocks;
}

// Identifiers declared with `:=` at brace-depth 0 relative to the snippet
// itself (i.e. not inside a nested block) -- these get blank-discarded so
// wrapping the snippet in `func main() { ... }` doesn't trip "declared and
// not used" on examples that only show their result as a trailing comment.
function collectTopLevelDeclarations(code) {
  const idents = new Set();
  let depth = 0;

  for (const rawLine of code.split('\n')) {
    const line = rawLine.trim();
    if (depth === 0) {
      const declMatch = line.match(/^([A-Za-z_][A-Za-z0-9_]*(?:\s*,\s*[A-Za-z_][A-Za-z0-9_]*)*)\s*:=/);
      if (declMatch) {
        declMatch[1].split(',').forEach((n) => {
          const name = n.trim();
          if (name && name !== '_') idents.add(name);
        });
      }
    }
    for (const ch of rawLine) {
      if (ch === '{') depth++;
      else if (ch === '}') depth--;
    }
  }

  return [...idents];
}

// Detects which samber/lo sub-package(s) a snippet needs, by the alias
// convention used throughout this repo's docs (lo., lom., lop., it.).
function detectImports(code) {
  const imports = [];
  if (/\blo\.[A-Z]/.test(code)) imports.push('"github.com/samber/lo"');
  if (/\blom\.[A-Z]/.test(code)) imports.push('lom "github.com/samber/lo/mutable"');
  if (/\blop\.[A-Z]/.test(code)) imports.push('lop "github.com/samber/lo/parallel"');
  if (/\bit\.[A-Z]/.test(code)) imports.push('"github.com/samber/lo/it"');
  return imports;
}

// Some doc examples show their own `import (...)` block to demonstrate the
// real import path -- illegal once dropped inside a function body. Pull any
// such declarations out and blank out their lines (not remove them, so
// every other line's number stays stable for error reporting).
function extractDeclaredImports(code) {
  const imports = [];
  const lines = code.split('\n');
  const result = [];
  let i = 0;

  while (i < lines.length) {
    const line = lines[i];
    const trimmed = line.trim();

    const singleLine = trimmed.match(/^import\s+((?:[\w.]+\s+)?"[^"]+")\s*$/);
    if (singleLine) {
      imports.push(singleLine[1]);
      result.push('');
      i++;
      continue;
    }

    if (/^import\s*\(\s*$/.test(trimmed)) {
      result.push('');
      i++;
      while (i < lines.length && lines[i].trim() !== ')') {
        const spec = lines[i].trim();
        if (spec) imports.push(spec);
        result.push('');
        i++;
      }
      if (i < lines.length) {
        result.push(''); // the closing ')'
        i++;
      }
      continue;
    }

    result.push(line);
    i++;
  }

  return { imports, code: result.join('\n') };
}

// Some doc examples declare a helper type (with methods) needed to satisfy
// a generic constraint, e.g. Fill's `Clonable[T]` requirement -- a type or
// func declaration is illegal inside a function body, so anything matching
// `type X struct/interface {` or `func (recv T) Method(...) {` (or a plain
// top-level `func name(...) {`) at brace-depth 0 gets hoisted above
// `func main()` instead of left inline. Lines are blanked, not removed, to
// keep line numbers stable for error reporting.
function hoistTopLevelDecls(code) {
  const lines = code.split('\n');
  const bodyLines = [];
  const hoisted = [];
  let i = 0;

  const declStart = /^(type\s+\w+\s+(struct|interface)\s*{|func\s*(\([^)]*\))?\s*\w+\s*\([^{]*\)[^{]*{)/;

  while (i < lines.length) {
    const trimmed = lines[i].trim();
    if (declStart.test(trimmed)) {
      let depth = 0;
      const declLines = [];
      do {
        declLines.push(lines[i]);
        for (const ch of lines[i]) {
          if (ch === '{') depth++;
          else if (ch === '}') depth--;
        }
        bodyLines.push('');
        i++;
      } while (i < lines.length && depth > 0);
      hoisted.push(declLines.join('\n'));
      continue;
    }
    bodyLines.push(lines[i]);
    i++;
  }

  return { hoisted, code: bodyLines.join('\n') };
}

function wrapBlock(rawCode) {
  const { imports: declaredImports, code: afterImports } = extractDeclaredImports(rawCode);
  const { hoisted, code } = hoistTopLevelDecls(afterImports);
  const imports = [...new Set([...declaredImports, ...detectImports(code), ...detectImports(hoisted.join('\n'))])];
  const discards = collectTopLevelDeclarations(code)
    .map((n) => `\t_ = ${n}`)
    .join('\n');

  const header = [
    'package main',
    '',
    imports.length ? `import (\n${imports.map((i) => '\t' + i).join('\n')}\n)\n` : '',
    ...hoisted,
    'func main() {',
  ].join('\n');

  const wrapped = [header, code, discards, '}', ''].join('\n');
  const headerLineCount = header.split('\n').length;
  return { wrapped, headerLineCount };
}

function main() {
  const files = [...listMarkdownFilesRecursive(DATA_DIR), ...listMarkdownFilesRecursive(DOCS_DIR)]
    // exp/simd helpers require go1.26+goexperiment.simd+amd64 build tags
    // (see docs/CLAUDE.md); they can't be built with a plain `go build` on
    // most machines/architectures and are already exempt from having a Go
    // Playground example for the same reason.
    .filter((f) => !path.basename(f).startsWith('simd-'));

  const blocksToCheck = [];
  files.forEach((filePath) => {
    extractGoBlocks(filePath).forEach((block, idx) => {
      if (!block.code.trim()) return;
      // Skip blocks that are pure output/type comments, not real code
      // (rare, but some docs show a bare `// map[int][]int{...}` block).
      const realCodeLines = block.code
        .split('\n')
        .filter((l) => l.trim() && !l.trim().startsWith('//'));
      if (realCodeLines.length === 0) return;
      blocksToCheck.push({ filePath, blockIndex: idx, ...block });
    });
  });

  if (blocksToCheck.length === 0) {
    console.log('No ```go code blocks found.');
    return;
  }

  const tmpRoot = fs.mkdtempSync(path.join(os.tmpdir(), 'lo-validate-go-blocks-'));
  const blockMeta = [];

  blocksToCheck.forEach((block, i) => {
    const dir = path.join(tmpRoot, `block_${i}`);
    fs.mkdirSync(dir, { recursive: true });
    const { wrapped, headerLineCount } = wrapBlock(block.code);
    fs.writeFileSync(path.join(dir, 'main.go'), wrapped, 'utf8');
    blockMeta.push({
      dirName: `block_${i}`,
      filePath: block.filePath,
      startLine: block.startLine,
      headerLineCount,
    });
  });

  if (process.env.DEBUG_BLOCK_MAP) {
    fs.writeFileSync(process.env.DEBUG_BLOCK_MAP, JSON.stringify(blockMeta, null, 2));
  }

  fs.writeFileSync(
    path.join(tmpRoot, 'go.mod'),
    [
      'module lo-validate-go-blocks',
      '',
      // 1.23+ so range-over-func / iter.Seq examples (lo/it) compile too.
      'go 1.23',
      '',
      'require github.com/samber/lo v1.53.0',
      '',
      `replace github.com/samber/lo => ${REPO_ROOT}`,
      '',
    ].join('\n'),
    'utf8'
  );

  try {
    execFileSync('go', ['mod', 'tidy'], { cwd: tmpRoot, stdio: 'pipe' });
  } catch (err) {
    console.error('go mod tidy failed:');
    console.error(err.stderr ? err.stderr.toString() : err.message);
    process.exitCode = 1;
    return;
  }

  // Run goimports per block directory, not as one batch over tmpRoot: a
  // single block that isn't valid Go to begin with (illustrative pseudocode,
  // or a fence mislabeled ```go around a shell snippet) would otherwise
  // abort import resolution for every other block in the batch. Letting
  // each one fail independently means it still gets reported by the build
  // step below, for the right reason, without collateral damage.
  let goimportsAvailable = true;
  blockMeta.forEach((meta) => {
    try {
      execFileSync('goimports', ['-w', path.join(tmpRoot, meta.dirName)], { stdio: 'pipe' });
    } catch (err) {
      if (/ENOENT/.test(err.message)) goimportsAvailable = false;
      // else: this block's own error will surface from `go build` below.
    }
  });
  if (!goimportsAvailable) {
    console.warn('Warning: goimports is not installed -- stdlib import resolution skipped for all blocks.');
    console.warn('Install it with: go install golang.org/x/tools/cmd/goimports@latest');
  }

  // goimports may have added imports for third-party packages (e.g. a doc
  // example using google/uuid) that weren't in go.mod yet -- tidy again so
  // `go build` doesn't fail on those with a "no required module" error
  // instead of a real compile error.
  try {
    execFileSync('go', ['mod', 'tidy'], { cwd: tmpRoot, stdio: 'pipe' });
  } catch {
    // Non-fatal: an unresolvable third-party import will still surface
    // clearly from `go build` below.
  }

  let buildOutput = '';
  let buildFailed = false;
  try {
    execFileSync('go', ['build', './...'], { cwd: tmpRoot, stdio: 'pipe' });
  } catch (err) {
    buildFailed = true;
    buildOutput = err.stderr ? err.stderr.toString() : err.message;
  }

  if (!buildFailed) {
    console.log(`OK: all ${blocksToCheck.length} Go code blocks compile.`);
    fs.rmSync(tmpRoot, { recursive: true, force: true });
    return;
  }

  // Map "block_42/main.go:7:2: ..." lines back to the original doc file
  // and an approximate line number in it.
  const byDir = new Map(blockMeta.map((m) => [m.dirName, m]));
  const errorLineRegex = /^(block_\d+)[\\/]main\.go:(\d+):(\d+):\s*(.+)$/;
  const reported = new Set();

  console.error(`FAILED: ${blocksToCheck.length} blocks checked, build errors below.\n`);

  buildOutput.split('\n').forEach((line) => {
    const m = line.match(errorLineRegex);
    if (!m) return;
    const [, dirName, wrappedLine, col, message] = m;
    const meta = byDir.get(dirName);
    if (!meta) return;
    const approxLine = meta.startLine + (parseInt(wrappedLine, 10) - meta.headerLineCount - 1);
    const relPath = path.relative(REPO_ROOT, meta.filePath);
    const key = `${relPath}:${approxLine}:${message}`;
    if (reported.has(key)) return;
    reported.add(key);
    console.error(`${relPath}:~${approxLine} -- ${message}`);
  });

  if (reported.size === 0) {
    // Errors we couldn't map (e.g. go mod/build-system errors) -- dump raw.
    console.error(buildOutput);
  }

  fs.rmSync(tmpRoot, { recursive: true, force: true });
  process.exitCode = 1;
}

main();
