import siteConfig from '@generated/docusaurus.config';
export default function prismIncludeLanguages(PrismObject) {
  const {
    themeConfig: {prism},
  } = siteConfig;
  const {additionalLanguages, defaultLanguage} = prism;
  
  // Add Go language support
  globalThis.Prism = PrismObject;
  
  // eslint-disable-next-line import/no-dynamic-require, global-require
  require(`prismjs/components/prism-${defaultLanguage}`);
  
  additionalLanguages.forEach((lang) => {
    // eslint-disable-next-line import/no-dynamic-require, global-require
    require(`prismjs/components/prism-${lang}`);
  });
  
  delete globalThis.Prism;
}
