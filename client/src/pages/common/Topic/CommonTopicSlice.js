import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";
import { convertDateToTimestamp } from "../../../utils/time";

const CommonTopicSlice = createSlice({
  name: "topic",
  initialState: {
    status: "idle",
  },
  reducers: {
    setResult: (state, action) => {
      state.result = {}
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchTopics.pending, (state, action) => {
        state.status = "loading";
      })
      .addCase(fetchTopics.fulfilled, (state, action) => {
        state.status = "idle";
        state.topics = action.payload.topic;
        state.total = action.payload.total;
      })
      .addCase(registationTopic.pending, (state, action) => {
        state.status = "loading";
        state.result = {}
      })
      .addCase(registationTopic.fulfilled, (state, action) => {
        state.status = "idle";
        state.result = action.payload.result
      })
      .addCase(registationTopic.rejected, (state, action) => {
        state.error = action.payload.error;
        state.result = {}
        state.status = "error"
      })
      .addCase(cancelTopic.pending, (state, action) => {
        state.status = "loading"
        state.result = {}
      })
      .addCase(cancelTopic.fulfilled, (state, action) => {
        state.status = "idle"
        state.result = action.payload.result
      })
      .addCase(cancelTopic.rejected, (state, action) => {
        state.status = "error"
        state.error = action.payload.error
        state.result = {}
      });
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

export const addTopic = createAsyncThunk("addTopic", async (payload) => {
  const resp = await client.post(`/api/topic`, {
    name: payload.name,
    departmentID: payload.departmentID,
    lectureID: payload.lectureID,
    estimateStudent: payload.estimateStudent,
    description: payload.description,
  })

  fetchTopics({lectureID: payload.lectureID})

  return resp.data
})

export default CommonTopicSlice;
export const CommonTopicReducer = CommonTopicSlice.reducer;
export const CommonTopicAction = CommonTopicSlice.actions;
