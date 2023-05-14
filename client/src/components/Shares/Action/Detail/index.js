import { useState } from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import Placeholder from "react-bootstrap/Placeholder";
import { useDispatch, useSelector } from "react-redux";

import "./style.css";
import { fetchTopicDetail } from "./TopicDetailSlice";
import { topicDetailSelector } from "../../../../store/selectors";
import { topicStatus } from "../../../../const/const";
import { convertTimestampToDateString } from "../../../../utils/time";

const Detail = ({ name, topicID }) => {
  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const dispatch = useDispatch();

  const handleShow = () => {
    setShow(true);
    dispatch(fetchTopicDetail(topicID));
  };

  const data = useSelector(topicDetailSelector);
  const topicDetail = data.topicDetail;
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
          {data.status === "loading" ? (
            <div className="topic-info">
              <div className="topic-info-item row">
                <div className="col-2">Tên đề tài</div>
                <div className="col-10">
                  <Placeholder xs={8} />
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Giảng viên hướng dẫn</div>
                <div className="col-10">
                  <Placeholder xs={3} />
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Liên hệ</div>
                <div className="col-10">
                  <Placeholder xs={4} />
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Khoa</div>
                <div className="col-10">
                  <Placeholder xs={3} />
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Bộ môn</div>
                <div className="col-10">
                  <Placeholder xs={3} />
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Trạng thái</div>
                <div className="col-10">
                  <Placeholder xs={2} />
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Nhóm sinh viên</div>
                <div className="col-10">
                  <Placeholder xs={8} />
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Thời gian thực hiện</div>
                <div className="col-10">
                  <Placeholder xs={6} />
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Tiểu ban</div>
                <div className="col-10">
                  <Placeholder xs={5} />
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Đại hội</div>
                <div className="col-10">
                  <Placeholder xs={6} />
                </div>
              </div>
            </div>
          ) : (
            <div className="topic-info">
              <div className="topic-info-item row">
                <div className="col-2">Tên đề tài</div>
                <div className="col-10">{topicDetail.name}</div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Đại hội</div>
                <div className="col-10">{topicDetail.event}</div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Tiểu ban</div>
                <div className="col-10">{topicDetail.subcommittee ? topicDetail.subcommittee : " ... "}</div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Giảng viên hướng dẫn</div>
                <div className="col-10">{topicDetail.lectureInfo.name}</div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Liên hệ</div>
                <div className="col-10">{topicDetail.lectureInfo.phone}</div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Khoa</div>
                <div className="col-10">{topicDetail.faculty}</div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Bộ môn</div>
                <div className="col-10">{topicDetail.department}</div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Trạng thái</div>
                <div className="col-10">
                  {topicStatus[topicDetail.status - 1]}
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Nhóm sinh viên:</div>
                <div className="col-10">
                  {topicDetail.listStudent? <></> : " ... "}
                </div>
              </div>
              <div className="topic-info-item row">
                <div className="col-2">Thời gian thực hiện</div>
                <div className="col-10">
                  Từ {convertTimestampToDateString(topicDetail.timeStart)} đến{" "}
                  {convertTimestampToDateString(topicDetail.timeEnd)}
                </div>
              </div>
            </div>
          )}
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
