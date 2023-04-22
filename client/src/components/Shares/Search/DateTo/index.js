import Col from "react-bootstrap/Col";
import Form from "react-bootstrap/Form";

const DateTo = ({setDateTo}) => {
  const handleDateTo = (e) => {
    setDateTo(e.target.value);
  }

  return (
      <Form.Group as={Col} className="col-12 col-sm-12 col-md-6 col-lg-3">
        <Form.Control type="date" placeholder="Ngày kết thúc" onChange={handleDateTo}/>
      </Form.Group>
  );
};

export default DateTo;
