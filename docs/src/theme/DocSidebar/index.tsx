import React from 'react';
import OriginalDocSidebar from '@theme-original/DocSidebar';
import type DocSidebarProps from '@theme/DocSidebar';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import {useColorMode} from '@docusaurus/theme-common';

type Sponsor = {
  name: string;
  url: string;
  title: string;
  logo_light: string;
  logo_dark: string;
};

export default function DocSidebarWrapper(props: DocSidebarProps) {
  const {siteConfig} = useDocusaurusContext();
  const sponsors = (siteConfig.customFields?.sponsors ?? []) as Sponsor[];
  const {colorMode} = useColorMode();

  return (
    <div className="docSidebarWithSpnsors">
      <OriginalDocSidebar {...props} />

      {sponsors.length > 0 && (
        <div className="sidebar-spnsors">
          <div className="sidebar-spnsors__title">💖 Sponsored by</div>
          <div className="sidebar-spnsors__logos">
            {sponsors.map((sponsor) => (
              <a
                key={sponsor.name}
                href={sponsor.url}
                target="_blank"
                rel="noopener noreferrer"
                className="sidebar-spnsors__logo-link"
                title={sponsor.title}
              >
                <div className="sidebar-spnsors__logo-wrapper">
                  <img
                    src={colorMode === 'dark' ? sponsor.logo_dark : sponsor.logo_light}
                    alt={sponsor.name}
                    className="sidebar-spnsors__logo"
                  />
                  <div className="sidebar-spnsors__logo-title">
                    {sponsor.title}
                  </div>
                </div>
              </a>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}

