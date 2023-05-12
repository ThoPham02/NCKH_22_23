import { BiUserCircle } from "react-icons/bi";
import { Link } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";

import "./navbar.css";
import { LoginActions } from "../../../../pages/common/Login/loginSlice";
import { userSelector } from "../../../../store/selectors";
import { userNav } from "./data";

const NavBar = () => {
  const dispatch = useDispatch();
  const user = useSelector(userSelector);
  function handleLogout() {
    dispatch(LoginActions.logout());
  }

  console.log(user.role)

  var loginAccount = (
    <div>
      <Link to="/login">
        <BiUserCircle />
        <span>Đăng nhập</span>
      </Link>
    </div>
  );

  if (user.role !== 0) {
    loginAccount = (
      <div>
        <span className="user-name">Xin Chào, {user.name}</span>
        <ul className="account__nav">
          <li>
            <Link to="/user-info">Thông tin</Link>
          </li>
          <li onClick={handleLogout} to="/login">
            Đăng xuất
          </li>
        </ul>
      </div>
    );
  }
  return (
    <ul className="nav">
      <li>
        <Link to="/home">Trang chủ</Link>
      </li>
      {userNav[user.role].map((item, index) => (
        <li key={index}>
          <Link to={item.path}>{item.name}</Link>
        </li>
      ))}
      <li>
        <Link to="/contact">Liên hệ</Link>
      </li>
      <li className="nav__login">
        {loginAccount}
      </li>
    </ul>
  );
};

export default NavBar;
