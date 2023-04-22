// import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
// import client from "../../apis";
// import { CompareTime } from "../../utils/time";

// const loginSlice = createSlice({
//   name: "user",
//   initialState: {
//     token: {},
//     user: {}
//   },
//   reducers: {
//     login:(state, action) => {
//       state.token = action.payload.token
//       state.user = action.payload.user
//     },
//     logout: (state, action) => {
//       state.token = {};
//       state.user = {};
//     },
//   },
//   extraReducers: (builder) => {
//     builder
//       .addCase(fetchLoginUser.pending, (state, action) => {
//         state.status = "loading";
//       })
//       .addCase(fetchLoginUser.fulfilled, (state, action) => {
//         state.token = {
//           accessToken: action.payload.access_token,
//           accessExpiredAt: action.payload.access_expired_at,
//           refreshToken: action.payload.refresh_token,
//           refreshExpiredAt: action.payload.refresh_expired_at,
//         };
//         state.user = {
//           name: action.payload.user.name,
//           email: action.payload.user.email,
//           type: action.payload.user.account_type,
//         }
//         state.status = "idle";
//       })
//       .addCase(getUserInfo.pending, (state,action) => {
//         state.status = "loading";
//       })
//       .addCase(getUserInfo.rejected, (state,action) => {
//         state.status = "rejected";
//         state.token = {}
//         state.user = {}
//         state.info = {}
//       })
//       .addCase(getUserInfo.fulfilled, (state, action) => {
//         state.status = "idle";
//         state.info = {}
//       });
//   },
// });

// export const fetchLoginUser = createAsyncThunk(
//   "user/login",
//   async (payload) => {
//     const response = await client.post("/api/user/login", {
//       name: payload.username,
//       password: payload.password,
//     });

//     switch (response.data.user.account_type) {
//       case 1:
//         response.data.user.account_type = "student"
//         break;
//       case 2:
//         response.data.user.account_type = "lecture"
//         break;
//       case 3:
//         response.data.user.account_type = "department"
//         break;
//       case 4:
//         response.data.user.account_type = "faculty"
//         break;
//       case 5:
//         response.data.user.account_type = "admin"
//         break;
//       default:
//         break;
//     }
//     return response.data;
//   }
// );

// const checkToken = async (token) => {
//   const isExpire = !CompareTime(token.accessExpiredAt)
//   if (isExpire) {
//     const response = await client.post("/api/user/refresh-token", {
//       access_token: token.accessToken,
//     });
    
//     return response.data.access_token;
//   }
//   return token.accessToken
// };

// export const getUserInfo = createAsyncThunk(
//   "/api/user/info",
//   async (_, { getState }) => {
//     const state = getState();
//     checkToken(state.login.token)
//     const response = await client.get("/api/user/info");

//     return response.data;
//   }
// );

// export default loginSlice;

// export const loginReducer = loginSlice.reducer;
