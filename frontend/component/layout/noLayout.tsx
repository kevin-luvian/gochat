import { Fragment } from "react";
import FCLayout from "../../definition/flayout";
import Head from "next/head";

const layout: FCLayout = ({ children }) => (
  <Fragment>
    <Head>
      <title>No Layout</title>
      <meta name="description" content="static empty page" />
      <link rel="icon" href="/favicon.ico" />
    </Head>
    {children}
  </Fragment>
);

export default layout;
