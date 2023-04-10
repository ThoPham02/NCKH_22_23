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
        state.listDepartments = action.payload.list_department;
      });
  },
});

export const fetchDepartment = createAsyncThunk(
  "getDepartment",
  async (_, {getState}) => {
    const state = getState()
    let url = "/api/department?faculty_id=" + state.topicFilter.faculity
    const response = await client.get(url);

    return response.data;
  }
);

export const DepartmentReducer = DepartmentSlice.reducer;

export default DepartmentSlice;
