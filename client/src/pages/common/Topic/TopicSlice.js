import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";
import { convertDateToTimestamp } from "../../../utils/time";

const TopicSlice = createSlice({
  name: "topic",
  initialState: {
    status: "idle",
    topics: [],
    total: 0,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
    .addCase(fetchTopics.pending, (state, action) => {
        state.status = "loading";
    })
    .addCase(fetchTopics.fulfilled, (state, action) => {
        state.status = "idle"
        state.topics = action.payload.topic
        state.total = action.payload.total
    });
  },
});

export const fetchTopics = createAsyncThunk("getTopics", async (payload) => {
  const timeStart = convertDateToTimestamp(payload.timeStart);
  const timeEnd = convertDateToTimestamp(payload.timeEnd);
  console.log(payload)
  const response = await client.get("/api/topics", {
    params: {
      search: payload.search === "" ? " " : payload.search,
      departmentID: payload.departmentID,
      facultyID: payload.facultyID,
      status: payload.status,
      lectureID: 0,
      eventID: 0,
      subcommitteeID: 0,
      timeStart: timeEnd,
      timeEnd: timeStart,
      limit: payload.limit,
      offset: payload.offset,
    },
  });

  return response.data;
});

export default TopicSlice;
export const TopicReducer = TopicSlice.reducer;
export const TopicAction = TopicSlice.actions;
