import { configureStore } from "@reduxjs/toolkit";
import { LoginReducer } from "../pages/common/Login/LoginSlice";
import { DepartmentReducer } from "../components/Shares/Search/Department/DepartmentSlice";
import { FacultyReducer } from "../components/Shares/Search/Faculty/FacultySlice";
import { AdminEventReducer } from "../pages/admin/DashBoard/EventSlice";
import { LectureTopicReducer } from "../pages/lecture/Topic/LectureTopicSlice";
import { CommonEventReducer } from "../components/Shares/Search/Event/CommonEventSlice";
import { CommonTopicReducer } from "../pages/common/Topic/CommonTopicSlice";
import { TopicDetailReducer } from "../components/Shares/Action/Detail/TopicDetailSlice";

const store = configureStore({
  reducer: {
    //share reducer
    department: DepartmentReducer,
    faculty: FacultyReducer,
    TopicDetailStore: TopicDetailReducer,


    // common reducer
    login: LoginReducer,
    CommonEventStore: CommonEventReducer,
    CommonTopicStore: CommonTopicReducer, 

    // admin reducer
    AdminEventStore: AdminEventReducer,

    //lecture reducer
    LectureTopicStore: LectureTopicReducer
  },
});

export default store;
