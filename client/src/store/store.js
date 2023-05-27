import { configureStore } from "@reduxjs/toolkit";
import { LoginReducer } from "../pages/common/Login/LoginSlice";
import { DepartmentReducer } from "../components/Shares/Search/Department/DepartmentSlice";
import { FacultyReducer } from "../components/Shares/Search/Faculty/FacultySlice";
import { AdminEventReducer } from "../pages/admin/DashBoard/EventSlice";

const store = configureStore({
  reducer: {
    // common reducer
    login: LoginReducer,
    department: DepartmentReducer,
    faculty: FacultyReducer,

    // admin reducer
    AdminEventStore: AdminEventReducer
  },
});

export default store;
