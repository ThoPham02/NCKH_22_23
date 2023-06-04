import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../apis";

const StudentMyTopicSlice = createSlice({
    name: "studentmytopic",
    initialState: {
        current: {},
        done: {},
    },
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(fetchStudentMyTopic.pending, (state, action) => {
                state.status = "loading"
            })
            .addCase(fetchStudentMyTopic.fulfilled, (state, action) => {
                state.status = "idle"
                state.current.topics = action.payload.current.topic
                state.done.topics = action.payload.done.topic
            })
    }
})

export const fetchStudentMyTopic = createAsyncThunk("fetchStudentMyTopic", async (payload) => {
    const resp = await client.get("/api/topics", {
        params: {
            studentID: payload.studentID,
            isCurrent: 1
        }
    })
    const resp2 = await client.get("/api/topics", {
        params: {
            studentID: payload.studentID,
            isCurrent: 2
        }
    })
    return { current: resp.data, done: resp2.data }
})

export default StudentMyTopicSlice;
export const StudentMyTopicReducer = StudentMyTopicSlice.reducer
export const StudentMyTopicAction = StudentMyTopicSlice.actions