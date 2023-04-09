import { useDispatch, useSelector } from "react-redux";
import { Link, useNavigate } from "react-router-dom";

import "./login.css";
import { fetchLoginUser } from "./LoginSlice";
import { loginSelector } from "../../store/selectors";
import { useEffect } from "react";

const Login = () => {
  const dispatch = useDispatch();
  const navige = useNavigate();
  const isAuthenticated = useSelector(loginSelector).user.name !== undefined;
  const handleLoginSubmit = (e) => {
    e.preventDefault();
    const payload = {
      username: e.target[0].value,
      password: e.target[1].value
    }
    
    dispatch(fetchLoginUser(payload))
  };

  useEffect(() => {
    if (isAuthenticated) {
      navige("/")
    }
  }, [isAuthenticated, navige])

  return (
    <div className="loginContainer">
      <form className="login__form" onSubmit={handleLoginSubmit}>
        <div className="login__header">
          <h1>Tài khoản</h1>
        </div>
        <div className="login__input">
          <input type="text" name="username" placeholder="Tên đăng nhập" />
          <span className="err"></span>
        </div>
        <div className="login__input">
          <input type="password" name="password" placeholder="Mật khẩu" />
          <span className="err"></span>
        </div>
        <button className="login__submitBtn" type="submit">Đăng nhập</button>
        <Link to="/">Trở về trang chủ</Link>
      </form>
    </div>
  );
};

export default Login;
