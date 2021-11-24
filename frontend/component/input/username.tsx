import {
  FormControl,
  IconButton,
  InputAdornment,
  InputLabel,
  OutlinedInput,
} from "@mui/material";
import { Person } from "@mui/icons-material";
import { useRef } from "react";
import { randstr } from "../../util/utils";
import InputDef from "./definition";
import { useCallback, useEffect } from "react";

const Username: InputDef<string> = ({
  className,
  label,
  value,
  errmsg = "",
  onChange,
}) => {
  const mRef = useRef<{ focus: () => void }>();

  let fid = "";
  useEffect(() => {
    fid = "pfield_" + randstr(25);
  }, []);

  const isErr = useCallback(() => errmsg.trim().length > 0, [errmsg]);

  return (
    <FormControl error={isErr()} className={className} variant="outlined">
      <InputLabel htmlFor={fid}>Username</InputLabel>
      <OutlinedInput
        id={fid}
        label={label}
        type="text"
        inputRef={mRef}
        value={value}
        onChange={(elem) => onChange(elem.target.value)}
        endAdornment={
          <InputAdornment position="end">
            <IconButton
              aria-label="username field"
              edge="end"
              onClick={() => mRef?.current?.focus()}
            >
              <Person />
            </IconButton>
          </InputAdornment>
        }
      />
    </FormControl>
  );
};

export default Username;
