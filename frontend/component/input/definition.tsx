import { FC, ReactNode, MouseEventHandler } from "react";

type InputDef<T, P = {}> = FC<
  {
    className?: string;
    label?: string;
    value: T;
    onChange: (val: T) => void;
    errmsg?: string;
  } & P
>;

export default InputDef;
