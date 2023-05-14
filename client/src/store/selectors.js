// select login info
export const loginSelector = (state) => state.login;
export const userSelector = (state) => state.login.user;
export const roleSelector = (state) => state.login.user.role;
export const tokenSelector = (state) => state.login.token;

// department
export const departmentSelector = (state) => state.department.departments

// faculty
export const facultySelector = (state) => state.faculty.faculties


// topic
export const topicsSelector = (state) => state.topic.topics
export const topicSelector = (state) => state.topic