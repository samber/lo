import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

// This runs in Node.js - Don't use client-side code here (browser APIs, JSX...)

const config: Config = {
  title: 'lo',
  tagline: 'A reusable utility library for Go: slices, maps, and more',
  favicon: 'img/favicon.ico',

  // Future flags, see https://docusaurus.io/docs/api/docusaurus-config#future
  future: {
    v4: {
      removeLegacyPostBuildHeadAttribute: true,
      useCssCascadeLayers: true,
    },
    experimental_faster: {
      swcJsLoader: true,
      swcJsMinimizer: true,
      swcHtmlMinimizer: true,
      lightningCssMinimizer: true,
      rspackBundler: true,
      rspackPersistentCache: true,
      ssgWorkerThreads: true,
      mdxCrossCompilerCache: true,
    },
    experimental_storage: {
      type: 'localStorage',
      namespace: true,
    },
  },

  // Set the production url of your site here
  url: 'https://ro.samber.dev',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'samber', // Usually your GitHub org/user name.
  projectName: 'lo', // Usually your repo name.

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'throw',
  onBrokenAnchors: 'throw',

  markdown: {
    anchors: {
      maintainCase: true,
    },
    mermaid: true,
  },

  // Storage configuration for better performance
  staticDirectories: ['static'],

  // Optional: Enable hash router for offline support (experimental)
  // Uncomment if you need offline browsing capability
  // router: 'hash',
  
  // Future-proofing configurations
  clientModules: [
    require.resolve('./src/theme/prism-include-languages.js'),
  ],

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  headTags: [
    // DNS prefetch for better performance
    {
      tagName: 'link',
      attributes: {
        rel: 'dns-prefetch',
        href: '//fonts.googleapis.com',
      },
    },
    {
      tagName: 'link',
      attributes: {
        rel: 'preconnect',
        href: 'https://fonts.gstatic.com',
        crossorigin: 'anonymous',
      },
    },
    {
      tagName: 'meta',
      attributes: {
        name: 'keywords',
        content: 'go, golang, lo, slices, maps, strings, channels, functions, helpers, concurrency, error handling, resusable, utility, framework, library, samber',
      },
    },
    {
      tagName: 'meta',
      attributes: {
        property: 'og:image',
        content: 'https://lo.samber.dev/img/cover.png',
      },
    },
    {
      tagName: 'meta',
      attributes: {
        name: 'twitter:card',
        content: 'summary_large_image',
      },
    },
    {
      tagName: 'meta',
      attributes: {
        name: 'twitter:image',
        content: 'https://lo.samber.dev/img/cover.png',
      },
    },
    {
      tagName: 'meta',
      attributes: {
        name: 'twitter:creator',
        content: '@samuelberthe',
      },
    },
    {
      tagName: 'link',
      attributes: {
        rel: 'canonical',
        href: 'https://lo.samber.dev',
      },
    },
  ],

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
          'https://github.com/samber/lo/tree/master/docs/',
          showLastUpdateAuthor: true,
          showLastUpdateTime: true,
          // Enhanced docs features from 3.8+
          breadcrumbs: true,
          sidebarCollapsed: false,
          numberPrefixParser: false,
          // Enable admonitions
          admonitions: {
            keywords: ['note', 'tip', 'info', 'danger', 'warning'],
            extendDefaults: true,
          },
          // Enhanced markdown features
          remarkPlugins: [],
          rehypePlugins: [],
        },      
        sitemap: {
          lastmod: 'date',
          changefreq: 'weekly',
          priority: 0.7,
          ignorePatterns: ['/tags/**'],
          filename: 'sitemap.xml',
          // Enhanced sitemap features from 3.8+
          createSitemapItems: async (params) => {
            const {defaultCreateSitemapItems, ...rest} = params ;
            const items = await defaultCreateSitemapItems(rest);
            // Add custom priority for specific pages
            return items.map((item) => {
              if (item.url.includes('/docs/getting-started')) {
                return {...item, priority: 1.0};
              }
              if (item.url.includes('/docs/')) {
                return {...item, priority: 0.8};
              }
              return item;
            });
          },
        },
        theme: {
          customCss: './src/css/custom.css',
        },
        gtag: {
          trackingID: 'G-VVXXV8747F',
          anonymizeIP: false,
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    // Replace with your project's social card
    image: 'img/cover.png',
    colorMode: {
      defaultMode: 'light',
      disableSwitch: false,
      respectPrefersColorScheme: true,
    },

    // Mermaid configuration
    mermaid: {
      theme: {light: 'neutral', dark: 'dark'},
      options: {
        maxTextSize: 50000,
      },
    },
    
    // Enhanced metadata
    metadata: [
      {name: 'og:type', content: 'website'},
    ],

    navbar: {
      title: '🏎️ samber/lo',
      logo: {
        alt: 'lo - Reusable utility library for Go',
        src: 'img/icon.png',
      },
      items: [
        {
          type: 'docSidebar',
          sidebarId: 'docSidebar',
          position: 'left',
          label: 'Doc',
        }, 
        {
          to: 'https://pkg.go.dev/github.com/samber/lo',
          label: 'GoDoc',
          position: 'left',
        },
        {
          to: 'community',
          label: 'Community',
          position: 'left',
        },
        {
          to: 'https://github.com/samber/lo/releases',
          label: 'Changelog',
          position: 'right',
        },
        {
          to: 'https://github.com/sponsors/samber',
          label: '💖 Sponsor',
          position: 'right',
        },
        {
          href: 'https://github.com/samber/lo',
          label: 'GitHub',
          position: 'right',
          className: 'header-github-link',
          'aria-label': 'GitHub repository',
        },
        {
          type: 'search',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Project',
          items: [
            {
              label: 'Documentation',
              to: '/docs/getting-started',
            },
            {
              label: 'Changelog',
              to: 'https://github.com/samber/ro/releases',
            },
            {
              label: 'Godoc',
              to: 'https://pkg.go.dev/github.com/samber/ro',
            },
            {
              label: 'License',
              to: 'https://github.com/samber/ro/blob/master/LICENSE',
            },
            {
              label: '💖 Sponsor',
              to: 'https://github.com/sponsors/samber',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'New issue',
              to: 'https://github.com/samber/lo/issues',
            },
            {
              label: 'GitHub',
              to: 'https://github.com/samber/lo',
            },
            {
              label: 'Stack Overflow',
              to: 'https://stackoverflow.com/search?q=samber+lo',
            },
            {
              label: 'Twitter',
              to: 'https://twitter.com/samuelberthe',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} lo.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      defaultLanguage: 'go',
      additionalLanguages: ['bash', 'diff', 'json', 'yaml', 'go'],
      magicComments: [
        {
          className: 'theme-code-block-highlighted-line',
          line: 'highlight-next-line',
          block: {start: 'highlight-start', end: 'highlight-end'},
        },
        {
          className: 'code-block-error-line',
          line: 'error-next-line',
          block: {start: 'error-start', end: 'error-end'},
        },
      ],
    },
    algolia: {
      appId: 'XKJFLJJWS2',
      // bearer:disable javascript_lang_hardcoded_secret
      apiKey: 'a1c30d9a943ef77d4cd26ac3aca68be7',
      externalUrlRegex: 'lo\\.samber\\.dev',
      indexName: 'lo.samber.dev',
      contextualSearch: true,
      searchParameters: {
        // facetFilters: ['type:lvl1'],
      },
      searchPagePath: 'search',
      // Enhanced search features from 3.8+
      insights: true,
    },
  } satisfies Preset.ThemeConfig,

  themes: ['@docusaurus/theme-mermaid'],
  
  plugins: [
    // Add ideal image plugin for better image optimization
    [
      '@docusaurus/plugin-ideal-image',
      {
        quality: 70,
        max: 1030,
        min: 640,
        steps: 2,
        disableInDev: false,
      },
    ],
    [
      'vercel-analytics',
      {
        debug: true,
        mode: 'auto',
      },
    ],
    // Custom plugin to generate helper category pages from data markdown
    [
      require.resolve('./plugins/helpers-pages'),
      {},
    ],
  ],
};

export default config;
