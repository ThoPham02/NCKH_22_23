import {
    FaUserAlt,
    FaUserGraduate,
    FaUsers,
    FaHouseUser,
  } from "react-icons/fa";
  import { MdTopic } from "react-icons/md";
  
  import Card from "../../../components/Shares/Card";
  import NumberCounter from "../../../components/Shares/NumberCount";
  import "./style.css";
  
  const DashBoard = () => {
    return (
      <div className="admin_dash_board">
        <Card title={"Thống kê chung"}>
          <h3>Đại hội nghiên cứu khoa học lần thứ 36 HUMG</h3>
          <div className="dash_board_items">
            <div className="dash_board_item">
              <FaHouseUser />
              <div className="dash_board_item_content">
                <NumberCounter target={12} duration={1} className={"counter"} /> Khoa
              </div>
            </div>
            <div className="dash_board_item">
              <FaUsers />
              <div className="dash_board_item_content">
                <NumberCounter target={60} duration={1} className={"counter"} /> Bộ môn
              </div>
            </div>
  
            <div className="dash_board_item">
              <MdTopic />
              <div className="dash_board_item_content">
                <NumberCounter target={100} duration={1} className={"counter"} /> Đề tài
              </div>
            </div>
  
            <div className="dash_board_item">
              <FaUserGraduate />
              <div className="dash_board_item_content">
                <NumberCounter target={50} duration={1} className={"counter"} /> Giảng viên hướng dẫn
              </div>
            </div>
  
            <div className="dash_board_item">
              <FaUserAlt />
              <div className="dash_board_item_content">
                <NumberCounter target={200} duration={1} className={"counter"} /> Sinh viên
              </div>
            </div>
          </div>
        </Card>
      </div>
    );
  };
  
  export default DashBoard;
  