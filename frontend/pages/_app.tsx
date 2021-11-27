import "../styles/globals.css";
import "bootstrap/dist/css/bootstrap.css";
import DefaultLayout from "../component/layout/default";
import { Fragment, useEffect, useState } from "react";
import AppPropsDef from "../definition/appProps";
import { cunord, getJWTExp } from "../util/utils";
import { CIF } from "../component/helper/condition";
import { Provider } from "react-redux";
import { store } from "../redux/store";
import AuthGateway from "../component/auth/gateway";

function MyApp({ Component, pageProps }: AppPropsDef) {
  const Layout = cunord(Component.layout, DefaultLayout);
  const requiredLogin = cunord(Component.loginRequired, false);
  const Gateway = requiredLogin ? AuthGateway : Fragment;
  return (
    <Provider store={store}>
      <Gateway>
        <Layout>
          <CIF condition={requiredLogin}>
            <p>Login is required for this page</p>
          </CIF>
          <Component {...pageProps} />
        </Layout>
      </Gateway>
    </Provider>
  );
}

export default MyApp;
