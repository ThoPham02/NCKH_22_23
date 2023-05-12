import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const LoginSlice = createSlice({
  name: "user",
  initialState: {
    status: "idle",
    token: {},
    user: { typeAccount: 0 },
  },
  reducers: {
    logout: (state, _) => {
      state.token = {};
      state.user = { typeAccount: 0 };
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchLogin.pending, (state, acction) => {
        state.status = "loading";
      })
      .addCase(fetchLogin.fulfilled, (state, acction) => {
        state.token = acction.payload.token;
        state.user = acction.payload.user;
        state.status = "idle";
      });
  },
});

export const fetchLogin = createAsyncThunk("user/fetchLogin", async (payload) => {
  const url = "/user/login"
  const response = await client.post(url, {
    username: payload.username,
    password: payload.password
  })
  console.log(response)

  return response
})



export default LoginSlice;
export const LoginReducer = LoginSlice.reducer;
export const LoginActions = LoginSlice.actions;
