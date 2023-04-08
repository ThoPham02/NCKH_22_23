import "./header.css";
import logo from "./lg.jpg";
import { loginSelector } from "../../../../store/selectors";

import { Link } from "react-router-dom";
import { useSelector } from "react-redux";
import { BiUserCircle } from "react-icons/bi";
function Header() {
  const login = useSelector(loginSelector);

  var loginAccount = (
    <Link to="login">
      <BiUserCircle />
      <span>Đăng nhập</span>
    </Link>
  );

  if (login.user.name !== undefined) {
    loginAccount = <span>Xin Chào, {login.user.name}</span>;
  }

  return (
    <div className="header">
      <div className="header__above">
        <div className="header-logo">
          <div className="logo">
            <img src={logo} alt="###" />
            <ul className="logo-name">
              <li className="name-VN">Trường đại học mỏ - địa chất</li>
              <li className="name-EN">
                hanoi university of mining and geology
              </li>
            </ul>
          </div>
          <div className="header-logo__right">
            <p>
              hệ thống quản lí <br></br>nghiên cứu khoa học
            </p>
          </div>
          <div className="account">{loginAccount}</div>
        </div>
        <div className="nav">
          <ul>
            <li>
              <Link to="/">Trang chủ</Link>
            </li>
            <li>
              <Link to="/topic">đề tài</Link>
            </li>
            <li>
              <Link to="/statistical">thống kê</Link>
            </li>
            <li>
              <Link to="/category">danh mục</Link>
            </li>
            <li>
              <Link to="/contact">liên hệ</Link>
            </li>
          </ul>
        </div>
      </div>
      <div className="header__below"></div>
    </div>
  );
}

export default Header;
