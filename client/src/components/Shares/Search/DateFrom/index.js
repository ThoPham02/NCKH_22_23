import Form from "react-bootstrap/Form";
import Col from "react-bootstrap/esm/Col";

const DateFrom = ({ dateFromRef }) => {
  
  return (
    <Form.Group as={Col} className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Control
        type="date"
        placeholder="Từ"
        ref={dateFromRef}
      />
    </Form.Group>
  );
};

export default DateFrom;
