import { HiOutlineDocument } from "react-icons/hi";
import { BiTimeFive } from "react-icons/bi";
import { AiFillStar, AiFillHome } from "react-icons/ai";
import { FaUsers, FaUserGraduate } from "react-icons/fa";
import { MdOutlineDriveFileMove } from "react-icons/md";
import { DiCode } from "react-icons/di";

import "./style.css";

const TopicDetail = (props) => {
  return (
    <div className="topic-info">
      <div className="topic-info-item">
        <DiCode /> <p>{props.id}</p>
      </div>
      <div className="topic-info-item">
        <HiOutlineDocument /> <p>{props.name}</p>
      </div>
      <div className="topic-info-item">
        <AiFillHome /> <p>{props.faculty}</p>
      </div>
      <div className="topic-info-item">
        <BiTimeFive />
        <p>
          Từ {props.timeStart.split(" ")[0]} đến {props.timeEnd.split(" ")[0]}
        </p>
      </div>
      <div className="topic-info-item">
        <AiFillStar /> <p>{props.status}</p>
      </div>
      <div className="topic-info-item">
        <FaUserGraduate />
        <p>{props.lecture}</p>
      </div>
      <div className="topic-info-item">
        <FaUsers />
        <p>{props.students.length === 0 ? "" :props.students.map((item) => item + ", ")}</p>
      </div>

      {props.resultUrl === "" ? (
        <></>
      ) : (
        <div className="topic-info-item">
          <MdOutlineDriveFileMove />
          <p>
            <a href={props.resultUrl}>Tệp đính kèm</a>
          </p>
        </div>
      )}
    </div>
  );
};

export default TopicDetail;
