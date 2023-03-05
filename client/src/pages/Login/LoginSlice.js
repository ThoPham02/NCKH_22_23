import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../apis";

const loginSlice = createSlice({
  name: "user",
  initialState: {
    status: "idle",
    accessToken: "",
    user: {},
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchLoginUser.pending, (state, action) => {
        state.status = "loading";
      })
      .addCase(fetchLoginUser.fulfilled, (state, action) => {
        state.accessToken = action.payload.access_token;
        state.user = action.payload.user;
        state.status = "idle";
      });
  },
});

export const fetchLoginUser = createAsyncThunk(
  "user/login",
  async (payload) => {
    const response = await client.post("/user/login", {
      username: payload.username,
      password: payload.password,
    });
    return response.data;
  }
);



export default loginSlice;
