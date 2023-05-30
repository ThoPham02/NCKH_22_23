import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";

import "./style.css";
import Loading from "../Loading";
import { useState } from "react";
import { useDispatch } from "react-redux";

const Confirm = (props) => {
  const { title, action, isLoading, content, onClick } = props
  const [show, setShow] = useState(false)

  const dispatch = useDispatch()
  const handle = () => {
    dispatch(action)
    setShow(false)
    onClick?.();
  }
  return (
    <>
      <Button onClick={() => setShow(true)} variant="danger">{title}</Button>

      <Modal
        size="sm"
        aria-labelledby="contained-modal-title-vcenter"
        show={show}
        backdrop="static"
        keyboard={false}
        centered
      >
        <Modal.Body>{content}</Modal.Body>
        <Modal.Footer>
          <Button onClick={() => setShow(false)} variant="danger">
            Close
          </Button>
          <Button onClick={handle} >
            {isLoading ? <Loading></Loading> : <></>}
            Xác nhận
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
};

export default Confirm;
