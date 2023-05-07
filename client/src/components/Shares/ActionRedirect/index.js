import { Link } from "react-router-dom";
import { Button } from "react-bootstrap";

import "./style.css";

const ActionRedirect = ({ todo }) => {
  return (
    <div className="action">
      <button className="dots-container">
        <div className="dot"></div>
        <div className="dot"></div>
        <div className="dot"></div>
      </button>
      <div className="action-items">
        {todo.map((item, index) => {
          return (
            <div key={index} className="action-item">
              <Button className="button">
                <Link to={item.href} className="button">
                  {item.name}
                </Link>
              </Button>
            </div>
          );
        })}
        <div className="triangle"></div>
      </div>
    </div>
  );
};

export default ActionRedirect;
