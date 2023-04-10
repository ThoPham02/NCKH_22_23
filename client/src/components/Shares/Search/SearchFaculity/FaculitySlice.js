import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const FaculitySlice = createSlice({
  name: "faculity",
  initialState: {
    status: "idle",
    listFaculities: [],
  },
  reducers: {
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchFaculity.pending, (state, action) => {
        state.status = "pending";
      })
      .addCase(fetchFaculity.fulfilled, (state, action) => {
        state.status = "idle";
        state.listFaculities = action.payload;
      });
  },
});

export const fetchFaculity = createAsyncThunk(
  "getFaculities",
  async () => {
    const response = await client.get("/api/faculity");

    return response.data;
  }
);

export const FaculityReducer = FaculitySlice.reducer;

export default FaculitySlice;