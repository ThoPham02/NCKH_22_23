import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const LectureTopicSlice = createSlice({
  name: "lectureTopic",
  initialState: {
    status: "idle",
    topic: [],
    currentEvent: {},
    result: {},
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchCurrentEvent.pending, (state, action) => {
        state.status = "loading";
      })
      .addCase(fetchCurrentEvent.fulfilled, (state, action) => {
        state.status = "idle";
        state.currentEvent = action.payload.event;
      });
  },
});

export const fetchCurrentEvent = createAsyncThunk(
  "fetchCurrentEvent",
  async () => {
    const resp = await client.get("/api/event-current");

    return resp.data;
  }
);

export const createTopic = createAsyncThunk("createTopic", async (payload) => {
  const resp = await client.post("/api/topic", {
    name: payload.name,
    departmentID: payload.departmentID,
    eventID: payload.eventID,
    lectureID: payload.lectureID,
  });

  return resp.data
});

export default LectureTopicSlice;
export const LectureTopicReducer = LectureTopicSlice.reducer;
export const LectureTopicAction = LectureTopicSlice.actions;
