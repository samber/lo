#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

// Read all markdown files in docs/data directory
const dataDir = path.join(__dirname, '../data');
const files = fs.readdirSync(dataDir).filter(f => f.endsWith('.md'));

const combinations = new Set();
const coreCategories = new Set();
const parallelCategories = new Set();
const itCategories = new Set();
const mutableCategories = new Set();

// Extract type+category combinations from each file
files.forEach(file => {
  const filePath = path.join(dataDir, file);
  const content = fs.readFileSync(filePath, 'utf8');

  const typeMatch = content.match(/^category:\s*(.+)$/m);
  const categoryMatch = content.match(/^subCategory:\s*(.+)$/m);

  if (typeMatch && categoryMatch) {
    const type = typeMatch[1].trim();
    const category = categoryMatch[1].trim();
    const combination = `${type}|${category}`;

    combinations.add(combination);

    if (type === 'core') {
      coreCategories.add(category);
    } else if (type === 'it') {
      itCategories.add(category);
    } else if (type === 'mutable') {
      mutableCategories.add(category);
    } else if (type === 'parallel') {
      parallelCategories.add(category);
    } else {
      throw new Error(`Unknown type: ${type}`);
    }
  }
});

console.log('=== TYPE+CATEGORY COMBINATIONS FOUND ===');
Array.from(combinations).sort().forEach(comb => console.log(comb));
console.log('\n=== CORE CATEGORIES ===');
Array.from(coreCategories).sort().forEach(cat => console.log(cat));
console.log('\n=== IT CATEGORIES ===');
Array.from(itCategories).sort().forEach(cat => console.log(cat));
console.log('\n=== MUTABLE CATEGORIES ===');
Array.from(mutableCategories).sort().forEach(cat => console.log(cat));
console.log('\n=== PARALLEL CATEGORIES ===');
Array.from(parallelCategories).sort().forEach(cat => console.log(cat));

// Check existing pages
const corePagesDir = path.join(__dirname, '../docs/core');
const itPagesDir = path.join(__dirname, '../docs/it');
const mutablePagesDir = path.join(__dirname, '../docs/mutable');
const parallelPagesDir = path.join(__dirname, '../docs/parallel');

const existingCorePages = new Set();
const existingItPages = new Set();
const existingMutablePages = new Set();
const existingParallelPages = new Set();

if (fs.existsSync(corePagesDir)) {
  fs.readdirSync(corePagesDir)
    .filter(f => f.endsWith('.md'))
    .forEach(f => existingCorePages.add(f.replace('.md', '')));
}
if (fs.existsSync(itPagesDir)) {
  fs.readdirSync(itPagesDir)
    .filter(f => f.endsWith('.md'))
    .forEach(f => existingItPages.add(f.replace('.md', '')));
}
if (fs.existsSync(mutablePagesDir)) {
  fs.readdirSync(mutablePagesDir)
    .filter(f => f.endsWith('.md'))
    .forEach(f => existingMutablePages.add(f.replace('.md', '')));
}
if (fs.existsSync(parallelPagesDir)) {
  fs.readdirSync(parallelPagesDir)
    .filter(f => f.endsWith('.md'))
    .forEach(f => existingParallelPages.add(f.replace('.md', '')));
}

console.log('\n=== EXISTING CORE PAGES ===');
Array.from(existingCorePages).sort().forEach(page => console.log(page));
console.log('\n=== EXISTING IT PAGES ===');
Array.from(existingItPages).sort().forEach(page => console.log(page));
console.log('\n=== EXISTING MUTABLE PAGES ===');
Array.from(existingMutablePages).sort().forEach(page => console.log(page));
console.log('\n=== EXISTING PARALLEL PAGES ===');
Array.from(existingParallelPages).sort().forEach(page => console.log(page));

// Find missing pages
console.log('\n=== MISSING CORE PAGES ===');
Array.from(coreCategories).sort().forEach(category => {
  if (!existingCorePages.has(category)) {
    console.log(`MISSING: core/${category}.md`);
  }
});
console.log('\n=== MISSING IT PAGES ===');
Array.from(itCategories).sort().forEach(category => {
  if (!existingItPages.has(category)) {
    console.log(`MISSING: it/${category}.md`);
  }
});
console.log('\n=== MISSING MUTABLE PAGES ===');
Array.from(mutableCategories).sort().forEach(category => {
  if (!existingMutablePages.has(category)) {
    console.log(`MISSING: mutable/${category}.md`);
  }
});
console.log('\n=== MISSING PARALLEL PAGES ===');
Array.from(parallelCategories).sort().forEach(category => {
  if (!existingParallelPages.has(category)) {
    console.log(`MISSING: parallel/${category}.md`);
  }
});

// Check for duplicates
console.log('\n=== VALIDATION RESULTS ===');
let hasErrors = false;

Array.from(coreCategories).sort().forEach(category => {
  if (!existingCorePages.has(category)) {
    console.log(`❌ ERROR: Missing core page for category: ${category}`);
    hasErrors = true;
  }
});
Array.from(itCategories).sort().forEach(category => {
  if (!existingItPages.has(category)) {
    console.log(`❌ ERROR: Missing it page for category: ${category}`);
    hasErrors = true;
  }
});
Array.from(mutableCategories).sort().forEach(category => {
  if (!existingMutablePages.has(category)) {
    console.log(`❌ ERROR: Missing mutable page for category: ${category}`);
    hasErrors = true;
  }
});
Array.from(parallelCategories).sort().forEach(category => {
  if (!existingParallelPages.has(category)) {
    console.log(`❌ ERROR: Missing parallel page for category: ${category}`);
    hasErrors = true;
  }
});

if (!hasErrors) {
  console.log('✅ All helper categories have corresponding pages!');
} else {
  console.log('\n❌ Found missing pages. Please create them as shown above.');
  process.exit(1);
}