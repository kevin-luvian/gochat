import { Fragment, FC, useState, useEffect } from "react";
import { useAppSelector, useAppDispatch } from "../../../redux/hooks";
import {
  changeAccessToken,
  changeRefreshToken,
  selectAccessToken,
} from "../../../redux/states/auth";
import { getJWTExp } from "../../../util/utils";
import { CIF } from "../../helper/condition";
import Router from "next/router";

const AuthGateway: FC = ({ children }) => {
  const [loading, setLoading] = useState(true);

  const access_token = useAppSelector(selectAccessToken);
  const dispatch = useAppDispatch();

  const isAuthenticated = async () => {
    if (access_token === "") return false;

    // check if access_token expired.
    const exp = getJWTExp(access_token);
    const fiveMinutes = 1000 * 60 * 5;

    if (exp - new Date().getTime() > fiveMinutes) return true;

    // if expired, refresh the token.

    // dipatch accesstoken and refreshtoken change.
    dispatch(changeAccessToken(""));
    dispatch(changeRefreshToken(""));
    return true;
  };

  useEffect(() => {
    (async () => {
      const isAuth = await isAuthenticated();
      if (isAuth) {
        setLoading(false);
      } else {
        Router.push("/login");
      }
    })();
  }, []);

  return (
    <Fragment>
      <CIF condition={loading}>Loading...</CIF>
      <CIF condition={!loading}>{children}</CIF>
    </Fragment>
  );
};

export default AuthGateway;
