import React from 'react';
import HelperCard from './HelperCard';
import {usePluginData} from '@docusaurus/useGlobalData';
import type {HelperDefinition} from '../index';
import HelperTOC from './HelperTOC';
import './helper-components.css';

interface HelperListProps {
  category: string;
  subCategory: string;
}

export default function HelperList({
  category,
  subCategory,
}: HelperListProps) {
  const gridStyle = {
    marginTop: '2rem'
  };

  const data = usePluginData('helpers-pages') as {helpers: HelperDefinition[]};
  const helpers = (data?.helpers ?? [])
    .filter((h) => h.category === category && h.subCategory === subCategory)
    .slice()
    .sort((a, b) => a.position - b.position);

  return (
    <div className="row">
      <div className="col col--10">
        <ul className="helper-list" style={gridStyle}>
          {helpers.map((helper) => (
            <li key={helper.slug}>
              <HelperCard
                helper={helper}
              />
            </li>
          ))}
        </ul>
      </div>
      <div className="col col--2">
        <HelperTOC category={category} subCategory={subCategory} />
      </div>
    </div>
  );
}
