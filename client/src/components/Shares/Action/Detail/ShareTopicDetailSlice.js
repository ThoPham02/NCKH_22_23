import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const ShareTopicDetailSlice = createSlice({
    name: "share-topic-detail",
    initialState: {
        status: "idle"
    },
    reducers: {
        clearStore: (state, action) => {
            state = {status: "idle"}
        }
    },
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
        .addCase(addStudentGroup.pending, (state, action) => {
            state.status = "loading";
        }).addCase(addStudentGroup.fulfilled, (state, action) => {
            state.status = "idle";
            state.topic = action.payload.topic;
            state.event = action.payload.event;
            state.subcommittee = action.payload.subcommittee
            state.reports = action.payload.reports
            state.marks = action.payload.marks
            state.listStudent = action.payload.listStudent
        })
        .addCase(deleteStudentGroup.pending, (state, action) => {
            state.status = "loading";
        }).addCase(deleteStudentGroup.fulfilled, (state, action) => {
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

export const addStudentGroup = createAsyncThunk("addStudentGroup", async (payload) => {
    await client.put(`/api/topic-student-group/${payload.topicID}`, {
        studentID: payload.studentID
    })

    const resp = await client.get(`/api/topic/${payload.topicID}`)
    
    return resp.data
})

export const deleteStudentGroup = createAsyncThunk("deleteStudentGroup", async (payload) => {
    await client.delete(`/api/topic-student-group/${payload.studentID}`)

    const resp = await client.get(`/api/topic/${payload.topicID}`)
    
    return resp.data
})

export default ShareTopicDetailSlice;
export const ShareTopicDetailReducer = ShareTopicDetailSlice.reducer;
export  const ShareTopicDetailAction = ShareTopicDetailSlice.actions;