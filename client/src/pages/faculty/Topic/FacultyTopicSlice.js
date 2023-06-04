import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const FacultyTopicSlice = createSlice({
    name: "departmenttopic",
    initialState: {
        current: {},
        done: {}
    },
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(fetchFacultyCurentTopics.pending, (state, action) => {
                state.current.status = "loading"
            })
            .addCase(fetchFacultyCurentTopics.fulfilled, (state, action) => {
                state.current.status = "idle"
                state.current.topics = action.payload.topic
                state.current.total = action.payload.total
            })
            .addCase(updateStatus.pending, (state, action) => {
                state.current.status = "loading"
            })
            .addCase(updateStatus.fulfilled, (state, action) => {
                state.current.status = "idle"
                state.current.topics = action.payload.topic
                state.current.total = action.payload.total
            })
            .addCase(fetchFacultyDoneTopics.pending, (state, action) => {
                state.done.status = "loading"
            })
            .addCase(fetchFacultyDoneTopics.fulfilled, (state, action) => {
                state.done.status = "idle"
                state.done.topics = action.payload.topic
                state.done.total = action.payload.total
            })
    }
})

export const fetchFacultyCurentTopics = createAsyncThunk("fetchFacultyCurentTopics", async (payload) => {
    const resp = await client.get("/api/topics", {
        params: {
            facultyID: payload.facultyID,
            departmentID: payload.departmentID ? payload.departmentID : 0,
            search: payload.search ? payload.search : "",
            status: payload.status ? payload.status : 0,
            limit: payload.limit ? payload.limit : 0,
            offset: payload.offset ? payload.offset : 0,
            isCurrent: 1
        }
    })

    return resp.data
})
export const fetchFacultyDoneTopics = createAsyncThunk("fetchFacultyDoneTopics", async (payload) => {
    const resp = await client.get("/api/topics", {
        params: {
            facultyID: payload.facultyID,
            departmentID: payload.departmentID ? payload.departmentID : 0,
            search: payload.search ? payload.search : "",
            status: payload.status ? payload.status : 0,
            limit: payload.limit ? payload.limit : 0,
            offset: payload.offset ? payload.offset : 0,
            isCurrent: 2
        }
    })

    return resp.data
})

export const updateStatus = createAsyncThunk("updateStatus", async (payload) => {
    await client.put(`/api/topic-status/${payload.id}`, { status: payload.status })

    const resp = await client.get("/api/topics", {
        params: {
            facultyID: payload.facultyID,
            departmentID: payload.departmentID ? payload.departmentID : 0,
            isCurrent: 1
        }
    })

    return resp.data
})

export default FacultyTopicSlice;
export const FacultyTopicReducer = FacultyTopicSlice.reducer;
export const FacultyTopicAction = FacultyTopicSlice.actions;