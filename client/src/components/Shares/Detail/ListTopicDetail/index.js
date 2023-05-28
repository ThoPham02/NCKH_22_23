import {CgNametag} from "react-icons/cg";
import { FaUserGraduate, FaUsers } from "react-icons/fa";
import { AiFillPhone, AiOutlineMail, AiOutlineStar } from "react-icons/ai";

import "./style.css"

const ListTopicDetail = (props) => {
    const {name,lectureName,phone, email, estimateStudent, status} = props
    return (
        <>
            <div className="row">
                <div className="col-1 list-topic-icon"><CgNametag/></div>
                <div className="col-10">{name}</div>
            </div>
            <div className="row">
                <div className="col-1 list-topic-icon"><FaUserGraduate/></div>
                <div className="col-10">{lectureName}</div>
            </div>
            <div className="row">
                <div className="col-1 list-topic-icon"><AiFillPhone/></div>
                <div className="col-10">{phone}</div>
            </div>
            <div className="row">
                <div className="col-1 list-topic-icon"><AiOutlineMail/></div>
                <div className="col-10">{email}</div>
            </div>
            <div className="row">
                <div className="col-1 list-topic-icon"><AiOutlineStar/></div>
                <div className="col-10">{status}</div>
            </div>
            <div className="row">
                <div className="col-1 list-topic-icon"><FaUsers/></div>
                <div className="col-10">{estimateStudent}</div>
            </div>
        </>
    )
}

export default ListTopicDetail;