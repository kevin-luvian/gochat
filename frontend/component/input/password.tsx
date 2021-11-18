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
import { FC, useState } from "react";
import { randstr } from "../../util/utils";
import { CIF } from "../helper/condition";

const Password: FC<{
  className?: string;
  label: string;
  value: string;
  error?: boolean;
  onChange: (val: string) => void;
}> = ({ className, label, value, error = false, onChange }) => {
  const [show, setShow] = useState(false);
  const fid = "pfield_" + randstr(25);

  return (
    <FormControl error={error} className={className} variant="outlined">
      <InputLabel htmlFor={fid}>Password</InputLabel>
      <OutlinedInput
        id={fid}
        label={label}
        type={show ? "text" : "password"}
        value={value}
        onChange={(elem) => onChange(elem.target.value)}
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
      <CIF condition={error}>
        <FormHelperText>username or password is invalid</FormHelperText>
      </CIF>
    </FormControl>
  );
};

export default Password;
