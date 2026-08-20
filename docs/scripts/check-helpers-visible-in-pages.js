#!/usr/bin/env node
//
// Verifies that every category/subCategory pair declared in docs/data/*.md
// has a corresponding page under docs/docs/<category>/<subCategory>.md.
//
// This used to hardcode 4 categories (core/it/mutable/parallel) and threw on
// any other category value -- which meant it could never actually run
// against real data, since helpers are tagged `category: iter`, not `it`,
// and `experimental` wasn't handled at all. It silently never caught the
// 11 `iter/math` and `iter/condition` helpers that had no rendering page.
const fs = require('fs');
const path = require('path');

const dataDir = path.join(__dirname, '..', 'data');
const docsDir = path.join(__dirname, '..', 'docs');

const files = fs.readdirSync(dataDir).filter((f) => f.endsWith('.md'));

// category -> Set of subCategory
const combinationsByCategory = new Map();

files.forEach((file) => {
  const content = fs.readFileSync(path.join(dataDir, file), 'utf8');
  const categoryMatch = content.match(/^category:\s*(.+)$/m);
  const subCategoryMatch = content.match(/^subCategory:\s*(.+)$/m);
  if (!categoryMatch || !subCategoryMatch) return;

  const category = categoryMatch[1].trim();
  const subCategory = subCategoryMatch[1].trim();

  if (!combinationsByCategory.has(category)) combinationsByCategory.set(category, new Set());
  combinationsByCategory.get(category).add(subCategory);
});

console.log('=== CATEGORY/SUBCATEGORY COMBINATIONS FOUND ===');
for (const [category, subCategories] of combinationsByCategory) {
  Array.from(subCategories).sort().forEach((sc) => console.log(`${category}/${sc}`));
}

let hasErrors = false;
console.log('\n=== VALIDATION RESULTS ===');

for (const [category, subCategories] of combinationsByCategory) {
  const pagesDir = path.join(docsDir, category);
  const existingPages = fs.existsSync(pagesDir)
    ? new Set(fs.readdirSync(pagesDir).filter((f) => f.endsWith('.md')).map((f) => f.replace('.md', '')))
    : new Set();

  Array.from(subCategories).sort().forEach((subCategory) => {
    if (!existingPages.has(subCategory)) {
      hasErrors = true;
      console.log(`❌ ERROR: docs/data/*.md declares ${category}/${subCategory} but docs/docs/${category}/${subCategory}.md does not exist`);
    }
  });
}

if (!hasErrors) {
  console.log('✅ All helper category/subCategory pairs have a corresponding page!');
} else {
  console.log('\n❌ Found helper data with no rendering page. Create the missing docs/docs/<category>/<subCategory>.md file(s) above.');
  process.exit(1);
}
