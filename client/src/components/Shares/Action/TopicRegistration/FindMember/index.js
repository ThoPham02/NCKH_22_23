import { FaCode, FaUser } from "react-icons/fa";
import "./style.css";

const FindMemberItem = ({ name, id, setMemberList, setFindName}) => {
  const handleStudentClick = () => {
    setMemberList(prev => [...prev, {name: name, id: id}])
    setFindName("")
  }
  return (
    <div className="find-member-item" onClick={handleStudentClick}>
      <div>
        <FaCode /> {id}
      </div>
      <div>
        <FaUser /> {name}
      </div>
    </div>
  );
};

export default FindMemberItem;
