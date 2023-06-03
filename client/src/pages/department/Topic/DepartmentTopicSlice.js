import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const DepartmentTopicSlice = createSlice({
    name: "departmenttopic",
    initialState: {
        gender: {},
        regis: {},
        suggest: {},
    },
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(fetchTopics.pending, (state, action) => {
                state.gender.status = "loading"
            })
            .addCase(fetchTopics.fulfilled, (state, action) => {
                state.gender.status = "idle"
                state.gender.topics = action.payload.topic
                state.gender.total = action.payload.total
            })
            .addCase(updateStatus.pending, (state, action) => {
                state.status = "loading"
            })
            .addCase(updateStatus.fulfilled, (state, action) => {
                state.status = "idle"
                state.topics = action.payload.topic
                state.total = action.payload.total
            })
    }
})

export const fetchTopics = createAsyncThunk("fetchTopics", async (payload) => {
    const resp = await client.get("/api/topics", {
        params: {

        }
    })

    return resp.data
})

export const updateStatus = createAsyncThunk("updateStatus", async (payload) => {
    await client.put(`/api/topic-status/${payload.id}`, { status: 4 })

    const resp = await client.get("/api/topics", {
        params: {
            
        }
    })

    return resp.data
})

export default DepartmentTopicSlice;
export const DepartmentTopicReducer = DepartmentTopicSlice.reducer;
export const DepartmentTopicAction = DepartmentTopicSlice.actions;