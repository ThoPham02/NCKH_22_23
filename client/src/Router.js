import Home from "./pages/Home";
import Topic from "./pages/Topic"
import Login from "./pages/Login"
import LoginLayout from "./components/Layout/LoginLayout";
import UserInfo from "./pages/UserInfo";

const publicRoutes = [
  { path: "/", component: Home },
  { path: "/home", component: Home },
  { path: "/topic", component: Topic},
  { path: "/login", component: Login, Layout: LoginLayout},
  { path: "/category", component: Topic},
  { path: "/contact", component: Topic},
  { path: "/user-info", component: UserInfo},
];

const privateRoutes = [
  { path: "/student/topic", component: Topic},
  { path: "/student/statistical", component: Topic},
  { path: "/lecture/topic", component: Topic},
  { path: "/lecture/statistical", component: Topic},
];

export {
  publicRoutes,
  privateRoutes,
};
