import { AdminLayout, LoginLayout } from "./components/Layout";
import { PublicHome, PublicLogin, PublicTopic, PublicResult, PublicContact } from "./pages/common";
import { StudentMyTopic, StudentResult, StudentTopic } from "./pages/student";
import { LectureMyTopic, LectureResult, LectureTopic } from "./pages/lecture";
import { FacultyDashBoard, FacultyEvent, FacultyReport, FacultyResult, FacultyTopic } from "./pages/faculty";
import { AdminDashBoard, AdminEvent, AdminReport, AdminResult, AdminTopic } from "./pages/admin";

const publicRoutes = [
  { path: "/", component: PublicHome },
  { path: "/home", component: PublicHome },
  { path: "/login", component: PublicLogin, Layout: LoginLayout},
  { path: "/contact", component: PublicContact},
];

const privateRoutes = [
  // user dont login
  { path: "/topic", component: PublicTopic, role: 0},
  { path: "/result", component: PublicResult, role: 0},
  // student routes
  { path: "/topic", component: StudentTopic, role: 1},
  { path: "/action", component: StudentMyTopic, role: 1},
  { path: "/result", component: StudentTopic, role: 1},
  { path: "/report", component: StudentResult, role: 1},

  // lecture routes
  { path: "/topic", component: LectureTopic, role: 2},
  { path: "/action", component: LectureMyTopic, role: 2},
  { path: "/result", component: LectureResult, role: 2},
  { path: "/report", component: LectureResult, role: 2},

  // faculty routes
  { path: "/admin-home", component: FacultyDashBoard, role: 4, Layout: AdminLayout},
  { path: "/admin-topic", component: FacultyTopic, role: 4, Layout: AdminLayout},
  { path: "/admin-result", component: FacultyResult, role: 4, Layout: AdminLayout},
  { path: "/admin-event", component: FacultyEvent, role: 4, Layout: AdminLayout},
  { path: "/admin-report", component: FacultyReport, role: 4, Layout: AdminLayout},

  // admin routes
  { path: "/admin-home", component: AdminDashBoard, role: 5, Layout: AdminLayout},
  { path: "/admin-topic", component: AdminTopic, role: 5, Layout: AdminLayout},
  { path: "/admin-result", component: AdminResult, role: 5, Layout: AdminLayout},
  { path: "/admin-event", component: AdminEvent, role: 5, Layout: AdminLayout},
  { path: "/admin-report", component: AdminReport, role: 5, Layout: AdminLayout},
];

export {
  publicRoutes,
  privateRoutes,
};
