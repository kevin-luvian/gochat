import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { RootState } from "../store";

const initialState = {
  access_token: "",
  refresh_token: "",
};

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    changeAccessToken: (state, action: PayloadAction<string>) => {
      state.access_token = action.payload;
    },
    changeRefreshToken: (state, action: PayloadAction<string>) => {
      state.refresh_token = action.payload;
    },
    clearTokens: (state) => {
      state = initialState;
    },
  },
});

export const { changeAccessToken, changeRefreshToken, clearTokens } =
  authSlice.actions;

export const selectAccessToken = (state: RootState) => state.auth.access_token;
export const selectRefreshToken = (state: RootState) =>
  state.auth.refresh_token;

export default authSlice.reducer;
