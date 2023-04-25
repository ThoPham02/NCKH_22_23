// import { useEffect, useState } from "react";
import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";

import { fetchFaculity } from "./FaculitySlice";
import { facultySelector } from "../../../../store/selectors";
import { useEffect } from "react";

const Faculty = ({value, setFilter }) => {
  const dispatch = useDispatch();
  const list = useSelector(facultySelector);
  useEffect(() => {
    dispatch(fetchFaculity());
  }, [dispatch]);


  const handleSearchFaculity = (e) => {
    setFilter(prevState => ({
      ...prevState,
      facultyId: e.target.value,
      departmentId: 0
    }));
  };
  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select onChange={handleSearchFaculity} value={value}>
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
