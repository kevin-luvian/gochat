import { Fragment } from "react";
import FCLayout from "../../definition/flayout";
import Head from "next/head";
import styles from "../../styles/Home.module.css";
import Image from "next/image";

const layout: FCLayout = ({ children }) => (
  <Fragment>
    <Head>
      <title>No Layout</title>
      <meta name="description" content="static empty page" />
      <link rel="icon" href="/favicon.ico" />
    </Head>
    {children}
    <footer className={styles.footer}>
      <a
        href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
        target="_blank"
        rel="noopener noreferrer"
      >
        Powered by{" "}
        <span className={styles.logo}>
          <Image src="/vercel.svg" alt="Vercel Logo" width={72} height={16} />
        </span>
      </a>
    </footer>
  </Fragment>
);

export default layout;
