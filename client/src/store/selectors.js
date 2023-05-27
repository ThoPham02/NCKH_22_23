// common selector
export const loginSelector = (state) => state.login;
export const userSelector = (state) => state.login.user;
export const roleSelector = (state) => state.login.user.role;
export const tokenSelector = (state) => state.login.token;
export const departmentSelector = (state) => state.department.departments
export const facultySelector = (state) => state.faculty.faculties

export const CommonEventSelector = (state) => state.CommonEventStore

// admin selector
export const AdminEventSelector = (state) => state.AdminEventStore

// lecture selector
export const LectureTopicSelector = (state) => state.LectureTopicStore








// topic
export const topicsSelector = (state) => state.topic.topics
export const topicSelector = (state) => state.topic
export const resultSelector = (state) => state.topic.result
export const statusSelector = (state) => state.topic.status

//topic detail
export const topicDetailSelector = (state) => state.topicDetail

//lecture topic
export const currentEventSelector = (state) => state.lectureTopic.currentEvent
export const myTopicSelector = state => state.lectureTopic.topic

//stage 
export const stageSelector = (state) => state.stage