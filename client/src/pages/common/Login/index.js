import { useDispatch, useSelector } from "react-redux";
import { Link, useNavigate } from "react-router-dom";
import { useEffect, useRef } from "react";

import "./style.css"
// import client from "../../../apis";
import { fetchLogin } from "./loginSlice";
import { userSelector } from "../../../store/selectors";
import Loading from "../../../components/Shares/Loading";
import { loginSelector } from "../../../store/selectors";

const Login = () => {
    const usernameRef = useRef();
    const passwordRef = useRef();
    const dispatch = useDispatch();

    const onSubmit = (e) => {
      e.preventDefault();
      const value = {
        username: usernameRef.current.value,
        password: passwordRef.current.value
      }
      dispatch(fetchLogin(value))
    };

    const user = useSelector(loginSelector)
  
    const navige = useNavigate();
    const isAuthenticated = useSelector(userSelector).role !== 0;
    useEffect(() => {
      if (isAuthenticated) {
        navige('/')
      }
    }, [isAuthenticated, navige])

    return (
        <div className="loginContainer">
          <form className="login__form" onSubmit={onSubmit}>
            <div className="login__header">
              <h1>Tài khoản</h1>
            </div>
            <div className="login__input">
              <input type="text" name="name" placeholder="Tên đăng nhập" ref={usernameRef} autoComplete="off"/>
              <span className="err"></span>
            </div>
            <div className="login__input">
              <input type="password" name="password" placeholder="Mật khẩu" ref={passwordRef} />
              <span className="err"></span>
            </div>
            <button className="login__submitBtn" type="submit">
            {user.status === "loading" ? <Loading/> : <></>}
              Đăng nhập
              </button>
            <Link to="/">Trở về trang chủ</Link>
          </form>
        </div>
      );
}

export default Login;