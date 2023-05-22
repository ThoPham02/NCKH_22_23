import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";

import { fetchFaculty } from "./FacultySlice";
import { facultySelector } from "../../../../store/selectors";
import { useEffect } from "react";

const Faculty = ({faculty, setFaculty, facultyRef }) => {
  const dispatch = useDispatch();
  const list = useSelector(facultySelector);
  useEffect(() => {
    dispatch(fetchFaculty());
  }, [dispatch]);

  const handleChangle = (e) => {
    setFaculty(e.target.value)
  }

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select ref={facultyRef} value={faculty} onChange={handleChangle}>
        <option value={0}>Tất cả các khoa</option>
        {list.map((faculty, index) => {
          return (
            <option value={faculty.id} key={index}>
              {faculty.name}
            </option>
          );
        })}
      </Form.Select>
    </Form.Group>
  );
};

export default Faculty;
