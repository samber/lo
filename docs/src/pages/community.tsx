import React from 'react';

import styles from './community.module.css';
import classnames from 'classnames';
import Layout from '@theme/Layout';
import Link from '@docusaurus/Link';

import useDocusaurusContext from '@docusaurus/useDocusaurusContext';

function Community() {
  const context = useDocusaurusContext();

  return (
    <Layout title="Community" description="Where to ask questions and find your soulmate">
      <header className="hero">
        <div className="container text--center">
          <h1>Community</h1>
          <div className="hero--subtitle">
            These are places where you can ask questions and find your soulmate (no promises).
            <br/>
            "If you want to go fast, go alone. If you want to go far, go together."
          </div>
          <img className={styles.headerImg} src="/img/go-community.png" />
        </div>
      </header>
      <main>
        <div className="container">
          <div className="row margin-vert--lg">
            <div className="col text--center padding-vert--md">
              <div className="card">
                <div className="card__header">
                  <i className={classnames(styles.icon, styles.chat)}></i>
                </div>
                <div className="card__body">
                  <p>Report bugs or suggest improvements</p>
                </div>
                <div className="card__footer">
                  <Link to="https://github.com/samber/lo/issues" className="button button--outline button--primary button--block">Open new issue</Link>
                </div>
              </div>
            </div>

            <div className="col text--center padding-vert--md">
              <div className="card">
                <div className="card__header">
                  <i className={classnames(styles.icon, styles.chat)}></i>
                </div>
                <div className="card__body">
                  <p>You like this project?</p>
                </div>
                <div className="card__footer">
                  <Link to="https://github.com/samber/lo?tab=readme-ov-file#-contributing" className="button button--outline button--primary button--block">Start contributing!</Link>
                </div>
              </div>
            </div>

            <div className="col text--center padding-vert--md">
              <div className="card">
                <div className="card__header">
                  <i className={classnames(styles.icon, styles.twitter)}></i>
                </div>
                <div className="card__body">
                  <p>Follow &#64;samuelberthe on Twitter</p>
                </div>
                <div className="card__footer">
                  <Link to="https://twitter.com/samuelberthe" className="button button--outline button--primary button--block">Follow &#64;SamuelBerthe</Link>
                </div>
              </div>
            </div>

            <div className="col text--center padding-vert--md">
              <div className="card">
                <div className="card__header">
                  <i className={classnames(styles.icon, styles.email)}></i>
                </div>
                <div className="card__body">
                  <p>For sensitive or security-related queries, send us an email</p>
                </div>
                <div className="card__footer">
                  <Link to="mailto:contact@samuel-berthe.fr" className="button button--outline button--primary button--block">contact&#64;samuel-berthe.fr</Link>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </Layout>
  );
}

export default Community;