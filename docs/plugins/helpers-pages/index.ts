import fs from 'fs';
import path from 'path';
import matter from 'gray-matter';
import type {LoadContext, Plugin} from '@docusaurus/types';

type HelperFrontMatter = {
  name: string;
  slug: string;
  sourceRef: string;
  category: 'core' | 'mutable' | 'parallel';
  subCategory: 'slice'
    | 'map'
    | 'channel'
    | 'string'
    | 'function'
    | 'find'
    | 'condition'
    | 'intersect'
    | 'type'
    | 'tuple'
    | 'math'
    | 'retry'
    | 'error-handling'
    | 'concurrency'
    | 'time';
  signatures: string[];
  playUrl?: string;
  variantHelpers: string[];
  similarHelpers: string[];
  position?: number;
};

export type HelperDefinition = HelperFrontMatter & {
  content: string;
  filePath: string;
};

function readAllHelperMarkdownFiles(dataDir: string): HelperDefinition[] {
  if (!fs.existsSync(dataDir)) {
    return [];
  }
  const entries = fs.readdirSync(dataDir, {withFileTypes: true});
  const mdFiles = entries.filter((e) => e.isFile() && e.name.endsWith('.md')).map((e) => path.join(dataDir, e.name));

  const items: HelperDefinition[] = [];
  for (const filePath of mdFiles) {
    const raw = fs.readFileSync(filePath, 'utf8');
    const {data, content} = matter(raw);
    const fm = data as Partial<HelperFrontMatter>;
    if (!fm.name || !fm.slug || !fm.category || !fm.subCategory) {
      continue;
    }
    items.push({
      name: fm.name,
      slug: fm.slug,
      sourceRef: fm.sourceRef,
      category: fm.category,
      subCategory: fm.subCategory,
      signatures: fm.signatures,
      playUrl: fm.playUrl,
      variantHelpers: Array.isArray(fm.variantHelpers) ? (fm.variantHelpers as string[]) : [],
      similarHelpers: Array.isArray(fm.similarHelpers) ? (fm.similarHelpers as string[]) : [],
      position: typeof fm.position === 'number' ? fm.position : 9999,
      content,
      filePath,
    } as HelperDefinition);
  }
  // stable sort within category/subcategory
  items.sort((a, b) => {
    if (a.category !== b.category) return a.category.localeCompare(b.category);
    if (a.subCategory !== b.subCategory) return a.subCategory.localeCompare(b.subCategory);
    if ((a.position ?? 9999) !== (b.position ?? 9999)) return (a.position ?? 9999) - (b.position ?? 9999);
    return a.name.localeCompare(b.name);
  });
  return items;
}

export default function pluginHelpersPages(context: LoadContext): Plugin<void> {
  let helpersCache: HelperDefinition[] = [];
  return {
    name: 'helpers-pages',

    async loadContent() {
      const dataDir = path.resolve(context.siteDir, 'data');
      helpersCache = readAllHelperMarkdownFiles(dataDir);
    },

    async contentLoaded({content, actions}) {
      const {setGlobalData} = actions as any;
      setGlobalData({helpers: helpersCache});
    },
  };
}

  