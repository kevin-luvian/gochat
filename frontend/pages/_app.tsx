import "../styles/globals.css";
import "bootstrap/dist/css/bootstrap.css";
import DefaultLayout from "../component/layout/default";
import { Fragment } from "react";
import AppPropsDef from "../definition/appProps";
import { cunord } from "../util/utils";
import { CIF } from "../component/helper/condition";

function MyApp({ Component, pageProps }: AppPropsDef) {
  const Layout = cunord(Component.layout, DefaultLayout);
  const requiredLogin = cunord(Component.loginRequired, false);
  return (
    <Fragment>
      <Layout>
        <CIF condition={requiredLogin}>
          <p>Login is required for this page</p>
        </CIF>
        <Component {...pageProps} />
      </Layout>
    </Fragment>
  );
}

export default MyApp;
