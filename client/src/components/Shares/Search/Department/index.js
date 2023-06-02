import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";

import { departmentSelector } from "../../../../store/selectors";
import { fetchDepartment } from "./DepartmentSlice";

const Department = ({faculty, department, setDepartment, style, defaultValue }) => {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchDepartment(faculty));
    setDepartment(0)
  }, [dispatch, faculty, setDepartment]);

  const list = useSelector(departmentSelector).filter(item => faculty*1 === 0 || item.facultyID === faculty*1)
  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3" style={style}>
      <Form.Select value={department} onChange={e => setDepartment(e.target.value)}>
        <option value="0">{defaultValue}</option>
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

export default Department;
