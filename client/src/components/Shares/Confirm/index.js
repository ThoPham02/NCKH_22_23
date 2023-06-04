import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import { useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import "./style.css";
import Loading from "../Loading";
import { userSelector } from "../../../store/selectors";
import { useNavigate } from "react-router-dom";

const Confirm = (props) => {
  const { title, action, isLoading, content, onClick, variant, isAction } = props
  const [show, setShow] = useState(false)

  const dispatch = useDispatch()
  const handle = () => {
    dispatch(action)
    if (!isLoading) {
      setShow(false)
    }
    onClick?.();
  }

  const user = useSelector(userSelector)
  const navige = useNavigate();
  const handleCLick = () => {
    if (user.role === 0) {
      navige("/login")
    }
    setShow(true)
  }
  return (
    <>
      <Button onClick={handleCLick} variant={variant} className={isAction ? "button" : ""} >{title}</Button>

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
            Hủy
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
