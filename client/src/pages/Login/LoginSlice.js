import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../apis";
import { CompareTime } from "../../utils/time";

const loginSlice = createSlice({
  name: "user",
  initialState: {
    status: "idle",
    token: {},
    user: {},
    info: {},
  },
  reducers: {
    logout: (state, action) => {
      state.status = "idle";
      state.token = {};
      state.user = {};
    },
  },
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
          refreshExpiredAt: action.payload.refresh_expired_at,
        };
        state.user = action.payload.user;
        state.status = "idle";
      })
      .addCase(getUserInfo.pending, (state,action) => {
        state.status = "loading";
      })
      .addCase(getUserInfo.rejected, (state,action) => {
        state.status = "rejected";
        state.token = {}
        state.user = {}
        state.info = {}
      })
      .addCase(getUserInfo.fulfilled, (state, action) => {
        state.status = "idle";
        state.info = action.payload;
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

const checkToken = async (token) => {
  const isExpire = !CompareTime(token.accessExpiredAt)
  if (isExpire) {
    const response = await client.post("/api/user/refresh-token", {
      access_token: token.accessToken,
    });
    
    return response.data.access_token;
  }
  return token.accessToken
};

export const getUserInfo = createAsyncThunk(
  "/api/user/info",
  async (_, { getState }) => {
    const state = getState();
    checkToken(state.login.token)
    const response = await client.get("/api/user/info");

    return response.data;
  }
);

export default loginSlice;

export const loginReducer = loginSlice.reducer;
