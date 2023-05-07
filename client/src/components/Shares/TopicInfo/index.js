import { HiOutlineDocument } from "react-icons/hi";
import { BiTimeFive } from "react-icons/bi";
import { AiFillStar } from "react-icons/ai";
import { FaUsers } from "react-icons/fa";
import { MdOutlineDriveFileMove } from "react-icons/md";

import "./style.css";

const TopicInfo = ({ name, dateTo, dateFrom, status, students, fileUrl }) => {
  return (
    <div className="topic-info">
      <div className="topic-info-item">
        <HiOutlineDocument /> <p>{name}</p>
      </div>
      <div className="topic-info-item">
        <BiTimeFive />{" "}
        <p>
          Thời gian từ {dateFrom} đến {dateTo}
        </p>
      </div>
      <div className="topic-info-item">
        <AiFillStar /> <p>{status}</p>
      </div>
      <div className="topic-info-item">
        <FaUsers />{" "}
        <p>{students.length === 0 ? "" : students.map((student) => student)}</p>
      </div>
      {fileUrl === "" ? (
        <></>
      ) : (
        <div className="topic-info-item">
          <MdOutlineDriveFileMove />{" "}
          <p>
            <a href={fileUrl}>Tệp đính kèm</a>
          </p>
        </div>
      )}
    </div>
  );
};

export default TopicInfo;
