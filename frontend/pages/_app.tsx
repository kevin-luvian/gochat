import "../styles/globals.css";
import "bootstrap/dist/css/bootstrap.css";
import NoLayout from "../component/layout/noLayout";
import { Fragment } from "react";
import AppPropsDef from "../definition/appProps";
import { cunord } from "../util/utils";
import { CIF } from "../component/helper/condition";

function MyApp({ Component, pageProps }: AppPropsDef) {
  const Layout = cunord(Component.layout, NoLayout);
  const requiredLogin = cunord(Component.loginRequired, false);
  return (
    <Fragment>
      <Layout>
        <CIF condition={requiredLogin}>
          <h3>Login is required</h3>
        </CIF>
        <Component {...pageProps} />
      </Layout>
    </Fragment>
  );
}

export default MyApp;
