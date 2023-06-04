import { BiUserCircle } from "react-icons/bi";
import { Link, useNavigate } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";

import "./navbar.css";
import { LoginActions } from "../../../../pages/common/Login/LoginSlice";
import { userSelector } from "../../../../store/selectors";

const NavBar = () => {
  const navige = useNavigate()
  const dispatch = useDispatch();
  const user = useSelector(userSelector);
  function handleLogout() {
    dispatch(LoginActions.logout());

    navige('/')
  }

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
    <div className="navbar">
      <ul className="navbar-content container">
        <li>
          <Link to="/home">Trang chủ</Link>
        </li>
        <li>
          <Link to="#">
            Đề tài
          </Link>
          <ul className="subnav">
            <li>
              <Link to={"/topics"}>Tất cả đề tài</Link>
            </li>
            <li>
              <Link to={user.role === 0 ? "/login" : "/my-topic"}>Đề tài của tôi</Link>
            </li>
          </ul>
        </li>
        {user.role !== 0 ?
          <li>
            <Link to="#">Nghiệm Thu</Link>
            <ul className="subnav">
              <li>
                <Link to={"/subcommittee-report"}>Cấp Tiểu Ban</Link>
              </li>
              <li>
                <Link to={"/school-report"}>Cấp Trường</Link>
              </li>
              {
                user.role === 2 ?
                  <li>
                    <Link to={"/mark"}>Chấm điểm</Link>
                  </li> : <></>
              }
            </ul>
          </li>
          : <></>}
        <li>
          <Link to="/result">Kết quả</Link>
        </li>
        <li>
          <Link to="/contact">Liên hệ</Link>
        </li>
        <li className="nav__login">{loginAccount}</li>
      </ul>
    </div>
  );
};

export default NavBar;
