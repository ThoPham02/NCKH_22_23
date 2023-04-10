import { useEffect, useState } from "react";
import Form from "react-bootstrap/Form";
import { useDispatch, useSelector } from "react-redux";
import { faculitySelector } from "../../../../store/selectors";
import { fetchFaculity } from "./FaculitySlice";
import { FilterActions } from "../SearchSlice";

const SearchFaculity = ({ setSearchFaculity }) => {
  const dispatch = useDispatch();
  const list = useSelector(faculitySelector);
  const [listFaculties, setListFaculties] = useState(list);
  useEffect(() => {
    dispatch(fetchFaculity());
  }, [dispatch]);

  useEffect(() => {
    setListFaculties(list);
  }, [list]);

  const handleSearchFaculity = (e) => {
    setSearchFaculity(e.target.value);
    dispatch(FilterActions.searchFaculityChange(e.target.value));
  };
  return (
    <Form.Select onChange={handleSearchFaculity}>
      <option value="0">Tất cả các khoa</option>
      {listFaculties.map((faculity) => {
        console.log(faculity)
        return <option value={faculity.id}>{faculity.name}</option>;
      })}
    </Form.Select>
  );
};

export default SearchFaculity;
