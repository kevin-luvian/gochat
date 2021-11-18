import Wrapper from "../../definition/wrapper";
import { cat } from "../../util/utils";
import styles from "./button.module.scss";

export const Base: Wrapper = ({ className = "", children }) => (
  <div className={cat(className, styles.rounded_button, "unselectable")}>
    <div className="mx-auto w-fit">{children}</div>
  </div>
);

const Primary: Wrapper = ({ className = "", children }) => (
  <Base className={cat(className, styles.btn_primary)}>{children}</Base>
);

const Rounded = {
  Primary,
};
export default Rounded;
