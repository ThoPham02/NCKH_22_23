import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const TopicDetailSlice = createSlice({
    name: "topicDetail",
    initialState: {
        status: "idle",
        topicDetail: {}
    },
    reducers: {},
    extraReducers: builder => {
        builder.addCase(fetchTopicDetail.pending, (state, action) => {
            state.status = "loading";
        }).addCase(fetchTopicDetail.fulfilled, (state, action) => {
            state.status = "idle";
            state.topicDetail = action.payload.topicDetail
        })
    }
})

export const fetchTopicDetail = createAsyncThunk("fetchTopicDetail", async (topicID) => {
    const resp = await client.get(`/api/topic/${topicID}`)

    return resp.data
})

export default TopicDetailSlice;
export const TopicDetailReducer = TopicDetailSlice.reducer
export const TopicDetailAction = TopicDetailSlice.actions