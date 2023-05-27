import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";

import "./style.css";
import Loading from "../Loading";

const Confirm = (props) => {
  console.log(props.loading)
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
        <Button onClick={props.handleConfirmButton}>
          {props.loading ? <Loading></Loading> : <></>}
          Xác nhận
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default Confirm;
