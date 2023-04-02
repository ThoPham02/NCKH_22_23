import { useDispatch } from "react-redux";
import { Link } from "react-router-dom";

import "./login.css";
import { fetchLoginUser } from "./LoginSlice";

const Login = () => {
  const dispatch = useDispatch()

  const handleLoginSubmit = (e) => {
    e.preventDefault();
    const payload = {
      username: e.target[0].value,
      password: e.target[1].value
    }
    dispatch(fetchLoginUser(payload))
  };

  

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
          <input type="text" name="password" placeholder="Mật khẩu" />
          <span className="err"></span>
        </div>
        <button className="login__submitBtn" type="submit">Đăng nhập</button>
        <Link to="/">Trở về trang chủ</Link>
      </form>
    </div>
  );
};

export default Login;
