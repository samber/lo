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
    <div className="docSidebarWithSponsors">
      <OriginalDocSidebar {...props} />

      {sponsors.length > 0 && (
        <div className="sidebar-sponsors">
          <div className="sidebar-sponsors__title">💖 Sponsored by</div>
          <div className="sidebar-sponsors__logos">
            {sponsors.map((sponsor) => (
              <a
                key={sponsor.name}
                href={sponsor.url}
                target="_blank"
                rel="noopener noreferrer"
                className="sidebar-sponsors__logo-link"
                title={sponsor.title}
              >
                <div className="sidebar-sponsors__logo-wrapper">
                  <img
                    src={colorMode === 'dark' ? sponsor.logo_dark : sponsor.logo_light}
                    alt={sponsor.name}
                    className="sidebar-sponsors__logo"
                  />
                  <div className="sidebar-sponsors__logo-title">
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

