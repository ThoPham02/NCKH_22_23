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
      })
      .addCase(fetchMyTopic.pending, (state, action) => {
        state.status = "loading"
      })
      .addCase(fetchMyTopic.fulfilled, (state, action) => {
        state.topic = action.payload.topic;
        state.status = "idle"
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

export const fetchMyTopic = createAsyncThunk("fetchMyTopic", async (payload) => {
  const resp = await client.get("/api/topics", {
    params: {
      search: " ",
      departmentID: 0,
      facultyID: 0,
      status: 0,
      lectureID: payload.lectureID,
      eventID: 0,
      subcommitteeID: 0,
      timeStart: 0,
      timeEnd: 0,
      limit: 0,
      offset: 0,
    },
  })

  return resp.data
})

export default LectureTopicSlice;
export const LectureTopicReducer = LectureTopicSlice.reducer;
export const LectureTopicAction = LectureTopicSlice.actions;
