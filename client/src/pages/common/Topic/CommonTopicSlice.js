import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";
import { convertDateToTimestamp } from "../../../utils/time";

const CommonTopicSlice = createSlice({
  name: "topic",
  initialState: {
    current: {},
    done: {}
  },
  reducers: {
    setResult: (state, action) => {
      state.current.result = {}
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchTopics.pending, (state, action) => {
        state.current.status = "loading";
      })
      .addCase(fetchTopics.fulfilled, (state, action) => {
        state.current.status = "idle";
        state.current.topics = action.payload.topic;
        state.current.total = action.payload.total;
      })
      .addCase(fetchCommonDoneTopics.pending, (state, action) => {
        state.done.status = "loading";
      })
      .addCase(fetchCommonDoneTopics.fulfilled, (state, action) => {
        state.done.status = "idle";
        state.done.topics = action.payload.topic;
        state.done.total = action.payload.total;
      })
      .addCase(registationTopic.pending, (state, action) => {
        state.current.status = "loading";
        state.current.result = {}
      })
      .addCase(registationTopic.fulfilled, (state, action) => {
        state.current.status = "idle";
        state.current.result = action.payload.result
      })
      .addCase(registationTopic.rejected, (state, action) => {
        state.current.error = action.payload.error;
        state.current.result = {}
        state.current.status = "error"
      })
      .addCase(cancelTopic.pending, (state, action) => {
        state.current.status = "loading"
        state.current.result = {}
      })
      .addCase(cancelTopic.fulfilled, (state, action) => {
        state.current.status = "idle"
        state.current.result = action.payload.result
      })
      .addCase(cancelTopic.rejected, (state, action) => {
        state.current.status = "error"
        state.current.error = action.payload.error
        state.current.result = {}
      })
  },
});

export const fetchTopics = createAsyncThunk("getTopics", async (payload) => {
  const timeStart = convertDateToTimestamp(payload.timeStart ? payload.timeStart : "")
  const timeEnd = convertDateToTimestamp(payload.timeEnd ? payload.timeEnd : "");
  const response = await client.get("/api/topics", {
    params: {
      search: payload.search ? payload.search : "",
      departmentID: payload.departmentID ? payload.departmentID : 0,
      facultyID: payload.facultyID ? payload.facultyID : 0,
      status: payload.status ? payload.status : 0,
      eventID: payload.eventID ? payload.eventID : 0,
      timeStart: timeStart,
      timeEnd: timeEnd,
      limit: payload.limit ? payload.limit : 0,
      offset: payload.offset ? payload.offset : 0,
      isCurrent: 1
    },
  });

  return response.data;
});

export const fetchCommonDoneTopics = createAsyncThunk("fetchCommonDoneTopics", async (payload) => {
  const timeStart = convertDateToTimestamp(payload.timeStart ? payload.timeStart : "")
  const timeEnd = convertDateToTimestamp(payload.timeEnd ? payload.timeEnd : "");
  const response = await client.get("/api/topics", {
    params: {
      search: payload.search ? payload.search : "",
      departmentID: payload.departmentID ? payload.departmentID : 0,
      facultyID: payload.facultyID ? payload.facultyID : 0,
      status: payload.status ? payload.status : 0,
      eventID: payload.eventID ? payload.eventID : 0,
      timeStart: timeStart,
      timeEnd: timeEnd,
      limit: payload.limit ? payload.limit : 0,
      offset: payload.offset ? payload.offset : 0,
      isCurrent: 2
    },
  });

  return response.data;
});

export const registationTopic = createAsyncThunk(
  "registrationTopic",
  async (payload) => {
    const resp = await client.put(`/api/topic-student-group/${payload.topicID}`, {
      studentID: payload.studentID,
    });

    return resp.data;
  }
);

export const cancelTopic = createAsyncThunk(
  "cancelTopic", 
  async (payload) => {
    const resp = await client.delete(`/api/topic-student-group/${payload.studentID}`)

    return resp.data
  }
)

export default CommonTopicSlice;
export const CommonTopicReducer = CommonTopicSlice.reducer;
export const CommonTopicAction = CommonTopicSlice.actions;
