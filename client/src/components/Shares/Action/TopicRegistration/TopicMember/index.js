import CloseButton from 'react-bootstrap/CloseButton';
import "./style.css";

const Member = ({name, id, memberList, setMemberList}) => {
    const handleClick = () => {
        const list = memberList;
        console.log(list)
        const indexToDelete = list.findIndex(obj => obj.ID === id);
        list.splice(indexToDelete, 1);
        console.log(list)
        setMemberList(list)
    }

    return (
        <div className='member'>
            <div className='member-name'>{name}</div>
            <CloseButton onClick={handleClick}/>
        </div>
    )
}

export default Member;