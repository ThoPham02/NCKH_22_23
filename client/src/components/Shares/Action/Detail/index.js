import { useState } from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";

import "./style.css";
import useApi from "../../../../hooks/useGetApi";
import Loading from "../../Loading";
import TopicDetail from "../../TopicDetail";

const Detail = (props) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const { data, isLoading } = useApi(`/api/topic/${props.topicID}`);
  let topic = {
    id: 0,
    name: "",
    timeStart: "",
    timeEnd: "",
    lecture: "",
    faculty: "",
    listStudents: [],
    status: "",
    resultUrl: ""
  }
  if (data !== null && data !== undefined) {
    topic = data.topic
  }

  return (
    <>
      <Button className="button" onClick={handleShow}>
        {props.name}
      </Button>

      <Modal show={show} onHide={handleClose} size="xl">
        <Modal.Header closeButton>
          <Modal.Title>Chi tiết đề tài</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          {isLoading ? (
            <Loading></Loading>
          ) : (
            <TopicDetail
              id={topic.id}
              name={topic.name}
              timeStart={topic.timeStart}
              timeEnd={topic.timeEnd}
              lecture={topic.lecture}
              faculty={topic.faculty}
              students={topic.listStudents}
              status={topic.status}
              resultUrl={topic.resultUrl}
            />
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
