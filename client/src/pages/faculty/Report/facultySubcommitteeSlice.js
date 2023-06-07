import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const FacultySubcommitteeSLice = createSlice({
    name: "faculty_subcommittee",
    initialState: {
        status: "idle"
    },
    reducers: {

    },
    extraReducers: builder => {
        builder
        .addCase(fetchSubcommittee.pending, (state, action) => {
            state.status = "loading"
        })
        .addCase(fetchSubcommittee.fulfilled, (state, action) => {
            state.status = "idle"
            state.subcommittee = action.payload.subcommittees
        })
        .addCase(addSubcommittee.pending, (state, action) => {
            state.status = "loading"
        })
        .addCase(addSubcommittee.fulfilled, (state, action) => {
            state.status = "idle"
            state.subcommittee = action.payload.subcommittees
        })
        .addCase(fetchTopicsBySubcommittee.pending, (state, action) => {
            state.status = "loading"
        })
        .addCase(fetchTopicsBySubcommittee.fulfilled, (state, action) => {
            state.status = "idle"
            state.topics = action.payload.topic
        } )
        .addCase(addTopicsToSub.pending, (state, action) => {
            state.status = "loading"
        })
        .addCase(addTopicsToSub.fulfilled, (state, action) => {
            state.status = "idle"
            state.topics = action.payload.topic
            state.suggestTopics = action.payload.topic.filter(item => item.subcommitteeID === 0)
        } )
    }
})

export const fetchSubcommittee = createAsyncThunk("fetchSubcommittee", async () => {
    const resp = await client.get("/api/subcommittees")

    return resp.data
})

export const fetchTopicsBySubcommittee = createAsyncThunk("fetchTopicsBySubcommittee", async (payload) => {
    const resp = await client.get("/api/topics", {
        params: {
            isCurrent: 1,
            status: 128,
            facultyID: payload.facultyID,
            search: payload.search,
            subcommitteeID: payload.subcommitteeID
        }
    })

    return resp.data
})

export const addSubcommittee = createAsyncThunk("addSubcommittee", async (payload) => {
    await client.post("/api/subcommittee", {
        name: payload.name,
        listLectures: payload.listLectures,
        facultyID: payload.facultyID,
        eventID: payload.eventID
    })

    const resp = await client.get("/api/subcommittees")

    return resp.data
})

export const addTopicsToSub = createAsyncThunk("addTopicsToSub", async (payload) => {
    await client.put("/api/topic-subcommittee", {
        subcommitteeID: payload.id,
        listTopicID: payload.listTopicID,
    })

    const resp = await client.get("/api/topics", {
        params: {
            isCurrent: 1,
            status: 128,
            facultyID: payload.facultyID,
        }
    })

    return resp.data
})


export default FacultySubcommitteeSLice
export const FacultySubcommitteeReducer = FacultySubcommitteeSLice.reducer