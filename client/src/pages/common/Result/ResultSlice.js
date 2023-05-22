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
        builder.addCase(fetchResult.pending, (state, action) => {

        })
    }
})


const fetchResult = createAsyncThunk("fetchResult", async () => {
    const resp = await client.get()

    return resp.json()
})