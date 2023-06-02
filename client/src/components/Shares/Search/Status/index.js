import Form from "react-bootstrap/Form";
import { getListStatus } from "../../../../utils/getListStatus";
import { topicStatus } from "../../../../const/const";

const Status = ({ status, setStatus, sumStatus }) => {
  const listStatus = getListStatus(sumStatus)
  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select value={status} onChange={e => setStatus(e.target.value)}>
        <option value="0">Tất cả trạng thái</option>
        {topicStatus.filter((item) => listStatus.find(id => id === item.id)).map(item => {
          return <option value={item.id} key={item.id}>{item.name}</option>
        })}
      </Form.Select>
    </Form.Group>
  );
};

export default Status;