import { AdminLayout, LoginLayout } from "./components/Layout";
import { PublicHome, PublicLogin, PublicTopic, PublicResult, PublicContact } from "./pages/common";
import { StudentTopic } from "./pages/student"; // StudentMyTopic, StudentResult,
// import { LectureMyTopic, LectureResult, LectureTopic } from "./pages/lecture";
import { FacultyDashBoard} from "./pages/faculty"; //, FacultyReport, FacultyResult, FacultyTopic 
import { AdminDashBoard} from "./pages/admin"; // , AdminReport, AdminResult, AdminTopic 

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
  { path: "/action", component: PublicTopic, role: 1},
  { path: "/result", component: PublicTopic, role: 1},
  { path: "/report", component: PublicTopic, role: 1},

  // lecture routes
  { path: "/topic", component: PublicTopic, role: 2},
  { path: "/action", component: PublicTopic, role: 2},
  { path: "/result", component: PublicTopic, role: 2},
  { path: "/report", component: PublicTopic, role: 2},

  // faculty routes
  { path: "/admin-home", component: FacultyDashBoard, role: 4, Layout: AdminLayout},
  { path: "/admin-topic", component: PublicTopic, role: 4, Layout: AdminLayout},
  { path: "/admin-result", component: PublicTopic, role: 4, Layout: AdminLayout},
  { path: "/admin-report", component: PublicTopic, role: 4, Layout: AdminLayout},

  // admin routes
  { path: "/admin-home", component: AdminDashBoard, role: 5, Layout: AdminLayout},
  { path: "/admin-topic", component: PublicTopic, role: 5, Layout: AdminLayout},
  { path: "/admin-result", component: PublicTopic, role: 5, Layout: AdminLayout},
  { path: "/admin-report", component: PublicTopic, role: 5, Layout: AdminLayout},
];

export {
  publicRoutes,
  privateRoutes,
};
