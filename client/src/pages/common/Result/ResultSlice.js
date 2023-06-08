import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const ResultSlice = createSlice({
    name: "result",
    initialState: {
        status: "idle",
        error: null,
        data: [],
    },
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(fetchTopicMarkCurrent.pending, (state, action) => {
                state.status = "loading"
            })
            .addCase(fetchTopicMarkCurrent.fulfilled, (state, action) => {
                state.status = "idle"
                state.current = action.payload.topicMark
                state.total = action.payload.total
            })
            .addCase(fetchTopicMarkDone.pending, (state, action) => {
                state.status = "loading"
            })
            .addCase(fetchTopicMarkDone.fulfilled, (state, action) => {
                state.status = "idle"
                state.done = action.payload.topicMark
                state.total = action.payload.total
            })
    }
})

export const fetchTopicMarkCurrent = createAsyncThunk("fetchTopicMarkCurrent", async (payload) => {
    const resp = await client.get("/api/result/topic-mark", {
        params: {
            search: payload.search ? payload.search : "",
            departmentID: payload.departmentID ? payload.departmentID : 0,
            facultyID: payload.facultyID ? payload.facultyID : 0,
            subcommitteeID: payload.subcommitteeID ? payload.subcommitteeID : 0,
            limit: payload.limit ? payload.limit : 0,
            offset: payload.offset ? payload.offset : 0,
            isCurrent: 1
        }
    })

    return resp.data
})

export const fetchTopicMarkDone = createAsyncThunk("fetchTopicMarkDone", async (payload) => {
    const resp = await client.get("/api/result/topic-mark", {
        params: {
            search: payload.search ? payload.search : "",
            departmentID: payload.departmentID ? payload.departmentID : 0,
            facultyID: payload.facultyID ? payload.facultyID : 0,
            subcommitteeID: payload.subcommitteeID ? payload.subcommitteeID : 0,
            limit: payload.limit ? payload.limit : 0,
            offset: payload.offset ? payload.offset : 0,
            isCurrent: 2
        }
    })

    return resp.data
})

export default ResultSlice;
export const ResultReducer = ResultSlice.reducer;
export const ResultActions = ResultSlice.actions;