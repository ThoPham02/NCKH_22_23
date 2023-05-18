import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";

import "./style.css";

const Confirm = (props) => {
  return (
    <Modal
      size="sm"
      aria-labelledby="contained-modal-title-vcenter"
      show={props.confirmShow}
      backdrop="static"
      keyboard={false}
      centered
    >
      <Modal.Body>{props.content}</Modal.Body>
      <Modal.Footer>
        <Button onClick={() => props.setConfirmShow(false)} variant="danger">
          Close
        </Button>
        <Button onClick={props.handleConfirmButton}>Xác nhận</Button>
      </Modal.Footer>
    </Modal>
  );
};

export default Confirm;
