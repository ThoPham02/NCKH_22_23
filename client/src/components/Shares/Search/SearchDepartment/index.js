import Form from "react-bootstrap/Form";

const SearchDepartment = ({ setSearchDepartment }) => {
  const handleSearchDepartment = (e) => {
    setSearchDepartment(e.target.value);
  };

  return (
    <Form.Select onChange={handleSearchDepartment}>
      <option value="0">Tất cả các bộ môn</option>
      <option value="1">hoàn thành</option>
      <option value="2">quá hạn</option>
      <option value="3">đang thực hiện</option>
    </Form.Select>
  );
};

export default SearchDepartment;
