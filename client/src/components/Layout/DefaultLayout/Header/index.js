import "./header.css";
import logo from "./lg.jpg";
import { Link } from "react-router-dom";

import { BiUserCircle } from "react-icons/bi";
function Header() {
  return (
    <div className="header">
      <div className="header__above">
        <div className="header-logo">
          <div className="logo">
            <img src={logo} alt="###" />
            <ul className="logo-name">
              <li className="name-VN">Trường đại học mỏ - địa chất</li>
              <li className="name-EN">hanoi university of mining and geology</li>
            </ul>
          </div>
          <div className="header-logo__right">
            <p>hệ thống quản lí <br></br>nghiên cứu khoa học</p>
          </div>
        </div>
        <div className="nav">
          <ul>
            <li><Link to="/">Trang chủ</Link></li>
            <li><Link to="/topic">đề tài</Link></li>
            <li><Link to="/statistical">thống kê</Link></li>
            <li><Link to="/category">danh mục</Link></li>
            <li><Link to="/contact">liên hệ</Link></li>
          </ul>
          <div className="account">
            <Link to="login">
              <BiUserCircle />
              <span>Đăng nhập</span>
            </Link>
          </div>
        </div>
      </div>
      <div className="header__below"></div>
    </div>
  );
}

export default Header;
