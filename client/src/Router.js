import Home from "./pages/Home";

const publicRoutes = [
  { path: "/", component: Home },
  { path: "/home", component: Home },
];

const privateRoutes = [];

export {
  publicRoutes,
  privateRoutes,
};
