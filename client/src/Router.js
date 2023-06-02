import { AdminLayout, LoginLayout } from "./components/Layout";
import { PublicHome, PublicLogin, PublicTopic, PublicResult, PublicContact } from "./pages/common";
import { FacultyDashBoard } from "./pages/faculty";
import { AdminDashBoard } from "./pages/admin";
import { StudentMyTopic, StudentResult, StudentSchoolReport, StudentSubcommitteeReport, StudentTopics } from "./pages/student";
import { LectureMark, LectureMyTopic, LectureResult, LectureSchoolReport, LectureSubcommitteeReport, LectureTopic } from "./pages/lecture";

const publicRoutes = [
  { path: "/", component: PublicHome },
  { path: "/home", component: PublicHome },
  { path: "/login", component: PublicLogin, Layout: LoginLayout },
  { path: "/contact", component: PublicContact },
];

const privateRoutes = [
  // user dont login
  { path: "/topics", component: PublicTopic, role: 0 },
  { path: "/result", component: PublicResult, role: 0 },
  // student routes
  { path: "/topics", component: StudentTopics, role: 1 },
  { path: "/my-topic", component: StudentMyTopic, role: 1 },
  { path: "/subcommittee-report", component: StudentSubcommitteeReport, role: 1 },
  { path: "/school-report", component: StudentSchoolReport, role: 1 },
  { path: "/result", component: StudentResult, role: 1 },

  // lecture routes
  { path: "/topics", component: LectureTopic, role: 2 },
  { path: "/my-topic", component: LectureMyTopic, role: 2 },
  { path: "/subcommittee-report", component: LectureSubcommitteeReport, role: 2 },
  { path: "/school-report", component: LectureSchoolReport, role: 2 },
  { path: "/mark", component: LectureMark, role: 2 },
  { path: "/result", component: LectureResult, role: 2 },
  
  // department routes
  { path: "/admin-home", component: AdminDashBoard, role: 3, Layout: AdminLayout },
  { path: "/admin-topics", component: PublicTopic, role: 3, Layout: AdminLayout },
  { path: "/admin-my-topics", component: PublicTopic, role: 3, Layout: AdminLayout },
  { path: "/admin-result", component: PublicTopic, role: 3, Layout: AdminLayout },
  { path: "/admin-report", component: PublicTopic, role: 3, Layout: AdminLayout },

  // faculty routes
  { path: "/admin-home", component: FacultyDashBoard, role: 4, Layout: AdminLayout },
  { path: "/admin-topic", component: PublicTopic, role: 4, Layout: AdminLayout },
  { path: "/admin-result", component: PublicTopic, role: 4, Layout: AdminLayout },
  { path: "/admin-report", component: PublicTopic, role: 4, Layout: AdminLayout },

  // admin routes
  { path: "/admin-home", component: AdminDashBoard, role: 5, Layout: AdminLayout },
  { path: "/admin-topic", component: PublicTopic, role: 5, Layout: AdminLayout },
  { path: "/admin-result", component: PublicTopic, role: 5, Layout: AdminLayout },
  { path: "/admin-report", component: PublicTopic, role: 5, Layout: AdminLayout },
];

export {
  publicRoutes,
  privateRoutes,
};
