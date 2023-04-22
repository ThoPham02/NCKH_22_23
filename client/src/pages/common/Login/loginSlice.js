import { createSlice } from "@reduxjs/toolkit";

const LoginSlice = createSlice({
  name: "user",
  initialState: {
    token: {},
    user: {typeAccount: 0}
  },
  reducers: {
    login:(state, action) => {
      state.token = action.payload.token
      state.user = action.payload.user
    },
    logout: (state, _) => {
      state.token = {};
      state.user = {typeAccount: 0};
    },
  }
})

export default LoginSlice;
export const LoginReducer = LoginSlice.reducer;
export const LoginActions = LoginSlice.actions;