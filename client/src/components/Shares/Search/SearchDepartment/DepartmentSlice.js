import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const DepartmentSlice = createSlice({
  name: "department",
  initialState: {
    status: "idle",
    listDepartments: [],
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchDepartment.pending, (state, action) => {
        state.status = "pending";
      })
      .addCase(fetchDepartment.fulfilled, (state, action) => {
        state.status = "idle";
        state.listDepartments = action.payload;
      });
  },
});

export const fetchDepartment = createAsyncThunk(
  "getDepartment",
  async (payload) => {
    let url = "/api/department";
    if (payload.faculityId !== 0) {
      url += "?faculty_id=" + payload.facultyId;
    }
    const response = await client.get(url);

    return response.data;
  }
);

export const DepartmentReducer = DepartmentSlice.reducer;

export default DepartmentSlice;
