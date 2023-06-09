import { AdminLayout, LoginLayout } from "./components/Layout";
import { PublicHome, PublicLogin, PublicTopic, PublicResult, PublicContact } from "./pages/common";
import { FacultyDashBoard, FacultyReport, FacultyTopic } from "./pages/faculty";
import { AdminDashBoard, AdminReport, AdminTopic } from "./pages/admin";
import { StudentMyTopic, StudentSchoolReport, StudentSubcommitteeReport } from "./pages/student";
import { LectureMark, LectureMyTopic, LectureSchoolReport, LectureSubcommitteeReport } from "./pages/lecture";
import { DepartmentReport, DepartmentTopic } from "./pages/department";

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
  { path: "/admin-home", component: PublicHome, role: 0 },
  { path: "/admin-topic", component: PublicHome, role: 0 },
  { path: "/admin-result", component: PublicHome, role: 0 },
  { path: "/admin-report", component: PublicHome, role: 0 },
  { path: "/topics", component: PublicHome, role: 0 },
  { path: "/my-topic", component: PublicHome, role: 0 },
  { path: "/subcommittee-report", component: PublicHome, role: 0 },
  { path: "/school-report", component: PublicHome, role: 0 },
  { path: "/mark", component: PublicHome, role: 0 },
  { path: "/result", component: PublicHome, role: 0 },
  // student routes
  { path: "/topics", component: PublicTopic, role: 1 },
  { path: "/my-topic", component: StudentMyTopic, role: 1 },
  { path: "/subcommittee-report", component: StudentSubcommitteeReport, role: 1 },
  { path: "/school-report", component: StudentSchoolReport, role: 1 },
  { path: "/result", component: PublicResult, role: 1 },

  // lecture routes
  { path: "/topics", component: PublicTopic, role: 2 },
  { path: "/my-topic", component: LectureMyTopic, role: 2 },
  { path: "/subcommittee-report", component: LectureSubcommitteeReport, role: 2 },
  { path: "/school-report", component: LectureSchoolReport, role: 2 },
  { path: "/mark", component: LectureMark, role: 2 },
  { path: "/result", component: PublicResult, role: 2 },

  // department routes
  { path: "/admin-home", component: AdminDashBoard, role: 3, Layout: AdminLayout },
  { path: "/admin-topic", component: DepartmentTopic, role: 3, Layout: AdminLayout },
  { path: "/admin-result", component: PublicResult, role: 3, Layout: AdminLayout },
  { path: "/admin-report", component: DepartmentReport, role: 3, Layout: AdminLayout },

  // faculty routes
  { path: "/admin-home", component: FacultyDashBoard, role: 4, Layout: AdminLayout },
  { path: "/admin-topic", component: FacultyTopic, role: 4, Layout: AdminLayout },
  { path: "/admin-result", component: PublicResult, role: 4, Layout: AdminLayout },
  { path: "/admin-report", component: FacultyReport, role: 4, Layout: AdminLayout },

  // admin routes
  { path: "/admin-home", component: AdminDashBoard, role: 5, Layout: AdminLayout },
  { path: "/admin-topic", component: AdminTopic, role: 5, Layout: AdminLayout },
  { path: "/admin-result", component: PublicResult, role: 5, Layout: AdminLayout },
  { path: "/admin-report", component: AdminReport, role: 5, Layout: AdminLayout },
];

export {
  publicRoutes,
  privateRoutes,
};
