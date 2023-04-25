import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";

import { departmentSelector } from "../../../../store/selectors";
import { fetchDepartment } from "./DepartmentSlice";

const Department = ({ faculty, value, setFilter }) => {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchDepartment());
  }, [dispatch]);

  const list = useSelector(departmentSelector);

  const departments = list.filter(
    (item) => faculty === 0 || item.facultyId === faculty
  );
  const handleSearchDepartment = (e) => {
    setFilter((prevState) => ({
      ...prevState,
      departmentId: e.target.value * 1,
      facultyId:
        e.target.value * 1 === 0
          ? 0
          : list.filter((item) => item.id === e.target.value * 1)[0].facultyId,
    }));
  };

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select onChange={handleSearchDepartment} value={value}>
        <option value="0">Tất cả các bộ môn</option>
        {departments.map((department, index) => {
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
