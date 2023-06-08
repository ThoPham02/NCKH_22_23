import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const StudentReportSlice = createSlice({
    name: "student-report",
    initialState: {},
    reducers: {},
    extraReducers: builder => {
        builder
        .addCase()
    }
})

export const fetchTopics = createAsyncThunk("fetchTopics", async () => {
    resp = await client.get("/api/topics", {
        params: {
            
        }
    })
})

export default StudentReportSlice;