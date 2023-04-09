import Form from "react-bootstrap/Form";

const SearchWord = ({ setSearchWord }) => {
  const handleSearchChange = (e) => {
    setSearchWord(e.target.value);
  };

  return (
    <Form.Control
      type="text"
      placeholder="Tìm kiếm theo tên"
      onChange={handleSearchChange}
    ></Form.Control>
  );
};

export default SearchWord;
