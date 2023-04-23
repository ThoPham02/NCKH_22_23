import Col from "react-bootstrap/Col";
import Form from "react-bootstrap/Form";

const DateTo = ({ value, setFilter }) => {
  const handleDateTo = (e) => {
    setFilter(prevState => ({
      ...prevState,
      dateTo: e.target.value
    }));
  };

  return (
    <Form.Group as={Col} className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Control
        type="date"
        placeholder="Ngày kết thúc"
        onChange={handleDateTo}
        value={value}
      />
    </Form.Group>
  );
};

export default DateTo;
