import { FC, Fragment, ReactNode } from "react";
import { Children } from "react-transition-group/node_modules/@types/react";
import { cunord } from "../../util/utils";

type CardBorder = { radius?: string; width?: string; color?: string };

const DefaultBorder = {
  radius: "0.5rem",
  width: "1px",
  color: "#d5d5d5",
};

const resolveBorderDefaults = (border?: CardBorder): CardBorder => {
  border = cunord(border, DefaultBorder);
  border.width = cunord(border.width, DefaultBorder.width);
  border.radius = cunord(border.radius, DefaultBorder.radius);
  border.color = cunord(border.color, DefaultBorder.color);
  return border;
};

const Simple: FC<{
  className?: string;
  children: ReactNode;
  border?: CardBorder;
}> = ({ className, children, border }) => {
  border = resolveBorderDefaults(border);
  const styleObj = {
    borderStyle: "solid",
    borderWidth: border.width,
    borderRadius: border.radius,
    borderColor: "transparent",
    boxShadow: "0 0 10px 0 " + border.color,
  };
  return (
    <div className={className} style={styleObj}>
      {children}
    </div>
  );
};

const Rounded = {
  Simple,
};
export default Rounded;
