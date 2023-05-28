import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const ShareTopicDetailSlice = createSlice({
    name: "share-topic-detail",
    initialState: {
        status: "idle"
    },
    reducers: {},
    extraReducers: builder => {
        builder
        .addCase(fetchTopicDetail.pending, (state, action) => {
            state.status = "loading";
        }).addCase(fetchTopicDetail.fulfilled, (state, action) => {
            state.status = "idle";
            state.topic = action.payload.topic;
            state.event = action.payload.event;
            state.subcommittee = action.payload.subcommittee
            state.reports = action.payload.reports
            state.marks = action.payload.marks
            state.listStudent = action.payload.listStudent
        })
    }
})

export const fetchTopicDetail = createAsyncThunk("fetchTopicDetail", async (payload) => {
    const resp = await client.get(`/api/topic/${payload.id}`)
    
    return resp.data
})

export default ShareTopicDetailSlice;
export const ShareTopicDetailReducer = ShareTopicDetailSlice.reducer;
export  const ShareTopicDetailAction = ShareTopicDetailSlice.actions;