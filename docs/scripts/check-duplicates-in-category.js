#!/usr/bin/env node
const path = require('path');
const { loadHelpers } = require('./utils');

const dataDir = process.argv[2] || path.join(__dirname, '..', 'data');
const { byCategoryName } = loadHelpers(dataDir);

let hasError = false;
for (const [key, list] of byCategoryName.entries()) {
  if (list.length > 1) {
    hasError = true;
    const files = list.map((h) => h.fileName).join(', ');
    console.error(`Duplicate helper in category detected for ${key}: ${files}`);
  }
}

if (hasError) process.exit(1);
console.log('OK: no duplicate helpers within categories.');

