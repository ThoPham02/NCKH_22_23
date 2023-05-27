import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const EventSlice = createSlice({
    name: "admin-event",
    initialState: {},
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(fetchEvents.pending, (state, action) => {
                state.status = "loading"
            })
            .addCase(fetchEvents.fulfilled, (state, action) => {
                state.status = "idle"
                state.current = action.payload.event
                state.result = action.payload.result
            })
            .addCase(updateStage.pending, (state, action) => {
                state.status = "loading"
            })
            .addCase(updateStage.fulfilled, (state, action) => {
                state.status = "idle"
                state.current = action.payload.event
                state.result = action.payload.result
            })
            .addCase(createEvent.pending, (state, action) => {
                state.status = "loading"
            })
            .addCase(createEvent.fulfilled, (state, action) => {
                state.status = "idle"
                state.current = action.payload.event
                state.result = action.payload.result
            })
    }
})

export const fetchEvents = createAsyncThunk("fetchEvents", async () => {
    const resp = await client.get(`/api/event-current`)

    return resp.data
})

export const fetchDoneEvents = createAsyncThunk("fetchDoneEvents", async () => {
    const resp = await client.get(``)

    return resp.data
})

export const updateStage = createAsyncThunk("updateStage", async (payload) => {
    await client.put(`/api/stage/${payload.stageID}`, {
        description: payload.description,
        url: "",
        timeStart: payload.timeStart,
        timeEnd: payload.timeEnd
    })

    const fetch = await client.get(`/api/event-current`)

    return fetch.data
})

export const createEvent = createAsyncThunk("createEvent", async (payload) => {
    await client.post(`/api/event`, {
        name: payload.name,
        schoolYear: payload.schoolYear
    })

    const fetch = await client.get(`/api/event-current`)

    return fetch.data
})

export default EventSlice
export const AdminEventReducer = EventSlice.reducer
export const AdminEventActions = EventSlice.actions