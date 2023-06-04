import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const DepartmentTopicSlice = createSlice({
    name: "departmenttopic",
    initialState: {
        current: {},
        done: {}
    },
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(fetchDepartmentCurentTopics.pending, (state, action) => {
                state.current.status = "loading"
            })
            .addCase(fetchDepartmentCurentTopics.fulfilled, (state, action) => {
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
    }
})

export const fetchDepartmentCurentTopics = createAsyncThunk("fetchDepartmentCurentTopics", async (payload) => {
    const resp = await client.get("/api/topics", {
        params: {
            departmentID: payload.departmentID,
            search: payload.search ? payload.search : "",
            status: payload.status ? payload.status : 0,
            limit: payload.limit ? payload.limit : 0,
            offset: payload.offset ? payload.offset : 0,
            isCurrent: 1
        }
    })

    return resp.data
})

export const updateStatus = createAsyncThunk("updateStatus", async (payload) => {
    await client.put(`/api/topic-status/${payload.id}`, { status: payload.status })

    const resp = await client.get("/api/topics", {
        params: {
            departmentID: payload.departmentID,
            isCurrent: 1
        }
    })

    return resp.data
})

export default DepartmentTopicSlice;
export const DepartmentTopicReducer = DepartmentTopicSlice.reducer;
export const DepartmentTopicAction = DepartmentTopicSlice.actions;