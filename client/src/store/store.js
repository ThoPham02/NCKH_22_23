import { configureStore } from "@reduxjs/toolkit";
import { LoginReducer } from "../pages/common/Login/LoginSlice";
import { DepartmentReducer } from "../components/Shares/Search/Department/DepartmentSlice";
import { FacultyReducer } from "../components/Shares/Search/Faculty/FacultySlice";
import { AdminEventReducer } from "../pages/admin/DashBoard/EventSlice";
import { LectureTopicReducer } from "../pages/lecture/Topic/LectureTopicSlice";
import { CommonEventReducer } from "../components/Shares/Search/Event/CommonEventSlice";

const store = configureStore({
  reducer: {
    // common reducer
    login: LoginReducer,
    department: DepartmentReducer,
    faculty: FacultyReducer,
    CommonEventStore: CommonEventReducer, 

    // admin reducer
    AdminEventStore: AdminEventReducer,

    //lecture reducer
    LectureTopicStore: LectureTopicReducer
  },
});

export default store;
