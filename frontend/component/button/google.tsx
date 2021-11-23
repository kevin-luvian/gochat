import { FC } from "react";
import { Colored, ColorOptions } from "./mui";
import GoogleIcon from "@mui/icons-material/Google";
import { MouseEventHandler } from "react-transition-group/node_modules/@types/react";

const GBtn: FC<{ className?: string; onClick?: MouseEventHandler }> = (
  props
) => {
  const color: ColorOptions = { main: "#cf4332", contrastText: "white" };
  return (
    <Colored {...props} color={color} startIcon={<GoogleIcon />}>
      Login with Google
    </Colored>
  );
};
export default GBtn;
