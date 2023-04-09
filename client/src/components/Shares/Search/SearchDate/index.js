import Col from "react-bootstrap/Col";
import Form from "react-bootstrap/Form";
import Row from "react-bootstrap/Row";

const SearchDate = ({setDateFrom, setDateTo}) => {
  const handleDateFrom = (e) => {
    setDateFrom(e.target.value);
  }

  const handleDateTo = (e) => {
    setDateTo(e.target.value);
  }

  return (
    <Row>
      <Form.Group as={Col} className="search-col">
        <Form.Control type="date" placeholder="Ngày bắt đầu" onChange={handleDateFrom}/>
      </Form.Group>
      <Form.Group as={Col} className="search-col">
        <Form.Control type="date" placeholder="Ngày kết thúc" onChange={handleDateTo}/>
      </Form.Group>
    </Row>
  );
};

export default SearchDate;
