import "./header.css"
import logo from "./lg.jpg"
import {Link} from "react-router-dom";

import {BiUserCircle} from "react-icons/bi"
function Header() {
    return (
        <div id="header">
            <div id="header-logo">
                <div id ="logo">
                    <img src={logo} alt="###"/>
                    <ul className ="logo-name">
                        <li className="name-VN">Trường đại học mỏ - địa chất</li>
                        <li className="name-EN">hanoi university of mining and geology</li>
                    </ul>
                </div>
                <div id = "header-logo__right">
                    <p>hệ thống quản lí <br></br>nghiên cứu khoa học</p>
                </div>
            </div>
            <div id ="nav">
                <ul>
                    <li><Link to ="/">đề tài</Link></li>
                    <li><Link to ="/">thống kê</Link></li>
                    <li><Link to ="/">danh mục</Link></li>
                    <li><Link to ="/">liên hệ</Link></li>
                </ul>
                <div id="acc">
                    <Link to="/" id ="acc_1">
                        <BiUserCircle></BiUserCircle>
                    </Link>
                    <Link to="/"  id = "acc_2">
                        Đăng nhập
                    </Link>
                </div>
            </div>
        </div>
    )
}

export default Header;