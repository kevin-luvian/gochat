import { Fragment } from "react";
import Head from "next/head";
import FCLayout from "../../definition/flayout";

const layout: FCLayout = ({ children }) => (
  <Fragment>
    <Head>
      <title>One Layout</title>
      <meta name="description" content="static empty page" />
      <link rel="icon" href="/favicon.ico" />
    </Head>
    {children}
  </Fragment>
);

export default layout;
