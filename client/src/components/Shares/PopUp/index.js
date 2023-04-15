import React, { useState } from "react";
import "./popup.css";


const PopUp = ({ heading, onClose, content }) => {
  const [isOpen, setIsOpen] = useState(true);

  const handleClose = () => {
    setIsOpen(false);
    onClose();
  };

  return (
    <>
      {isOpen && (
        <div className="popup-container">
          <div className="popup-header">
            <h3>{heading}</h3>
            <button className="close-btn" onClick={handleClose}>
              x
            </button>
          </div>
          <div className="popup-content">
            <p>{content}</p>
          </div>
          <div className="popup-footer">
            <button className="close-btn" onClick={handleClose}>
              Close
            </button>
          </div>
        </div>
      )}
    </>
  );
};



export default PopUp;