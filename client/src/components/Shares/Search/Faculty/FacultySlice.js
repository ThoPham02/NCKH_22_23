import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const FacultySlice = createSlice({
  name: "faculty",
  initialState: {
    status: "idle",
    faculties: [],
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchFaculty.pending, (state, action) => {
        state.status = "pending";
      })
      .addCase(fetchFaculty.fulfilled, (state, action) => {
        state.status = "idle";
        if (action.payload.faculties !== null) {
          state.faculties = action.payload.faculties;
        }
      });
  },
});

export const fetchFaculty = createAsyncThunk("getFaculities", async () => {
  const response = await client.get("/api/faculties");

  return response.data;
});

export const FacultyReducer = FacultySlice.reducer;
export default FacultySlice;
