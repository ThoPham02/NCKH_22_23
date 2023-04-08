import 'bootstrap/dist/css/bootstrap.min.css';
import "./App.css";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import React from "react";
import { useSelector } from "react-redux";

import DefaultLayout from "./components/Layout/DefaultLayout";
import { publicRoutes, privateRoutes } from "./Router";
import { loginSelector } from "./store/selectors";
import { CompareTime } from './utils/time';


function App() {
  const login = useSelector(loginSelector)
  console.log(CompareTime(login.token.accessExpiredAt))
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
  )
  
}

export default App;
