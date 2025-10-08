import React, { useMemo, useCallback, useState, useEffect, useRef } from 'react';
import { ensurePrototypeGenericsHighlighted } from './highlightPrototypeGenerics';
import { marked } from 'marked';
import type { HelperDefinition } from '../index';
import Heading from '@theme/Heading';
import CodeBlock from '@theme/CodeBlock';
import '../../../src/prism-include-languages.js';

interface HelperCardProps {
  helper: HelperDefinition;
}

export default function HelperCard({ 
  helper, 
}: HelperCardProps) {
  // Extract function name from signature for godoc link
  const functionName = useMemo(() => {
    if (!helper.signatures) return '';
    // Extract function name from signature like "func Map[T any, R any](collection []T, iteratee func(T, int) R) []R"
    const match = helper.signatures.find(signature => signature.match(/func\s+(\w+)/))?.match(/func\s+(\w+)/);
    return match ? match[1] : '';
  }, [helper.signatures]);

  const godocUrl = useMemo(() => {
    let baseUrl = 'https://pkg.go.dev/github.com/samber/lo';
    switch (helper.category) {
      case 'core':
        baseUrl = 'https://pkg.go.dev/github.com/samber/lo';
        break;
      case 'mutable':
        baseUrl = 'https://pkg.go.dev/github.com/samber/lo/mutable';
        break;
      case 'parallel':
        baseUrl = 'https://pkg.go.dev/github.com/samber/lo/parallel';
        break;
      // case 'it':
      //   baseUrl = 'https://pkg.go.dev/github.com/samber/lo/it';
      //   break;
    }

    if (!functionName) return '';
    return `${baseUrl}#${functionName}`;
  }, [functionName]);

  const sourceRef = useMemo(() => {
    return `https://github.com/samber/lo/blob/master/${helper.sourceRef}`;
  }, [helper.sourceRef]);

  const renderedNodes = useMemo(() => {
    marked.setOptions({
      gfm: true,
      breaks: true,
    });
    const tokens = marked.lexer(helper.content || '');

    const elements: React.ReactNode[] = [];
    let pendingTokens: any[] = [];

    const flushPending = () => {
      if (pendingTokens.length === 0) return;
      const html = (marked as any).parser(pendingTokens);
      elements.push(
        React.createElement('div', {
          className: 'helper-card__markdown-chunk',
          dangerouslySetInnerHTML: { __html: html },
          key: `md-${elements.length}`,
        })
      );
      pendingTokens = [];
    };

    for (const token of tokens as any[]) {
      if (token.type === 'code') {
        flushPending();
        const lang = (token.lang || '').trim() || undefined;
        elements.push(
          React.createElement(CodeBlock as any, {
            language: lang,
            key: `code-${elements.length}`,
            children: token.text,
          })
        );
      } else {
        pendingTokens.push(token);
      }
    }
    flushPending();

    return elements;
  }, [helper.content]);

  const signatures = useMemo(() => {
    return helper.signatures.join('\n');
  }, [helper.signatures]);

  // Post-process prototype code to colorize generic type parameters locally
  const prototypeRowRef = useRef<HTMLDivElement>(null);
  useEffect(() => {
    // // Ensure highlighting even if Prism updates asynchronously
    // const raf = requestAnimationFrame(() => {
    //   ensurePrototypeGenericsHighlighted(prototypeRowRef.current as unknown as HTMLElement);
    // });
    // return () => cancelAnimationFrame(raf);
  }, [helper.signatures]);

  return (
    <div className="helper-card">
      <div className="helper-card__header">
        {/* Heading registered in MDX ToC */}
        <Heading as="h3" id={helper.slug} className="helper-card__title anchor">
          {helper.name}
        </Heading>
        {/* <h3 className="helper-card__title">
          <a href={`#${helper.slug}`}>
            {helper.name}
          </a>
        </h3> */}
        <div className="helper-card__actions">
          <div className="helper-card__badges">
            <span className="helper-card__badge helper-card__badge--category">
              {helper.subCategory}
            </span>
            <span className="helper-card__badge helper-card__badge--subcategory">
              {helper.category}
            </span>
          </div>
          {sourceRef && (
            <a 
              href={sourceRef}
              target="_blank"
              rel="noopener noreferrer"
              className="helper-card__source"
            >
              🧩 Source
            </a>
          )}
          {godocUrl && (
            <a 
              href={godocUrl}
              target="_blank"
              rel="noopener noreferrer"
              className="helper-card__godoc"
            >
              📚 GoDoc
            </a>
          )}
          {helper.playUrl && (
            <a 
              href={helper.playUrl}
              target="_blank"
              rel="noopener noreferrer"
              className="helper-card__playground"
            >
              🎮 Try on Go Playground
            </a>
          )}
        </div>
      </div>
      
      <div className="helper-card__content">
        <div className="helper-card__markdown">
          {renderedNodes}
        </div>

        <SimilarHelpers
          title="Variant"
          similarHelpers={helper.variantHelpers}
          currentType={helper.category || ''}
          currentCategory={helper.subCategory || ''}
          currentName={helper.name || ''}
        />
        <SimilarHelpers
          title="Similar"
          similarHelpers={helper.similarHelpers}
          currentType={helper.category || ''}
          currentCategory={helper.subCategory || ''}
          currentName={helper.name || ''}
        />
      </div>
      
      <div className="helper-card__prototype">
        <div className="helper-card__prototype-row" ref={prototypeRowRef}>
          <span className="helper-card__prototype-label">Prototype{helper.signatures.length > 1 ? 's' : ''}:</span>
          <CodeBlock language="go" className="helper-card__prototype-code" children={signatures} />
        </div>
      </div>
    </div>
  );
}

type SimilarHelpersProps = {
  title: string;
  similarHelpers: string[];
  currentType: string;
  currentCategory: string;
  currentName: string;
};

function SimilarHelpers({ 
  title, 
  similarHelpers, 
  currentType,
  currentCategory, 
  currentName,
}: SimilarHelpersProps) {
  const currentHelperLower = `${currentType}#${currentCategory}#${currentName}`.toLowerCase();
  similarHelpers = similarHelpers.filter((helper) => helper != currentHelperLower);

  if (similarHelpers.length === 0) {
    return null;
  }

  return (
    <div className="helper-card__similar">
      <div className="helper-card__similar-row">
        <span className="helper-card__prototype-label">{title}:</span>
        <div className="helper-card__similar-list">
          {similarHelpers
            .map((originalLabel) => {
              const parts = String(originalLabel).split('#');
              const typeRaw = parts[0];
              const categoryRaw = parts[1];
              const nameRaw = parts[2];
              const fallbackType = currentType || '';
              const fallbackCategory = currentCategory || '';
              const type = (typeRaw || fallbackType).toLowerCase();
              const category = (categoryRaw || fallbackCategory).toLowerCase();
              // Fallback for legacy 2-part labels: type#name
              const legacyName = parts.length === 2 ? parts[1] : undefined;
              const name = (nameRaw || legacyName || '').toLowerCase();
              return { originalLabel, type, category, name, nameRaw };
            })
            .sort((a, b) => {
              const currentTypeLower = (currentType || '').toLowerCase();
              const aSame = a.type === currentTypeLower ? 0 : 1;
              const bSame = b.type === currentTypeLower ? 0 : 1;
              if (aSame !== bSame) return aSame - bSame; // same type first
              return a.name.localeCompare(b.name);
            })
            .map(({ originalLabel, type, category, name, nameRaw }, index) => {
              const currentTypeLower = (currentType || '').toLowerCase();
              const href = `/docs/${type}/${category}#${name}`;
              const displayName = nameRaw || name;
              const isSameSection = type === currentTypeLower; // compare only type for label
              return (
                <a 
                  key={index}
                  href={href}
                  className="helper-card__similar-link"
                >
                  {isSameSection ? (
                    displayName
                  ) : (
                    <>
                      <span className="helper-card__similar-prefix">{type}›</span>{' '}
                      {displayName}
                    </>
                  )}
                </a>
              );
            })}
        </div>
      </div>
    </div>
  );
}
