import { useState } from "react";
import Pagination1 from "../../components/Shares/Pagination";
import PopUp from "../../components/Shares/PopUp";

const Home = () => {
  const [pagi, setPagi] = useState(1)
  const [showPopUp, setShowPopUp] = useState(false);
  const handleButtonClick = () => {
    setShowPopUp(true);
  };
  const handleClosePopUp = () => {
    setShowPopUp(false);
  };

  return (
    <div className="home">
      <Pagination1 currentPage={pagi} total={10} limit={2} setPagi={setPagi}/>
      <div>
      <button onClick={handleButtonClick}>Mở cửa sổ</button>
      {showPopUp && (
        <PopUp heading="Thông báo" onClose={handleClosePopUp} content="this is content">
        </PopUp>
      )}
    </div>
    </div>
  );
}

export default Home;
