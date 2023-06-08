import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const LectureMyTopicSlice = createSlice({
  name: "lecturemytopic",
  initialState: {
    status: "idle"
  },
  reducers: {},
  extraReducers: builder => {
    builder
      .addCase(fetchLectureTopic.pending, (state, action) => {
        state.status = "loading"
      })
      .addCase(fetchLectureTopic.fulfilled, (state, action) => {
        state.status = "idle"
        state.topics = action.payload.topic
      })
      .addCase(addTopic.pending, (state, action) => {
        state.status = "loading"
      })
      .addCase(addTopic.fulfilled, (state, action) => {
        state.status = "idle"
        state.topics = action.payload.topic
      })
  }
})

export const fetchLectureTopic = createAsyncThunk("fetchLectureTopic", async (payload) => {
  const response = await client.get("/api/topics", {
    params: {
      lectureID: payload.lectureID,
    },
  });

  return response.data;
})


export const addTopic = createAsyncThunk("addTopic", async (payload) => {
  await client.post(`/api/topic`, {
    name: payload.name,
    departmentID: payload.departmentID,
    lectureID: payload.lectureID,
    estimateStudent: payload.estimateStudent,
    description: payload.description,
  })

  const response = await client.get("/api/topics", {
    params: {
      lectureID: payload.lectureID,
    },
  });

  return response.data;
})


export const addReport = createAsyncThunk("addReport" , async (payload) => {
  await client.post("/api/report", {
    topicID: payload.id,
    description: payload.description,
    reportUrl: payload.url,
    stageID: payload.stageID
  })
  const response = await client.get("/api/topics", {
    params: {
      lectureID: payload.lectureID,
    },
  });

  return response.data;
})

export default LectureMyTopicSlice;
export const LectureMyTopicReducer = LectureMyTopicSlice.reducer
export const LectureMyTopicAction = LectureMyTopicSlice.actions