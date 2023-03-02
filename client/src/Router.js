import Home from "./pages/Home";
import Topic from "./pages/Topic"
const publicRoutes = [
  { path: "/", component: Home },
  { path: "/home", component: Home },
  { path: "/topic", component: Topic}
];

const privateRoutes = [];

export {
  publicRoutes,
  privateRoutes,
};
