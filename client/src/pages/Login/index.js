import { Link } from "react-router-dom";

import client from "../../apis";
import "./login.css";

const Login = () => {
  const handleLoginSubmit = (e) => {
    e.preventDefault();

    client
      .post("/user/login", {
        username: e.target.username.value,
        password: e.target.password.value,
      })
      .then((response) => console.log(response.data))
      .catch((error) => console.error(error));
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
