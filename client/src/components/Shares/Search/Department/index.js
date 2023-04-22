import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";

import { departmentSelector } from "../../../../store/selectors";
import { useEffect } from "react";
import { fetchDepartment } from "./DepartmentSlice";

const Department = ({ facultyId, setFilter }) => {
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(fetchDepartment());
  }, [dispatch]);

  const list = useSelector(departmentSelector);

  const handleSearchDepartment = (e) => {
    setFilter((prev) => (prev.department = e.target.value));
  };

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select onChange={handleSearchDepartment}>
        <option value="0">Tất cả các bộ môn</option>
        {list
          .filter((item) => item.facultyId === facultyId)
          .map((department, index) => {
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
