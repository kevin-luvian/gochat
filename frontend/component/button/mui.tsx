import { Button } from "@mui/material";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { ReactNode } from "react";
import Wrapper from "../../definition/wrapper";
import { colors } from "../../styles/themes";

export const Primary: Wrapper = (props) => (
  <Colored
    {...props}
    color={{
      main: colors.primary,
      contrastText: colors.onPrimary,
    }}
  />
);

export type ColorOptions = {
  main: string;
  light?: string;
  dark?: string;
  contrastText?: string;
};

export const Colored: Wrapper<{ color: ColorOptions; startIcon?: ReactNode }> =
  ({ children, color, ...props }) => {
    const theme = createTheme({
      components: {
        MuiButton: {
          styleOverrides: {
            root: {
              lineHeight: "inherit",
            },
          },
        },
      },
      palette: {
        primary: color,
      },
    });
    return (
      <ThemeProvider theme={theme}>
        <Button {...props} variant="contained" size="medium" color="primary">
          {children}
        </Button>
      </ThemeProvider>
    );
  };
