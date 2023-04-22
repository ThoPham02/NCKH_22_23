import Form from "react-bootstrap/Form";
import Col from "react-bootstrap/esm/Col";

const DateFrom = ({setDateFrom}) => {
  const handleDateFrom = (e) => {
    setDateFrom(e.target.value);
  };

  return (
    <Form.Group as={Col} className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Control
        type="date"
        placeholder="Ngày bắt đầu"
        onChange={handleDateFrom}
      />
    </Form.Group>
  );
};

export default DateFrom;
