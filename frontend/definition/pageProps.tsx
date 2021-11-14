import { NextComponentType, NextPageContext } from "next";
import FCLayout from "./flayout";

type PageProps = NextComponentType<NextPageContext, any, {}> & {
  loginRequired?: boolean;
  layout?: FCLayout;
};

export default PageProps;
