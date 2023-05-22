import Toast from "react-bootstrap/Toast";

const Error = (props) => {
  return (
    <Toast onClose={props.handleClose} show={props.show} delay={3000} bg="danger" autohide>
      <Toast.Header>
        <strong className="me-auto">{props.title}</strong>
      </Toast.Header>
      <Toast.Body>{props.message}</Toast.Body>
    </Toast>
  );
};

export default Error;
