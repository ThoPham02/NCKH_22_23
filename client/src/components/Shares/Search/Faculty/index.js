import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";

import { fetchFaculity } from "./FaculitySlice";
import { facultySelector } from "../../../../store/selectors";
import { useEffect } from "react";

const Faculty = ({ faculityRef }) => {
  const dispatch = useDispatch();
  const list = useSelector(facultySelector);
  useEffect(() => {
    dispatch(fetchFaculity());
  }, [dispatch]);

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select ref={faculityRef}>
        <option value={0}>Tất cả các khoa</option>
        {list.map((faculity, index) => {
          return (
            <option value={faculity.id} key={index}>
              {faculity.name}
            </option>
          );
        })}
      </Form.Select>
    </Form.Group>
  );
};

export default Faculty;
