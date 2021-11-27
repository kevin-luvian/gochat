import { FC } from "react";
import { Colored, ColorOptions } from "./mui";
import GoogleIcon from "@mui/icons-material/Google";
import authRepo from "../../axios/repositories/auth";

const GBtn: FC<{ className?: string }> = (props) => {
  const handleGLogin = async () => {
    const { oauth_url } = await authRepo.redirectGoogle("/redirect/google");
    window.location.assign(oauth_url);
  };

  const color: ColorOptions = { main: "#cf4332", contrastText: "white" };
  return (
    <Colored
      {...props}
      color={color}
      startIcon={<GoogleIcon />}
      onClick={handleGLogin}
    >
      Login with Google
    </Colored>
  );
};
export default GBtn;
