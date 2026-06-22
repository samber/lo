#!/usr/bin/env node
const path = require('path');
const { loadHelpers, toFullKey } = require('./utils');

const dataDir = process.argv[2] || path.join(__dirname, '..', 'data');
const { helpers, byFullKey } = loadHelpers(dataDir);

let hasError = false;
helpers.forEach((h) => {
  (h.similarHelpers || []).forEach((ref) => {
    const key = toFullKey(ref);
    if (!byFullKey.has(key)) {
      hasError = true;
      console.error(`Missing similar helper reference from ${h.fileName} -> ${key}`);
    }
  });
});

if (hasError) process.exit(1);
console.log('OK: all similarHelpers references point to existing helpers.');

