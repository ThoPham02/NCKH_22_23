import { useEffect, useState } from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

import "./style.css";
import { ShareTopicDetailSelector, userSelector } from "../../../../store/selectors";
import { TopicPlaceholder } from "../../Placeholder";
import TopicDetail from "./TopicDetail";
import { ShareTopicDetailAction, addStudentGroup, deleteStudentGroup, fetchTopicDetail } from "./ShareTopicDetailSlice";
import Confirm from "../../Confirm";


const Detail = ({ name, topicIn }) => {
  const dispatch = useDispatch();
  const [show, setShow] = useState(false);
  const topicDetail = useSelector(ShareTopicDetailSelector)
  const [detail, setDetail] = useState({})
  useEffect(() => {
    if (topicDetail.topic) {
      setDetail(topicDetail)
    }
    // eslint-disable-next-line
  }, [topicDetail.topic])

  let isLoading = detail.status === 'loading'

  const user = useSelector(userSelector)
  const [confirmShow, setConfirmShow] = useState(false)
  const navigate = useNavigate()
  const handleRegisButton = () => {
    if (user.role === 0) {
      navigate("/login")
    }
    setConfirmShow(true)
  }
  const handleCancelButton = () => {
    setConfirmShow(true)
  }
  const isCancel = user.role === 1 && detail.listStudent && detail.listStudent.find((item) => item.studentID === user.id);
  const handleConfirmButton = () => {
    if (isCancel) {
      dispatch(deleteStudentGroup({ topicID: topicIn.id, studentID: user.id }))
    } else {
      dispatch(addStudentGroup({ topicID: topicIn.id, studentID: user.id }))
    }
    setConfirmShow(false)
  }
  const handleShow = () => {
    dispatch(fetchTopicDetail({ id: topicIn.id }))
    setShow(true)
  }
  const handleClose = () => {
    setDetail({})
    dispatch(ShareTopicDetailAction.clearStore())
    setShow(false)
  }

  console.log(topicDetail)

  return (
    <>
      <Button className="button" onClick={handleShow}>
        {name}
      </Button>

      <Modal show={show} onHide={handleClose} size="xl" centered>
        <Modal.Header closeButton>
          <Modal.Title>Chi tiết đề tài</Modal.Title>
        </Modal.Header>
        <Modal.Body style={{ position: "relative" }}>
          {isLoading || !detail.topic ? (
            <TopicPlaceholder />
          ) : (
            <TopicDetail
              topic={detail.topic}
              event={detail.event}
              subcommittee={detail.subcommittee}
              reports={detail.reports}
              marks={detail.marks}
              listStudent={detail.listStudent}
              isLoading={isLoading}
              handleRegisButton={handleRegisButton}
              handleCancelButton={handleCancelButton}
              isCancel={isCancel}
            />
          )}
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>

      <Confirm confirmShow={confirmShow} setConfirmShow={setConfirmShow} isLoading={isLoading} content={"Xác nhận thao tác!"} handleConfirmButton={handleConfirmButton} />
    </>
  );
};

export default Detail;
