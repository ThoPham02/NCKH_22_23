import { configureStore } from "@reduxjs/toolkit";
import { LoginReducer } from "../pages/common/Login/LoginSlice";
import { DepartmentReducer } from "../components/Shares/Search/Department/DepartmentSlice";
import { FacultyReducer } from "../components/Shares/Search/Faculty/FacultySlice";
import { TopicReducer } from "../pages/common/Topic/TopicSlice";
import { TopicDetailReducer } from "../components/Shares/Action/Detail/TopicDetailSlice";

const store = configureStore({
  reducer: {
    login: LoginReducer,

    department: DepartmentReducer,
    faculty: FacultyReducer,
    topic: TopicReducer,
    topicDetail: TopicDetailReducer
  },
});

export default store;
