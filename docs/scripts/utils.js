const fs = require('fs');
const path = require('path');
const matter = require('gray-matter');

function readFile(filePath) {
  return fs.readFileSync(filePath, 'utf8');
}

function listMarkdownFiles(dirPath) {
  return fs
    .readdirSync(dirPath)
    .filter((f) => f.endsWith('.md'))
    .map((f) => path.join(dirPath, f));
}

// NOTE: this used to be a hand-rolled regex parser that only understood
// single-line arrays (`key: [a, b]`). Every helper doc uses multi-line YAML
// lists for similarHelpers/variantHelpers (`key:\n  - a\n  - b`), which the
// old parser silently dropped to an empty string on every file. As a result
// every check-* script downstream of loadHelpers() always saw empty
// similarHelpers/variantHelpers arrays and could never report a broken
// cross-reference. Using gray-matter (already a transitive dependency, and
// what the Docusaurus plugin itself uses) parses the real YAML correctly.
function parseFrontmatter(content) {
  const {data} = matter(content);
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
      variantHelpers: Array.isArray(fm.variantHelpers) ? fm.variantHelpers : [],
    };
    helpers.push(helper);
  });

  // Build index by keys for quick lookup
  const byCategoryName = new Map(); // key: `${category}#${name}` -> helper[]
  // Cross-reference labels (similarHelpers/variantHelpers) are written as
  // lowercase `category#subCategory#slug` (see docs/CLAUDE.md), not the
  // PascalCase `name` field — index on slug, lowercased, to match reality.
  const byFullKey = new Map(); // key: `${category}#${subCategory}#${slug}` (lowercased) -> helper
  const byPath = new Map(); // key: filename -> helper
  const bySlugDuplicates = new Map(); // key: `${category}#${subCategory}#${slug}` (lowercased) -> helper[]

  helpers.forEach((h) => {
    if (h.category && h.name) {
      const k = `${h.category}#${h.name}`;
      if (!byCategoryName.has(k)) byCategoryName.set(k, []);
      byCategoryName.get(k).push(h);
    }
    if (h.category && h.subCategory && h.slug) {
      const k2 = `${h.category}#${h.subCategory}#${h.slug}`.toLowerCase();
      byFullKey.set(k2, h);
      if (!bySlugDuplicates.has(k2)) bySlugDuplicates.set(k2, []);
      bySlugDuplicates.get(k2).push(h);
    }
    byPath.set(h.fileName, h);
  });

  return { helpers, byCategoryName, byFullKey, byPath, bySlugDuplicates };
}

// The `category` frontmatter value is "iter" (matches docs/docs/iter/ and
// the `iter#...` cross-reference prefix), but by long-standing filename
// convention every iterator helper file is prefixed `it-`, not `iter-`
// (see docs/CLAUDE.md: "Filename matches slug (with core- or mutable- or
// parallel- or it- prefix)"). Mapping straight from `category` produced
// 165+ false-positive "filename mismatch" errors on every run, which meant
// this check could never usefully run and real mismatches went unnoticed.
const CATEGORY_TO_FILE_PREFIX = {
  iter: 'it',
  experimental: 'simd',
};

function expectedFileName(helper) {
  if (!helper || !helper.category || !helper.slug) return null;
  const prefix = CATEGORY_TO_FILE_PREFIX[helper.category] || helper.category;
  return `${prefix}-${helper.slug}.md`;
}

function toFullKey(ref) {
  // ref format: category#subcategory#slug (lowercase)
  return ref.trim().toLowerCase();
}

module.exports = {
  readFile,
  listMarkdownFiles,
  parseFrontmatter,
  loadHelpers,
  expectedFileName,
  toFullKey,
};


