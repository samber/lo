import React from 'react';
import {usePluginData} from '@docusaurus/useGlobalData';
import type {HelperDefinition} from '../index';
import './helper-components.css';

interface HelperTOCProps {
  category: string;
  subCategory: string;
  title?: string;
}

export default function HelperTOC({
  category,
  subCategory,
  title = 'On this page',
}: HelperTOCProps) {
  const data = usePluginData('helpers-pages') as {helpers: HelperDefinition[]};
  const helpers = (data?.helpers ?? [])
    .filter((h) => h.category === category && h.subCategory === subCategory)
    .slice()
    .sort((a, b) => a.position - b.position);

  if (helpers.length === 0) {
    return null;
  }

  return (
    <nav className="helper-toc" aria-label="Page navigation">
      <div className="helper-toc__title">{title}</div>
      <ul className="helper-toc__list">
        {helpers.map((h) => (
          <li key={h.slug} className="helper-toc__item">
            <a className="helper-toc__link" href={`#${h.slug}`}>{h.name}</a>
          </li>
        ))}
      </ul>
    </nav>
  );
}


