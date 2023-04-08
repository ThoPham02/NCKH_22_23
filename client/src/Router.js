import Home from "./pages/Home";
import Topic from "./pages/Topic"
import Login from "./pages/Login"
import LoginLayout from "./components/Layout/LoginLayout";
import Category from "./pages/Category";

const publicRoutes = [
  { path: "/", component: Home },
  { path: "/home", component: Home },
  { path: "/topic", component: Topic},
  { path: "/login", component: Login, Layout: LoginLayout},
  { path: "/statistical", component: Topic},
  { path: "/category", component: Category},
  { path: "/contact", component: Topic},
];

const privateRoutes = [
  { path: "/statistical", component: Topic},
];

export {
  publicRoutes,
  privateRoutes,
};
