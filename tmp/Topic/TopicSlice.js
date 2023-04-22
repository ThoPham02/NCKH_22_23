import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const TopicSlice = createSlice({
    name: "topic",
    initialState: {
        status: "idle",
        topicList: [],
        total: 0,
        limit: 10,
        currentPage: 1,
    },
    reducers: {

    },
    extraReducers: (builder) => {
        builder
          .addCase(fetchTopic.pending, (state, action) => {
            state.status = "pending";
          })
          .addCase(fetchTopic.fulfilled, (state, action) => {
            state.status = "idle";
            state.topicList = action.payload.list_topic_registation;
            state.total = action.payload.total
          });
      },
})


export const fetchTopic = createAsyncThunk("fetchTopic", async (_, {getState}) => {
    // const filter = getState().topicFilter
    const topic = getState().topic

    let url = "/api/topic-registation?limit=" + topic.limit + "&offset=" + topic.currentPage
    const response = await client.get(url)

    return response.data
})

export default TopicSlice;

export const TopicReducer = TopicSlice.reducer;