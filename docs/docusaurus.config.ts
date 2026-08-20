import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

// This runs in Node.js - Don't use client-side code here (browser APIs, JSX...)

const config: Config = {
  title: 'samber/lo - Go Generics Utility Library',
  tagline: 'Type-safe utility library for Go using generics: slices, maps, and more',
  favicon: 'img/favicon.ico',

  // Future flags, see https://docusaurus.io/docs/api/docusaurus-config#future
  future: {
    v4: {
      removeLegacyPostBuildHeadAttribute: true,
      useCssCascadeLayers: true,
    },
    faster: {
        swcJsLoader: true,
        swcJsMinimizer: true,
        swcHtmlMinimizer: true,
        lightningCssMinimizer: true,
        rspackBundler: true,
        rspackPersistentCache: true,
        ssgWorkerThreads: true,
        mdxCrossCompilerCache: true,
    },
  },
    storage: {
        type: 'localStorage',
        namespace: true,
    },

    // Set the production url of your site here
  url: 'https://lo.samber.dev',
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
    require.resolve('./src/clientModules/posthog-events.ts'),
  ],

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

    headTags: [
	      // SEO
    {
      tagName: 'meta',
      attributes: {
        name: 'msvalidate.01',
        content: '4576E3F85783A82149A0DB35A150F7EB',
      },
    },
    {
      tagName: 'link',
      attributes: {
        rel: 'alternate',
        type: 'text/plain',
        href: '/llms.txt',
        title: 'llms.txt',
      },
    },
        {
        tagName: 'script',
        attributes: {
            'async': 'true',
            'src': 'https://analytics.ahrefs.com/analytics.js',
            'data-key': 'ZlVVDleFCGZPB8Nd2KkKrw'
        },
    },
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
        content: 'go, golang, lo, slices, maps, strings, channels, functions, helpers, concurrency, error handling, reusable, utility, framework, library, samber',
      },
    },
    // NOTE: og:image, twitter:card and twitter:image are NOT declared here.
    // Docusaurus already emits them automatically from `themeConfig.image`
    // (see below). Declaring them a second time here produced duplicate
    // meta tags with identical values on every page.
    {
      tagName: 'meta',
      attributes: {
        name: 'twitter:creator',
        content: '@samuelberthe',
      },
    },
    // twitter:site complements twitter:creator for card attribution
    {
      tagName: 'meta',
      attributes: {
        name: 'twitter:site',
        content: '@samuelberthe',
      },
    },
    // og:locale is NOT declared here: Docusaurus already emits it from
    // `i18n.defaultLocale` ("en"). A second, hardcoded "en_US" value here
    // produced two conflicting og:locale tags on every page.
    // og:site_name provides branding context in social share cards
    {
      tagName: 'meta',
      attributes: {
        property: 'og:site_name',
        content: 'samber/lo',
      },
    },
    // Global entity graph: ties the site, the software project and its
    // author together for search engines and AI answer engines. Injected
    // once in <head>, present on every page (Docusaurus renders headTags
    // statically at build time, so this is in the SSG HTML, not client-only).
    {
      tagName: 'script',
      attributes: {
        type: 'application/ld+json',
      },
      innerHTML: JSON.stringify({
        '@context': 'https://schema.org',
        '@graph': [
          {
            '@type': 'Person',
            '@id': 'https://lo.samber.dev/#author',
            name: 'Samuel Berthe',
            url: 'https://samuel-berthe.fr',
            sameAs: [
              'https://github.com/samber',
              'https://twitter.com/samuelberthe',
              'https://samuelberthe.substack.com',
            ],
          },
          {
            '@type': 'WebSite',
            '@id': 'https://lo.samber.dev/#website',
            url: 'https://lo.samber.dev/',
            name: 'samber/lo',
            description: 'Documentation for samber/lo, a Lodash-style generic utility library for Go.',
            inLanguage: 'en',
            publisher: {'@id': 'https://lo.samber.dev/#author'},
          },
          {
            '@type': 'SoftwareSourceCode',
            '@id': 'https://lo.samber.dev/#software',
            name: 'samber/lo',
            alternateName: 'lo',
            description: 'A Lodash-style utility library for Go based on generics. 449 documented, type-safe helpers for slices, maps, strings, channels and iterators.',
            url: 'https://lo.samber.dev/',
            codeRepository: 'https://github.com/samber/lo',
            programmingLanguage: {
              '@type': 'ComputerLanguage',
              name: 'Go',
              alternateName: 'Golang',
            },
            runtimePlatform: 'Go 1.18+',
            license: 'https://github.com/samber/lo/blob/master/LICENSE',
            author: {'@id': 'https://lo.samber.dev/#author'},
            keywords: 'go, golang, generics, lodash, slice, map, functional programming, utility library',
          },
        ],
      }),
    },
    // NOTE: do NOT add a global <link rel="canonical"> here.
    // Docusaurus injects a correct per-page canonical automatically
    // based on `url` + `baseUrl` + the page path. A static href here
    // would override every page's canonical to the homepage, causing
    // Google to treat all docs pages as non-canonical duplicates.
  ],

  customFields: {
    sponsors: [
      {
        name: 'DBOS',
        url: 'https://www.dbos.dev/?utm_campaign=gh-smbr',
        title: 'DBOS - Durable workflow orchestration library for Go',
        logo_light: '/img/sponsors/dbos-black.png',
        logo_dark: '/img/sponsors/dbos-white.png',
      },
    ],
  },

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
        // No blog content exists yet. The classic preset enables a blog by
        // default, which was shipping an empty, indexable /blog page (200,
        // in the sitemap, zero content) — disable it until there is a first
        // real post.
        blog: false,
          sitemap: {
          lastmod: 'date',
          changefreq: 'weekly',
          priority: 0.7,
          ignorePatterns: ['/tags/**', '/search'],
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
    // og:type defaults to "website"; individual doc pages that need
    // "article" should override via their page's <Layout> or frontmatter.
    metadata: [
      // Open Graph tags use the `property` attribute, not `name` -- `name`
      // is ignored by strict OG parsers (Facebook, LinkedIn...).
      {property: 'og:type', content: 'website'},
      // Fallback description for pages that don't set their own
      {name: 'description', content: 'A Lodash-style utility library for Go. Slices, maps, channels, and more — type-safe, using generics.'},
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
          // label: 'GitHub',
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
              to: 'https://github.com/samber/lo/releases',
            },
            {
              label: 'Godoc',
              to: 'https://pkg.go.dev/github.com/samber/lo',
            },
            {
              label: 'License',
              to: 'https://github.com/samber/lo/blob/master/LICENSE',
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
            {
              label: 'Substack',
              to: 'https://samuelberthe.substack.com',
            },
          ],
        },
        {
          title: 'Ecosystem',
          items: [
            {
              label: 'ro — reactive streams for Go',
              to: 'https://ro.samber.dev',
            },
            {
              label: 'do — dependency injection',
              to: 'https://github.com/samber/do',
            },
            {
              label: 'mo — monads (Option, Result...)',
              to: 'https://github.com/samber/mo',
            },
            {
              label: 'oops — errors with context',
              to: 'https://github.com/samber/oops',
            },
            {
              label: 'hot — in-memory caching',
              to: 'https://github.com/samber/hot',
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
      [
      "posthog-docusaurus",
      {
          apiKey: "phc_uA762TtYyJ6UrbF5nzWutAJojstpC2EDptFpd2bBvWFY",
          appHost: "https://hogpost.samber.dev",
          enableInDevelopment: false, // optional,
          disableSessionRecording: true,
      },
  ],
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
