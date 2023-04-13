import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";
import { departmentSelector } from "../../../../store/selectors";
import { useEffect, useState } from "react";
import { fetchDepartment } from "./DepartmentSlice";

const SearchDepartment = ({ setSearchDepartment }) => {
  const dispatch = useDispatch();
  const list = useSelector(departmentSelector);
  const [listDepartments, setListDepartments] = useState(list);
  useEffect(() => {
    dispatch(fetchDepartment());
  }, [dispatch]);

  useEffect(() => {
    setListDepartments(list);
  }, [list]);

  const handleSearchDepartment = (e) => {
    setSearchDepartment(e.target.value);
  };

  return (
    <Form.Select onChange={handleSearchDepartment}>
      <option value="0">Tất cả các bộ môn</option>
      {listDepartments.map((department, index) => {
        return <option value={department.id} key={index}>{department.name}</option>
      })}
    </Form.Select>
  );
};

export default SearchDepartment;
