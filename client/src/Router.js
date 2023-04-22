import { PublicHome, PublicLogin, PublicTopic, PublicTopicRegis, PublicContact } from "./pages/common";
import { StudentMyTopic, StudentResult, StudentTopic, StudentTopicRegis } from "./pages/student";
import { LectureMyTopic, LectureReport, LectureResult, LectureTopic, LectureTopicRegis } from "./pages/lecture";
import { FacultyResult, FacultySubcommitte, FacultyTopic, FacultyTopicRegis } from "./pages/faculty";
import { AdminCongress, AdminNotification, AdminResult, AdminSubcommittee, AdminTopic } from "./pages/admin";
import { LoginLayout } from "./components/Layout";

const publicRoutes = [
  { path: "/", component: PublicHome },
  { path: "/home", component: PublicHome },
  { path: "/login", component: PublicLogin, Layout: LoginLayout},
  { path: "/contact", component: PublicContact},
];

const privateRoutes = [
  // user dont login
  { path: "/topic", component: PublicTopic, role: 0},
  { path: "/topic-regis", component: PublicTopicRegis, role: 0},
  // student routes
  { path: "/topic", component: StudentTopic, role: 1},
  { path: "/my-topic", component: StudentMyTopic, role: 1},
  { path: "/topic-regis", component: StudentTopicRegis, role: 1},
  { path: "/result", component: StudentResult, role: 1},

  // lecture routes
  { path: "/topic", component: LectureTopic, role: 2},
  { path: "/my-topic", component: LectureMyTopic, role: 2},
  { path: "/topic-regis", component: LectureTopicRegis, role: 2},
  { path: "/result", component: LectureResult, role: 2},
  { path: "/report", component: LectureReport, role: 2},

  // faculty routes
  { path: "/topic", component: FacultyTopic, role: 4},
  { path: "/topic-regis", component: FacultyTopicRegis, role: 4},
  { path: "/result", component: FacultyResult, role: 4},
  { path: "/subcommittee", component: FacultySubcommitte, role: 4},

  // admin routes
  { path: "/topic", component: AdminTopic, role: 5},
  { path: "/result", component: AdminResult, role: 5},
  { path: "/notification", component: AdminNotification, role: 5},
  { path: "/congress", component: AdminCongress, role: 5},
  { path: "/subcommittee", component: AdminSubcommittee, role: 5},
];

export {
  publicRoutes,
  privateRoutes,
};
