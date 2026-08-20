#!/usr/bin/env node
// Verifies that no two helpers in the same category/subCategory share a
// slug. Two helpers with the same slug render the same HTML `id` on the
// same page, making one of them permanently unreachable by anchor.
const path = require('path');
const { loadHelpers } = require('./utils');

const dataDir = process.argv[2] || path.join(__dirname, '..', 'data');
const { bySlugDuplicates } = loadHelpers(dataDir);

let hasError = false;
for (const [key, list] of bySlugDuplicates.entries()) {
  if (list.length > 1) {
    hasError = true;
    const files = list.map((h) => h.fileName).join(', ');
    console.error(`Duplicate slug detected for ${key}: ${files}`);
  }
}

if (hasError) process.exit(1);
console.log('OK: no duplicate slugs within any category/subCategory.');
