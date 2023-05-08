import { HiOutlineDocument } from "react-icons/hi";
import { BiTimeFive } from "react-icons/bi";
import { AiFillStar } from "react-icons/ai";
import { FaUsers } from "react-icons/fa";
import { MdOutlineDriveFileMove } from "react-icons/md";

import "./style.css";

const TopicInfo = ({ name, dateTo, dateFrom, status, lecture, fileUrl }) => {
  return (
    <div className="topic-info">
      <div className="topic-info-item">
        <HiOutlineDocument /> <p>{name}</p>
      </div>
      <div className="topic-info-item">
        <BiTimeFive />
        <p>
          Từ {dateFrom.split(' ')[0]} đến {dateTo.split(' ')[0]}
        </p>
      </div>
      <div className="topic-info-item">
        <AiFillStar /> <p>{status}</p>
      </div>
      <div className="topic-info-item">
        <FaUsers />
        <p>{lecture}</p>
      </div>
      {fileUrl === "" ? (
        <></>
      ) : (
        <div className="topic-info-item">
          <MdOutlineDriveFileMove />
          <p>
            <a href={fileUrl}>Tệp đính kèm</a>
          </p>
        </div>
      )}
    </div>
  );
};

export default TopicInfo;
