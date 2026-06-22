function getGenericTypeNamesFromText(text: string): string[] {
  // Extract generic type names from the declaration brackets
  const bracketMatch = text.match(/\[([^\]]+)\]/);
  if (!bracketMatch) return [];
  const content = bracketMatch[1];
  
  // Split by commas and extract type names (before constraints like "any", "~[]T", etc.)
  const names = content
    .split(',')
    .map(s => s.trim())
    .map(s => (s.match(/^([A-Z][A-Za-z0-9_]*)\b/) || [])[1])
    .filter(Boolean) as string[];

  // Also look for any other capitalized identifiers that might be generic types
  // This catches cases where generics are used as types throughout the signature
  const allCapsIdentifiers = text.match(/\b([A-Z][A-Za-z0-9_]*)\b/g) || [];
  const additionalNames = allCapsIdentifiers.filter(name => 
    !['func', 'any', 'bool', 'int', 'string', 'float', 'byte', 'rune'].includes(name) &&
    !name.match(/^[A-Z][a-z]+$/) // Exclude function names
  );
  
  return Array.from(new Set([...names, ...additionalNames]));
}

export function highlightPrototypeGenerics(container: HTMLElement | null): boolean {
  if (!container) return false;
  const codeEl = container.querySelector('.helper-card__prototype-code code');
  if (!codeEl) return false;

  // Avoid double-processing
  if ((codeEl as HTMLElement).dataset.__genericsEnhanced === '1') return true;

  const textContent = codeEl.textContent || '';
  const genericNames = getGenericTypeNamesFromText(textContent);
  
  // Built-in types to color like other types
  const builtinTypes = [
    'bool','string','byte','rune','error',
    'int','int8','int16','int32','int64',
    'uint','uint8','uint16','uint32','uint64','uintptr',
    'float32','float64','complex64','complex128'
  ];
  const names = Array.from(new Set([...genericNames, ...builtinTypes]));
  if (names.length === 0) return false;

  // Simple approach: replace the entire content with highlighted version
  let newHtml = textContent;
  const pattern = new RegExp(`\\b(${names.join('|')})\\b`, 'g');
  newHtml = newHtml.replace(pattern, '<span class="token type type-parameter">$1</span>');

  if (newHtml !== textContent) {
    codeEl.innerHTML = newHtml;
    (codeEl as HTMLElement).dataset.__genericsEnhanced = '1';
    return true;
  }

  return false;
}

export function ensurePrototypeGenericsHighlighted(container: HTMLElement | null) {
  if (!container) return;

  // Try immediately
  if (highlightPrototypeGenerics(container)) return;

  // Try on next animation frame
  requestAnimationFrame(() => {
    if (highlightPrototypeGenerics(container)) return;
    
    // Try with a small delay
    setTimeout(() => highlightPrototypeGenerics(container), 50);
  });
}