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
        .addCase()
    }
})

export const fetchSubcommittee = createAsyncThunk("fetchSubcommittee", async () => {
    const resp = await client.get("")

    return resp.data
})