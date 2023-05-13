import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const DepartmentSlice = createSlice({
  name: "department",
  initialState: {
    status: "idle",
    departments: [],
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchDepartment.pending, (state, action) => {
        state.status = "pending";
      })
      .addCase(fetchDepartment.fulfilled, (state, action) => {
        state.status = "idle";
        if (action.payload.departments !== null) {
          state.departments = action.payload.departments;
        }
      });
  },
});

export const fetchDepartment = createAsyncThunk("getDepartment", async () => {
  const response = await client.get("/api/departments");

  return response.data;
});

export const DepartmentReducer = DepartmentSlice.reducer;
export default DepartmentSlice;
