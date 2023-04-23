import Form from "react-bootstrap/Form";

const Status = ({ value, setFilter }) => {
  const handleSearchStatus = (e) => {
    setFilter(prevState => ({
      ...prevState,
      status: e.target.value * 1
    }));
  };

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select onChange={handleSearchStatus} value={value}>
        <option value="0">Tất cả trạng thái</option>
        <option value="1">Hoàn thành</option>
        <option value="2">Quá hạn</option>
        <option value="3">Đang thực hiện</option>
      </Form.Select>
    </Form.Group>
  );
};

export default Status;
