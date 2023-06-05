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
    }
})

export const fetchSubcommittee = createAsyncThunk("fetchSubcommittee", async () => {
    const resp = await client.get("/api/subcommittees")

    return resp.data
})

export default FacultySubcommitteeSLice
export const FacultySubcommitteeReducer = FacultySubcommitteeSLice.reducer