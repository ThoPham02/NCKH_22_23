import Form from "react-bootstrap/Form";

const Word = ({ wordRef }) => {
  return (
    <Form.Group className="col-12 col-sm-12 col-md-12 col-lg-6">
      <Form.Control
        type="text"
        placeholder="Tìm kiếm theo tên"
        ref={wordRef}
      ></Form.Control>
    </Form.Group>
  );
};

export default Word;
