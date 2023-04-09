import Form from "react-bootstrap/Form";

const SearchStatus = ({setSearchStatus}) => {
  const handleSearchStatus = (e) => {
    setSearchStatus(e.target.value)
  }

  return (
    <Form.Select onChange={handleSearchStatus}>
      <option value="0">tất cả trạng thái</option>
      <option value="1">hoàn thành</option>
      <option value="2">quá hạn</option>
      <option value="3">đang thực hiện</option>
    </Form.Select>
  );
};

export default SearchStatus;
