import { useState } from "react";
import Placeholder from "react-bootstrap/Placeholder";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";

import "./style.css";
import Member from "./TopicMember";
import FindMemberItem from "./FindMember";

const TopicRegistration = ({ name, topicRegis }) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);
  const handleSubmit = (e) => {
    e.preventDefault();
    
  };

  return (
    <>
      <Button className="button" onClick={handleShow}>
        {name}
      </Button>

      <Modal show={show} onHide={handleClose} size="xl">
        <Modal.Header closeButton>
          <Modal.Title>Đăng ký đề tài</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" type="submit">
            Xác Nhận
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
};

export default TopicRegistration;
