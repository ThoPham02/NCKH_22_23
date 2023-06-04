import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";
import { fetchCommonEvents } from "./CommonEventSlice";
import { CommonEventSelector } from "../../../../store/selectors";


const Event = ({ event, setEvent }) => {
  const dispatch = useDispatch();
  const list = useSelector(CommonEventSelector).events;
  useEffect(() => {
    dispatch(fetchCommonEvents());
  }, [dispatch]);

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select value={event} onChange={(e) => setEvent(e.target.value)}>
        <option value={0}>Tất cả NCKH</option>
        {list ? list.map((item, index) => {
          return (
            <option value={item.id} key={index}>
              {item.name}
            </option>
          );
        }) : <></>}
      </Form.Select>
    </Form.Group>
  );
};
export default Event;