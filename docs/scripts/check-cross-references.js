#!/usr/bin/env node
const path = require('path');
const { loadHelpers } = require('./utils');

const dataDir = process.argv[2] || path.join(__dirname, '..', 'data');
const { helpers, byFullKey } = loadHelpers(dataDir);

let hasError = false;

helpers.forEach((h) => {
  const thisKey = `${h.category}#${h.subCategory}#${h.name}`;
  (h.similarHelpers || []).forEach((ref) => {
    const other = byFullKey.get(ref);
    if (!other) return; // Existence is checked by another script
    const otherHasBackRef = (other.similarHelpers || []).includes(thisKey);
    if (!otherHasBackRef) {
      hasError = true;
      console.error(`Cross-ref missing: ${h.fileName} -> ${ref} but not reciprocated.`);
    }
  });
});

if (hasError) process.exit(1);
console.log('OK: all similarHelpers are reciprocal.');
