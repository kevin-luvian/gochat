import { FC, Fragment, ReactNode, CSSProperties } from "react";
import { Children } from "react-transition-group/node_modules/@types/react";
import Wrapper from "../../definition/wrapper";
import { cunord } from "../../util/utils";

type ButtonColors = { bg: string };

const Base: Wrapper<{ colors: ButtonColors }> = ({
  className,
  children,
  colors,
}) => {
  const styles: CSSProperties = {
    backgroundColor: colors.bg,
  };
  return (
    <div className={className} style={styles}>
      {children}
    </div>
  );
};

const Primary: Wrapper = ({ className, children }) => {
  return (
    <Base className={className} colors={{ bg: "red" }}>
      <div className="mx-auto" style={{ width: "fit-content" }}>
        {children}
      </div>
    </Base>
  );
};

const Rounded = {
  Primary,
};
export default Rounded;
