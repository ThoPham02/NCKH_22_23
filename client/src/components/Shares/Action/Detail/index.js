import { useState } from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import { useDispatch, useSelector } from "react-redux";

import { TopicDetailSelector } from "../../../../store/selectors";
import { TopicPlaceholder } from "../../Placeholder";
import TopicDetail from "./TopicDetail";
import { fetchTopicDetail } from "./TopicDetailSlice";


const Detail = ({ name, topicIn }) => {
  const dispatch = useDispatch();
  const [show, setShow] = useState(false);
  const topicDetail = useSelector(TopicDetailSelector)

  const handleShow = () => {
    dispatch(fetchTopicDetail({ id: topicIn.id }))
    setShow(true);
  }

  const isLoading = topicDetail.status === 'loading'

  return (
    <>
      <Button className="button" onClick={handleShow} style={{ width: "120px" }}>
        {name}
      </Button>

      <Modal show={show} onHide={() => setShow(false)} size="xl" centered>
        <Modal.Header closeButton>
          <Modal.Title>Chi tiết đề tài</Modal.Title>
        </Modal.Header>
        <Modal.Body style={{ position: "relative" }}>
          {isLoading || !topicDetail.topic ?
            <TopicPlaceholder />
            :
            <TopicDetail data={topicDetail} />
          }
        </Modal.Body>
      </Modal>
    </>
  );
};

export default Detail;
