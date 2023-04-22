import Form from "react-bootstrap/Form";

const Word = ({ setSearchWord }) => {
  const handleSearchChange = (e) => {
    setSearchWord(e.target.value);
  };

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Control
        type="text"
        placeholder="Tìm kiếm theo tên"
        onChange={handleSearchChange}
      ></Form.Control>
    </Form.Group>
  );
};

export default Word;
