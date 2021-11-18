import Wrapper from "../../definition/wrapper";
import { cat } from "../../util/utils";
import styles from "./button.module.scss";
import { Base as RoundedBtn } from "./rounded";
import { FC } from "react";

const GBtn: FC<{ className?: string }> = ({ className = "" }) => (
  <RoundedBtn className={cat(className, styles.btn_google, "unselectable")}>
   <img src="/img/google_icon.jpg" /> Login with Google
  </RoundedBtn>
);

export default GBtn;
