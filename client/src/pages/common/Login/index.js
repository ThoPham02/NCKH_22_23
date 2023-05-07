import { useDispatch, useSelector } from "react-redux";
import { Link, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

import "./style.css"
import client from "../../../apis";
import { LoginActions } from "./loginSlice";
import { userSelector } from "../../../store/selectors";
import useForm from "../../../hooks/useForm";
import Loading from "../../../components/Shares/Loading";

const Login = () => {
    const dispatch = useDispatch();
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState(null);
    const onSubmit = async (values) => {
      setIsLoading(true)
      try {
        const res = await client.post("/api/user/login", values)
  
        dispatch(LoginActions.login(res.data))
        setIsLoading(false)
      } catch (err) {
        setError(err)
        console.log(error)
      } finally {
        setIsLoading(false)
      }
  
    };
    const {values, handleChange, handleSubmit} = useForm(onSubmit, { name: "", password: ""})
  
    const navige = useNavigate();
    const isAuthenticated = useSelector(userSelector).typeAccount !== 0;
    useEffect(() => {
      if (isAuthenticated) {
        navige('/')
      }
    }, [isAuthenticated, navige])

    return (
        <div className="loginContainer">
          <form className="login__form" onSubmit={handleSubmit}>
            <div className="login__header">
              <h1>Tài khoản</h1>
            </div>
            <div className="login__input">
              <input type="text" name="name" placeholder="Tên đăng nhập" value={values.name} onChange={handleChange} autoComplete="off"/>
              <span className="err"></span>
            </div>
            <div className="login__input">
              <input type="password" name="password" placeholder="Mật khẩu" value={values.password} onChange={handleChange} />
              <span className="err"></span>
            </div>
            <button className="login__submitBtn" type="submit">
            {isLoading ? <Loading/> : <></>}
              Đăng nhập
              </button>
            <Link to="/">Trở về trang chủ</Link>
          </form>
        </div>
      );
}

export default Login;