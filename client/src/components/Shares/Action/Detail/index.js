import {  useState } from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

import "./style.css";
import { userSelector } from "../../../../store/selectors";
import { TopicPlaceholder } from "../../Placeholder";
import TopicDetail from "./TopicDetail";
import {
  statusSelector,
  topicDetailSelector,
} from "../../../../store/selectors";
import { fetchTopicDetail } from "./TopicDetailSlice";
import {
  cancelTopic,
  registationTopic,
} from "../../../../pages/common/Topic/TopicSlice";
import Confirm from "../../Confirm";
import Toast from "../../Toast";
import { TopicAction } from "../../../../pages/common/Topic/TopicSlice";

const Detail = ({ name, topic }) => {
  const dispatch = useDispatch();
  const navige = useNavigate();

  const [show, setShow] = useState(false);
  const [confirmShow, setConfirmShow] = useState(false);
  const [action, setAction] = useState("");

  const user = useSelector(userSelector);
  const data = useSelector(topicDetailSelector);
  const status = useSelector(statusSelector);
  const topicDetail = data.topicDetail;
  let content = <p style={{fontWeight: "bold"}}>{topic.name}</p>

  const handleClose = () => setShow(false);
  const handleShow = () => {
    setShow(true);
    dispatch(TopicAction.setResult())
    dispatch(fetchTopicDetail(topic.id));
  };
  const handleRegisButton = () => {
    if (user.role === 0) {
      navige("/login");
    }

    content = <p style={{fontWeight: "bold"}}>Xác nhận đăng ký đề tài {topic.name}</p>
    setAction("regis")
    setConfirmShow(true);
  };
  const handleCancelButton = () => {
    if (user.role === 0) {
      navige("/login");
    }

    content = <p style={{fontWeight: "bold"}}>Xác nhận hủy đăng ký đề tài {topic.name}</p>
    setAction("cancel")
    setConfirmShow(true);
  };
  const handleConfirmButton = () => {
    if (action === "regis") {
      dispatch(registationTopic({ studentID: user.id, topicID: topic.id }));
    } else {
      dispatch(cancelTopic({ studentID: user.id }));
    }

    dispatch(fetchTopicDetail(topic.id));
    setConfirmShow(false);
  };

  return (
    <>
      <Button className="button" onClick={handleShow}>
        {name}
      </Button>

      <Modal show={show} onHide={handleClose} size="xl" centered>
        <Modal.Header closeButton>
          <Modal.Title>Chi tiết đề tài</Modal.Title>
        </Modal.Header>
        <Modal.Body style={{position: "relative"}}>
          {topicDetail.listStudent === undefined ? (
            <TopicPlaceholder />
          ) : (
            <TopicDetail
              topic={topic}
              topicDetail={topicDetail}
              user={user}
              status={status}
              handleRegisButton={handleRegisButton}
              handleCancelButton={handleCancelButton}
            />
          )}
          <Confirm
            confirmShow={confirmShow}
            setConfirmShow={setConfirmShow}
            handleConfirmButton={handleConfirmButton}
            content={content}
            variant="primary"
          />
          <Toast action={action}/>
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
