import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const LoginSlice = createSlice({
  name: "user",
  initialState: {
    status: "idle",
    token: {},
    user: { role: 0 },
  },
  reducers: {
    logout: (state, _) => {
      state.token = {};
      state.user = { role: 0 };
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchLogin.pending, (state, acction) => {
        state.status = "loading";
      })
      .addCase(fetchLogin.fulfilled, (state, acction) => {
        state.token = acction.payload.authToken;
        state.user = acction.payload.user;
        state.status = "idle";
      });
  },
});

export const fetchLogin = createAsyncThunk("user/fetchLogin", async (payload) => {
  const url = "/user/login"
  const response = await client.post(url, payload);

  return response.data
})



export default LoginSlice;
export const LoginReducer = LoginSlice.reducer;
export const LoginActions = LoginSlice.actions;
