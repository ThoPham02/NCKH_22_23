import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const CommonEventSlice = createSlice({
  name: "event",
  initialState: {
    status: "idle",
    events: [],
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchEvents.pending, (state, action) => {
        state.status = "pending";
      })
      .addCase(fetchEvents.fulfilled, (state, action) => {
        state.status = "idle";
        state.events = action.payload.events
      });
  },
});

export const fetchEvents = createAsyncThunk("fetchEvents", async () => {
  const response = await client.get("/api/events");

  return response.data;
});

export const CommonEventReducer = CommonEventSlice.reducer;
export const CommonEventAction = CommonEventSlice.actions;
export default CommonEventSlice;
