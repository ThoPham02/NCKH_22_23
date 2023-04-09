import Form from "react-bootstrap/Form";

const SearchFaculity = ({ setSearchFaculity }) => {
  const handleSearchFaculity = (e) => {
    setSearchFaculity(e.target.value);
  };
  return (
    <Form.Select onChange={handleSearchFaculity}>
      <option value="0">Tất cả các khoa</option>
      <option value="1">hoàn thành</option>
      <option value="2">quá hạn</option>
      <option value="3">đang thực hiện</option>
    </Form.Select>
  );
};

export default SearchFaculity;
