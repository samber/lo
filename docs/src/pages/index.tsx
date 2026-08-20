import type {ReactNode} from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import Heading from '@theme/Heading';
import styles from './index.module.css';

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Type-safe Utilities for Go',
    Svg: require('@site/static/img/oxygen-tube.svg').default,
    description: (
      <>
        Generic utilities bringing type safety and convenience to daily Go programming.
      </>
    ),
  },
  {
    title: 'Comprehensive API Coverage',
    Svg: require('@site/static/img/compass.svg').default,
    description: (
      <>
        A rich set of helpers for slices, maps, and more, ready for any data task in Go.
      </>
    ),
  },
  {
    title: 'Built for Productivity',
    Svg: require('@site/static/img/backpacks.svg').default,
    description: (
      <>
        Minimal dependencies and intuitive APIs for seamless adoption and fast results. No breaking changes.
      </>
    ),
  },
];

function Feature({title, Svg, description}: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

function HomepageFeatures(): ReactNode {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
          {siteConfig.title}
        </Heading>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div className={styles.buttons} style={{marginBottom: '10px'}}>
          <Link
            className="button button--secondary button--lg"
            to="/docs/about">
            Intro
          </Link>
        </div>
        <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg"
            to="/docs/getting-started">
            Getting started - 5min ⏱️
          </Link>
        </div>
        <div className={clsx('row', styles.stats)}>
          <div className="col text--center">
            <strong>449</strong> documented helpers
          </div>
          <div className="col text--center">
            <strong>21k+</strong> GitHub stars
          </div>
          <div className="col text--center">
            <strong>12,500+</strong> packages import it
          </div>
          <div className="col text--center">
            No breaking changes before v2, <strong>MIT</strong> licensed
          </div>
        </div>
        <p style={{marginTop: '1rem', fontSize: '0.9rem'}}>
          <Link to="/production-ready">Evaluating lo for production? See the full breakdown →</Link>
        </p>
      </div>
    </header>
  );
}

export default function Home(): JSX.Element {
  return (
    <Layout
      title="Go Generics Utility Library — 449 Type-safe Helpers"
      description="A Lodash-style utility library for Go 1.18+ generics: 449 type-safe helpers for slices, maps, strings, channels and iterators. Zero dependencies.">
      <HomepageHeader />
      <main>
        <HomepageFeatures />
      </main>
    </Layout>
  );
}
