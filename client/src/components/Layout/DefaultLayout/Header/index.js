import "./header.css";
import logo from "./lg.jpg";

function Header() {
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
              hệ thống quản lí <br></br>nghiên cứu khoa học sinh viên
            </p>
          </div>
        </div>
      </div>
      <div className="header__below"></div>
    </div>
  );
}

export default Header;
