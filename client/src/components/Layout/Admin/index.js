import { MdOutlineSpaceDashboard, MdTopic } from "react-icons/md";
import { BiUserCircle, BiCalendarEvent } from "react-icons/bi";
import { TbReport } from "react-icons/tb";
import {AiOutlineFileDone} from "react-icons/ai"
import { Link, useNavigate } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";

import "./style.css";
import logo from "./lg.jpg";
import { userSelector } from "../../../store/selectors";
import { LoginActions } from "../../../pages/common/Login/LoginSlice";

const Admin = ({ children }) => {
  const navige = useNavigate();
  const dispatch = useDispatch();
  const user = useSelector(userSelector);
  function handleLogout() {
    dispatch(LoginActions.logout());

    navige("/");
  }

  if (user.role !== 5 && user.role !== 4) {
    navige("/");
  }
  return (
    <div className="admin">
      <div className="admin_user">
        <div className="admin_user_logo">
          <img src={logo} alt="" style={{marginRight: "24px"}}/>
          Quản lý NCKH HUMG
        </div>
        <div className="admin_user_info">
          <BiUserCircle /> {user.name}
          <ul className="admin_user_action">
            <li>Thông tin cá nhân</li>
            <li onClick={handleLogout}>Đăng xuất</li>
          </ul>
        </div>
      </div>
      <div className="admin_body">
        <div className="admin_nav">
          <ul>
            <li className="admin_nav_item">
              <Link to={"/admin-home"}>
                <span className="nav_icon">
                  <MdOutlineSpaceDashboard />
                </span>
                <span className="nav_text">DashBoard</span>
              </Link>
            </li>
            <li className="admin_nav_item">
              <Link to={"/admin-event"}>
                <span className="nav_icon">
                  <BiCalendarEvent />
                </span>
                <span className="nav_text">Đại hội</span>
              </Link>
            </li>
            <li className="admin_nav_item">
              <Link to={"/admin-topic"}>
                <span className="nav_icon">
                  <MdTopic />
                </span>
                <span className="nav_text">Đề tài</span>
              </Link>
            </li>
            <li className="admin_nav_item">
              <Link to={"/admin-report"}>
                <span className="nav_icon">
                  <TbReport />
                </span>
                <span className="nav_text">Nghiệm thu</span>
              </Link>
            </li>
            <li className="admin_nav_item">
              <Link to={"/admin-result"}>
                <span className="nav_icon">
                  <AiOutlineFileDone />
                </span>
                <span className="nav_text">Kết quả</span>
              </Link>
            </li>
          </ul>
        </div>
        <div className="admin_content">{children}</div>
      </div>
    </div>
  );
};

export default Admin;
