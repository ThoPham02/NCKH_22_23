import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { useSelector } from "react-redux";
import React from "react";

import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";
import { publicRoutes, privateRoutes } from "./Router";
import { userSelector } from "./store/selectors";
import { DefaultLayout } from "./components/Layout";

function App() {
  const user = useSelector(userSelector);
  const role = user.typeAccount;
  return (
    <Router>
      <Routes>
        {role === 0
          ? publicRoutes.map((route, index) => {
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
            })
          : privateRoutes
              .filter((route) => route.role === role)
              .map((route, index) => {
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
