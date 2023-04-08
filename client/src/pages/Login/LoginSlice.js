import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../apis";

const loginSlice = createSlice({
  name: "user",
  initialState: {
    status: "idle",
    token: {},
    user: {},
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchLoginUser.pending, (state, action) => {
        state.status = "loading";
      })
      .addCase(fetchLoginUser.fulfilled, (state, action) => {
        state.token = {
          accessToken: action.payload.access_token,
          accessExpiredAt: action.payload.access_expired_at,
          refreshToken: action.payload.refresh_token,
          refreshExpiredAt: action.payload.refresh_expired_at
        }
        state.user = action.payload.user;
        state.status = "idle";
      });
  },
});

export const fetchLoginUser = createAsyncThunk(
  "user/login",
  async (payload) => {
    const response = await client.post("/api/user/login", {
      name: payload.username,
      password: payload.password,
    });
    return response.data;
  }
);

export default loginSlice;

export const loginReducer = loginSlice.reducer