import Form from "react-bootstrap/Form";

const Status = ({ setSearchStatus }) => {
  const handleSearchStatus = (e) => {
    setSearchStatus(e.target.value);
  };

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select onChange={handleSearchStatus}>
        <option value="0">tất cả trạng thái</option>
        <option value="1">hoàn thành</option>
        <option value="2">quá hạn</option>
        <option value="3">đang thực hiện</option>
      </Form.Select>
    </Form.Group>
  );
};

export default Status;
