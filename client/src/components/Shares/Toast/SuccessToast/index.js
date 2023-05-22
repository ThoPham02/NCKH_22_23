import Toast from "react-bootstrap/Toast";

const Success = (props) => {
  return (
    <Toast onClose={props.handleClose} show={props.show} delay={3000} autohide style={{backgroundColor: "#45F94B"}}>
      <Toast.Header>
        <strong className="me-auto">{"Thành công"}</strong>
      </Toast.Header>
      <Toast.Body>{props.action === "cancel" ?  "Hủy đăng ký đè tài thành công" :props.message}</Toast.Body>
    </Toast>
  );
};

export default Success;