
import { Button } from "react-bootstrap";
import "./style.css";
import Confirm from "../../Confirm";


const ConfirmStatus = ({ name }) => {
  return (
    <>
      <Button className="button">
        {name}
      </Button>

    <Confirm />
    </>
  );
};
export default ConfirmStatus;
