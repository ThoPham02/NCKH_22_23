import Form from "react-bootstrap/Form";
import { topicStatus } from "../../../../const/const";

const Status = ({ statusRef }) => {
  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select ref={statusRef}>
        <option value="0">Tất cả trạng thái</option>
        {topicStatus.map((item, index) => {
          return <option value={index+1} key={index}>{item}</option>
        })}
      </Form.Select>
    </Form.Group>
  );
};

export default Status;
