import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";

import { departmentSelector } from "../../../../store/selectors";
import { fetchDepartment } from "./DepartmentSlice";

const Department = ({faculty, departmentRef, style }) => {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchDepartment(faculty));
  }, [dispatch, faculty]);

  const list = useSelector(departmentSelector).filter(item => faculty === 0 || item.facultyID === faculty*1)

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3" style={style}>
      <Form.Select ref={departmentRef}>
        <option value="0">Tất cả các bộ môn</option>
        {list.map((department, index) => {
          return (
            <option value={department.id} key={index}>
              {department.name}
            </option>
          );
        })}
      </Form.Select>
    </Form.Group>
  );
};

export default Department;
