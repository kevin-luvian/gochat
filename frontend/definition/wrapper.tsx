import { FC, ReactNode } from "react";

type Wrapper<P = {}> = FC<{ className?: string; children: ReactNode } & P>;

export default Wrapper;
