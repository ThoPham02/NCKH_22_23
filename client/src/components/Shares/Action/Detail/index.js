import { useState } from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import { HiOutlineDocument } from "react-icons/hi";
import { BiTimeFive } from "react-icons/bi";
import { AiFillStar, AiFillHome } from "react-icons/ai";
import { FaUsers, FaUserGraduate } from "react-icons/fa";
import { MdOutlineDriveFileMove } from "react-icons/md";
import { DiCode } from "react-icons/di";

import "./style.css";
import { topicStatus } from "../../../../const/const";
import { convertTimestampToDateString } from "../../../../utils/time";

const Detail = ({ name, topic }) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  return (
    <>
      <Button className="button" onClick={handleShow}>
        {name}
      </Button>

      <Modal show={show} onHide={handleClose} size="xl">
        <Modal.Header closeButton>
          <Modal.Title>Chi tiết đề tài</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <div className="topic-info">
            <div className="topic-info-item">
              <DiCode /> <p>{topic.id}</p>
            </div>
            <div className="topic-info-item">
              <HiOutlineDocument /> <p>{topic.name}</p>
            </div>
            <div className="topic-info-item">
              <AiFillHome /> <p>{topic.departmentID}</p>
            </div>
            <div className="topic-info-item">
              <BiTimeFive />
              <p>
                Từ {convertTimestampToDateString(topic.timeStart)} đến{" "}
                {convertTimestampToDateString(topic.timeEnd)}
              </p>
            </div>
            <div className="topic-info-item">
              <AiFillStar /> <p>{topicStatus[topic.status - 1]}</p>
            </div>
            <div className="topic-info-item">
              <FaUserGraduate />
              <p>{topic.lectureName}</p>
            </div>
            <div className="topic-info-item">
              <FaUsers />
              {/* <p>{topic === 0 ? "" :topic.map((item) => item + ", ")}</p> */}
            </div>

            {!topic.resultUrl ? (
              <></>
            ) : (
              <div className="topic-info-item">
                <MdOutlineDriveFileMove />
                <p>
                  <a href={topic.resultUrl}>Tệp đính kèm</a>
                </p>
              </div>
            )}
          </div>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
};

export default Detail;
