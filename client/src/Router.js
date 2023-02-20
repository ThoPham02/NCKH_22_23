import Home from "./pages/Home";

const publicRoutes = [
    { path: "/", component: Home },
    { path: "/home", component: Home },
];

const adminRoutes = [];
const studentRoutes = [];
const lecturerRoutes = [];
const departmentRoutes = [];
const facultyRoutes = [];

export {
  publicRoutes,
  adminRoutes,
  studentRoutes,
  lecturerRoutes,
  departmentRoutes,
  facultyRoutes,
};
