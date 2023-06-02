export const loginSelector = (state) => state.login;
export const userSelector = (state) => state.login.user;
export const roleSelector = (state) => state.login.user.role;
export const tokenSelector = (state) => state.login.token;
export const departmentSelector = (state) => state.department.departments
export const facultySelector = (state) => state.faculty.faculties
// share selector
export const TopicDetailSelector = state => state.TopicDetailStore

// common selector
export const CommonEventSelector = (state) => state.CommonEventStore
export const CommonTopicSelector = (state) => state.CommonTopicStore

// admin selector
export const AdminEventSelector = (state) => state.AdminEventStore

// lecture selector
export const LectureTopicSelector = (state) => state.LectureTopicStore

