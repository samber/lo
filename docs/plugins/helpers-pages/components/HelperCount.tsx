import React from 'react';
import {usePluginData} from '@docusaurus/useGlobalData';
import type {HelperDefinition} from '../index';

interface HelperCountProps {
  category: string;
  subCategory: string;
}

export default function HelperCount({category, subCategory}: HelperCountProps) {
  const data = usePluginData('helpers-pages') as {helpers: HelperDefinition[]};
  const count = (data?.helpers ?? []).filter(
    (h) => h.category === category && h.subCategory === subCategory,
  ).length;

  return <>{count}</>;
}
