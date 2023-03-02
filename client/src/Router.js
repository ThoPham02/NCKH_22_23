import Home from "./pages/Home";
import Topic from "./pages/Topic"
import Login from "./pages/Login"
import LoginLayout from "./components/Layout/LoginLayout";

const publicRoutes = [
  { path: "/", component: Home },
  { path: "/home", component: Home },
  { path: "/topic", component: Topic},
  { path: "/login", component: Login, Layout: LoginLayout}
];

const privateRoutes = [];

export {
  publicRoutes,
  privateRoutes,
};
