import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.min.css';
import { publicRoutes, privateRoutes } from "./Router";
import "./App.css";
import DefaultLayout from "./components/Layout/DefaultLayout";
import { useSelector } from "react-redux";
import { loginSelector } from "./store/selectors";

function App() {
  const user = useSelector(loginSelector)
  console.log(user)
  return (
    <Router>
      <Routes>
        {publicRoutes.map((route, index) => {
          const Layout = route.Layout || DefaultLayout;
          const Page = route.component;
          return (
            <Route
              key={index}
              path={route.path}
              element={
                <Layout>
                  <Page />
                </Layout>
              }
            />
          );
        })}
        {privateRoutes.map((route, index) => {
          if (user.permission === undefined) {
            // TODO: redirect to login page
          }
          const Layout = route.Layout || DefaultLayout;
          const Page = route.component;
          return (
            <Route
              key={index}
              path={route.path}
              element={
                <Layout>
                  <Page />
                </Layout>
              }
            />
          );
        })}
      </Routes>
    </Router>
  );
}

export default App;
