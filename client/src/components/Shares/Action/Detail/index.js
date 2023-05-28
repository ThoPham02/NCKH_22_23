import {  useEffect, useState } from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

import "./style.css";
import { ShareTopicDetailSelector, userSelector } from "../../../../store/selectors";
import { TopicPlaceholder } from "../../Placeholder";
import TopicDetail from "./TopicDetail";
import { fetchTopicDetail } from "./ShareTopicDetailSlice";


const Detail = ({ name, topicIn }) => {
  const dispatch = useDispatch();
  const [show, setShow] = useState(false);
  const topicDetail = useSelector(ShareTopicDetailSelector)
  const [detail, setDetail] = useState({})

  useEffect(() => {
    if (!topicDetail.topic) {
      dispatch(fetchTopicDetail({id: topicIn.id}))
    }
     // eslint-disable-next-line
  }, [dispatch, topicDetail.topic])
  useEffect(() => {
    if (topicDetail.topic) {
      setDetail(topicDetail)
    }
     // eslint-disable-next-line
  }, [topicDetail.topic])

  let isLoading = detail.status === 'loading'

  const user = useSelector(userSelector)
  const navigate = useNavigate()
  const handleRegisButton = () => {
    if (user.role === 0) {
      navigate("/login")
    }

    
  }
  const handleCancelButton = () => {
    
  }
  return (
    <>
      <Button className="button" onClick={() => setShow(true)}>
        {name}
      </Button>

      <Modal show={show} onHide={() => setShow(false)} size="xl" centered>
        <Modal.Header closeButton>
          <Modal.Title>Chi tiết đề tài</Modal.Title>
        </Modal.Header>
        <Modal.Body style={{position: "relative"}}>
          {isLoading ? (
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
            />
          )}
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShow(false)}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
};

export default Detail;
