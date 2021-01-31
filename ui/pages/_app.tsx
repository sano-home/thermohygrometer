import React from 'react';
import { AppProps } from 'next/app';
import Head from 'next/head';
import { SWRConfig } from 'swr';

import { AppContainer } from '../components/AppContainer';

function MyApp({ Component, pageProps }: AppProps): JSX.Element {
  return (
    <>
      <Head>
        <title>Thermohygrometer</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <h2 className="app-title">Thermohygrometer</h2>
      <SWRConfig
        value={{
          refreshInterval: 10000,
          fetcher: (resource, init) =>
            fetch(resource, init).then((res) => res.json()),
        }}
      >
        <AppContainer>
          <Component {...pageProps} />
        </AppContainer>
      </SWRConfig>

      <style jsx global>{`
        html,
        body {
          padding: 0;
          margin: 0;
          font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto,
            Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue,
            sans-serif;
          background-color: #eeeeee;
        }

        * {
          box-sizing: border-box;
        }

        .app-title {
          text-align: center;
          background: #fff;
          margin: 0;
          padding: 16px;
        }
      `}</style>
    </>
  );
}

export default MyApp;
