import { useSelector } from "react-redux";

import "./userinfo.css"
import avtDefault from "./avata-default.jpg"
import { loginSelector } from "../../../store/selectors";

const UserInfo = () => {
    const userInfo = useSelector(loginSelector).info
    console.log(userInfo)    
    return (
        <div className="user">
            <div className="user__image">
                <img src={avtDefault} alt="" />
            </div>
        </div>
    )
}

export default UserInfo;