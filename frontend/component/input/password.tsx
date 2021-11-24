import {
  FormControl,
  IconButton,
  InputAdornment,
  InputLabel,
  OutlinedInput,
  FormHelperText,
} from "@mui/material";
import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import { useState } from "react";
import { randstr } from "../../util/utils";
import { CIF } from "../helper/condition";
import InputDef from "./definition";
import { useCallback, useEffect } from "react";

const Password: InputDef<string> = ({
  className,
  label,
  value,
  errmsg = "",
  onChange,
}) => {
  const [show, setShow] = useState(false);
  let fid = "";

  useEffect(() => {
    fid = "pfield_" + randstr(25);
  }, []);

  const handleChange = (val: string) => onChange?.(val.trim());

  const isErr = useCallback(() => errmsg.trim().length > 0, [errmsg]);

  return (
    <FormControl error={isErr()} className={className} variant="outlined">
      <InputLabel htmlFor={fid}>Password</InputLabel>
      <OutlinedInput
        id={fid}
        label={label}
        type={show ? "text" : "password"}
        value={value}
        onChange={(elem) => handleChange(elem.target.value)}
        endAdornment={
          <InputAdornment position="end">
            <IconButton
              aria-label="toggle password visibility"
              onClick={() => setShow(!show)}
              onMouseDown={(event) => event.preventDefault()}
              edge="end"
            >
              {show ? <VisibilityOff /> : <Visibility />}
            </IconButton>
          </InputAdornment>
        }
      />
      <CIF condition={isErr()}>
        <FormHelperText>{errmsg}</FormHelperText>
      </CIF>
    </FormControl>
  );
};

export default Password;
