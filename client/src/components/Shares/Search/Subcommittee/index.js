import { useEffect } from "react";
import { Form } from "react-bootstrap";
import { useDispatch, useSelector } from "react-redux";
import { fetchCommonSubcommittee } from "./CommonSubcommitteeSlice";
import { CommonSubcommitteeSelector } from "../../../../store/selectors";

const Subcommittee = ({subcommittee, setSubcommittee, facultyID}) => {
    const dispatch = useDispatch();
  const list = useSelector(CommonSubcommitteeSelector).subcommittees;
  useEffect(() => {
    dispatch(fetchCommonSubcommittee({facultyID}));
  }, [dispatch, facultyID]);

  return (
    <Form.Group className="col-12 col-sm-12 col-md-6 col-lg-3">
      <Form.Select value={subcommittee} onChange={(e) => setSubcommittee(e.target.value)}>
        <option value={0}>Tất cả Tiểu ban</option>
        {list ? list.map((item, index) => {
          return (
            <option value={item.id} key={index}>
              {item.name}
            </option>
          );
        }) : <></>}
      </Form.Select>
    </Form.Group>
  );
};

export default Subcommittee;