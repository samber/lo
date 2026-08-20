#!/usr/bin/env node
const path = require('path');
const { loadHelpers, toFullKey } = require('./utils');

const dataDir = process.argv[2] || path.join(__dirname, '..', 'data');
const { helpers, byFullKey } = loadHelpers(dataDir);

let hasError = false;

helpers.forEach((h) => {
  const thisKey = `${h.category}#${h.subCategory}#${h.slug}`.toLowerCase();
  (h.similarHelpers || []).forEach((ref) => {
    const other = byFullKey.get(toFullKey(ref));
    if (!other) return; // Existence is checked by another script
    const otherKeys = (other.similarHelpers || []).map(toFullKey);
    const otherHasBackRef = otherKeys.includes(thisKey);
    if (!otherHasBackRef) {
      hasError = true;
      console.error(`Cross-ref missing: ${h.fileName} -> ${ref} but not reciprocated.`);
    }
  });
});

if (hasError) process.exit(1);
console.log('OK: all similarHelpers are reciprocal.');
