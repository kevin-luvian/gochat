import { TextField } from "@mui/material";
import { FC } from "react-transition-group/node_modules/@types/react";

const Basic: FC<{ className?: string; label: string }> = ({
  className,
  label,
}) => <TextField className={className} label={label} variant="standard" />;

export const inpText = {
  Basic,
};
