import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";

import "./style.css";
import Loading from "../Loading";

const Confirm = (props) => {
  const {confirmShow, setConfirmShow, handleConfirmButton, isLoading, content} = props
  return (
    <Modal
      size="sm"
      aria-labelledby="contained-modal-title-vcenter"
      show={confirmShow}
      backdrop="static"
      keyboard={false}
      centered
    >
      <Modal.Body>{content}</Modal.Body>
      <Modal.Footer>
        <Button onClick={() => setConfirmShow(false)} variant="danger">
          Close
        </Button>
        <Button onClick={handleConfirmButton}>
          {isLoading ? <Loading></Loading> : <></>}
          Xác nhận
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default Confirm;
