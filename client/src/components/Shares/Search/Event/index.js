import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";
import { fetchEvents } from "./CommonEventSlice";
import { CommonEventSelector } from "../../../../store/selectors";


const Event = ({eventRef}) => {
  const dispatch = useDispatch();
  const list = useSelector(CommonEventSelector).events;
  useEffect(() => {
    dispatch(fetchEvents());
  }, [dispatch]);

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select ref={eventRef}>
        <option value={0}>Tất cả NCKH</option>
        {list.map((item, index) => {
          return (
            <option value={item.id} key={index}>
              {item.name}
            </option>
          );
        })}
      </Form.Select>
    </Form.Group>
  );
};
export default Event;