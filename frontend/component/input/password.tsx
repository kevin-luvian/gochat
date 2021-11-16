import {
  FormControl,
  IconButton,
  InputAdornment,
  InputLabel,
  OutlinedInput,
} from "@mui/material";
import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import { FC, useState } from "react";
import { randstr } from "../../util/utils";

const Password: FC<{
  className?: string;
  label: string;
  value: string;
  onChange: (val: string) => void;
}> = ({ className, label, value, onChange }) => {
  const [show, setShow] = useState(false);
  const fid = "pfield_" + randstr(25);

  return (
    <FormControl className={className} variant="outlined">
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
    </FormControl>
  );
};

export default Password;
