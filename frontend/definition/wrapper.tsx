import { FC, ReactNode, MouseEventHandler } from "react";

type Wrapper<P = {}> = FC<
  { className?: string; onClick?: MouseEventHandler; children: ReactNode } & P
>;

export default Wrapper;
