import { FC, Fragment, ReactNode } from "react";

/**
 * check if condition true then render the children elements
 */
export const CIF: FC<{ condition: boolean; children: ReactNode }> = ({
  condition,
  children,
}) => {
  return <Fragment>{condition && children}</Fragment>;
};
