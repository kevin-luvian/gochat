import {
  FormControl,
  IconButton,
  InputAdornment,
  InputLabel,
  OutlinedInput,
  FormHelperText,
} from "@mui/material";
import { Person } from "@mui/icons-material";
import { FC, useRef } from "react";
import { randstr } from "../../util/utils";

const Username: FC<{
  className?: string;
  label: string;
  value: string;
  error?: boolean;
  onChange: (val: string) => void;
}> = ({ className, label, value, error, onChange }) => {
  const fid = "ufield_" + randstr(25);
  const mRef = useRef<{ focus: () => void }>();
  //   sx={{ m: 1, width: "25ch" }}

  return (
    <FormControl error={error} className={className} variant="outlined">
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
