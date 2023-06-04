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
export const CommonCurrentTopicSelector = (state) => state.CommonTopicStore.current
export const CommonDoneTopicSelector = (state) => state.CommonTopicStore.done

// admin selector
export const AdminEventSelector = (state) => state.AdminEventStore

// lecture selector
export const LectureTopicSelector = (state) => state.LectureTopicStore
export const LectureMyTopicSelector = (state) => state.LectureMyTopicStore

// department selector 
export const DepartmentCurrentTopic = (state) => state.DepartmentTopicStore.current
export const DepartmentDoneTopic = (state) => state.DepartmentTopicStore.done

// faculty selector 
export const FacultyCurrentTopic = (state) => state.FacultyTopicStore.current
export const FacultyDoneTopic = (state) => state.FacultyTopicStore.done

// student selector
export const StudentMyTopicCurrent = (state) => state.StudentMyTopicStore.current
export const StudentMyTopicDone = (state) => state.StudentMyTopicStore.done
