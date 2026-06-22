#!/usr/bin/env node
const path = require('path');
const { loadHelpers, expectedFileName } = require('./utils');

const dataDir = process.argv[2] || path.join(__dirname, '..', 'data');
const { helpers } = loadHelpers(dataDir);

let hasError = false;
helpers.forEach((h) => {
  const expected = expectedFileName(h);
  if (!expected) {
    hasError = true;
    console.error(`Invalid or missing frontmatter (category/slug) in ${h.fileName}`);
    return;
  }
  if (h.fileName !== expected) {
    hasError = true;
    console.error(`Filename mismatch for ${h.fileName}, expected ${expected}`);
  }
});

if (hasError) process.exit(1);
console.log('OK: all filenames match category and slug.');

