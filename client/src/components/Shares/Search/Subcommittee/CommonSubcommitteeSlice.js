import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const CommonSubcommitteeSlice = createSlice({
  name: "event",
  initialState: {
    status: "idle",
    subcommittees: [],
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchCommonSubcommittee.pending, (state, action) => {
        state.status = "pending";
      })
      .addCase(fetchCommonSubcommittee.fulfilled, (state, action) => {
        state.status = "idle";
        state.subcommittees = action.payload
      });
  },
});

export const fetchCommonSubcommittee = createAsyncThunk("fetchCommonSubcommittee", async (payload) => {
  const response = await client.get("/api/subcommittees");

  return response.data.subcommittees.filter(c => !payload.facultyID || c.facultyID === payload.facultyID);
});

export const CommonSubcommitteeReducer = CommonSubcommitteeSlice.reducer;
export const CommonSubcommitteeAction = CommonSubcommitteeSlice.actions;
export default CommonSubcommitteeSlice;
