const fs = require('fs');
const path = require('path');

function readFile(filePath) {
  return fs.readFileSync(filePath, 'utf8');
}

function listMarkdownFiles(dirPath) {
  return fs
    .readdirSync(dirPath)
    .filter((f) => f.endsWith('.md'))
    .map((f) => path.join(dirPath, f));
}

function parseFrontmatter(content) {
  // Very small frontmatter parser tailored to our files
  // Expects YAML-like block between leading --- lines
  const fmMatch = content.match(/^---[\r\n]+([\s\S]*?)[\r\n]+---/);
  if (!fmMatch) return null;
  const fm = fmMatch[1];
  const data = {};

  // Capture simple key: value pairs and array values like: key: [a, b]
  fm.split(/\r?\n/).forEach((line) => {
    const m = line.match(/^([A-Za-z][A-Za-z0-9_-]*):\s*(.*)$/);
    if (!m) return;
    const key = m[1];
    const raw = m[2].trim();
    if (raw.startsWith('[') && raw.endsWith(']')) {
      const inner = raw.slice(1, -1).trim();
      if (inner.length === 0) {
        data[key] = [];
      } else {
        data[key] = inner
          .split(',')
          .map((s) => s.trim())
          .map((s) => (s.startsWith('"') && s.endsWith('"') ? s.slice(1, -1) : s));
      }
    } else if (raw.startsWith('"') && raw.endsWith('"')) {
      data[key] = raw.slice(1, -1);
    } else if (raw === '[]') {
      data[key] = [];
    } else if (raw === 'null') {
      data[key] = null;
    } else {
      data[key] = raw;
    }
  });
  return data;
}

function loadHelpers(dataDir) {
  const files = listMarkdownFiles(dataDir);
  const helpers = [];

  files.forEach((absPath) => {
    const filename = path.basename(absPath);
    const content = readFile(absPath);
    const fm = parseFrontmatter(content) || {};
    const helper = {
      filePath: absPath,
      fileName: filename,
      name: fm.name || null,
      slug: fm.slug || null,
      category: fm.category || null,
      subCategory: fm.subCategory || null,
      similarHelpers: Array.isArray(fm.similarHelpers) ? fm.similarHelpers : [],
    };
    helpers.push(helper);
  });

  // Build index by keys for quick lookup
  const byCategoryName = new Map(); // key: `${category}#${name}` -> helper
  const byFullKey = new Map(); // key: `${category}#${subCategory}#${name}` -> helper
  const byPath = new Map(); // key: filename -> helper

  helpers.forEach((h) => {
    if (h.category && h.name) {
      const k = `${h.category}#${h.name}`;
      if (!byCategoryName.has(k)) byCategoryName.set(k, []);
      byCategoryName.get(k).push(h);
    }
    if (h.category && h.subCategory && h.name) {
      const k2 = `${h.category}#${h.subCategory}#${h.name}`;
      byFullKey.set(k2, h);
    }
    byPath.set(h.fileName, h);
  });

  return { helpers, byCategoryName, byFullKey, byPath };
}

function expectedFileName(helper) {
  if (!helper || !helper.category || !helper.slug) return null;
  return `${helper.category}-${helper.slug}.md`;
}

function toFullKey(ref) {
  // ref format: category#subcategory#Name
  return ref.trim();
}

module.exports = {
  readFile,
  listMarkdownFiles,
  parseFrontmatter,
  loadHelpers,
  expectedFileName,
  toFullKey,
};


