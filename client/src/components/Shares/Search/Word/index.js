import Form from "react-bootstrap/Form";

const Word = ({ value, setFilter }) => {
  const handleSearchChange = (e) => {
    setFilter(prevState => ({
      ...prevState,
      word: e.target.value
    }));
  };

  return (
    <Form.Group className="col-12 col-sm-12 col-md-12 col-lg-6">
      <Form.Control
        type="text"
        placeholder="Tìm kiếm theo tên"
        onChange={handleSearchChange}
        value={value}
      ></Form.Control>
    </Form.Group>
  );
};

export default Word;
