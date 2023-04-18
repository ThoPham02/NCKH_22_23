import { useState, useEffect } from "react";
import "./home.css";
import { image1, image2, image3 } from "./imgHome";
import { AiOutlineLeft, AiOutlineRight } from "react-icons/ai";
import { MdOutlineKeyboardDoubleArrowRight } from "react-icons/md";
const Home = () => {
  const [imageIndex, setImageIndex] = useState(0);
  const images = [image1, image2, image3];
  useEffect(() => {
    const intervalId = setInterval(() => {
      setImageIndex((imageIndex + 1) % images.length);
    }, 4000);
    return () => clearInterval(intervalId);
  });

  const prevImg = () => {
    setImageIndex(imageIndex === 0 ? images.length - 1 : imageIndex - 1);
  };
  const nextImg = () => {
    setImageIndex(imageIndex === images.length - 1 ? 0 : imageIndex + 1);

    const [pagi, setPagi] = useState(1);
    const [showPopUp, setShowPopUp] = useState(false);
    const handleButtonClick = () => {
      setShowPopUp(true);
    };
    const handleClosePopUp = () => {
      setShowPopUp(false);
    };

    return (
      <div className="home">
        <div>
          <p className="name__head">Trang chủ</p>
        </div>
        <div className="home__child">
          <div className="newIMG">
            <AiOutlineLeft className="prevImg dtImg" onClick={prevImg} />
            <img className="img--new" src={images[imageIndex]} alt="###" />
            <AiOutlineRight className="nextImg dtImg" onClick={nextImg} />
          </div>
          <div className="home__notification">
            <p>Thông báo</p>
          </div>
          <div className="home__prize">
            <p>giải thưởng</p>
            <ul>
              <li>
                <MdOutlineKeyboardDoubleArrowRight />
                <a href="###">giải thưởng 1</a>
              </li>
              <li>
                <MdOutlineKeyboardDoubleArrowRight />
                <a href="###">giải thưởng 2</a>
              </li>
              <li>
                <MdOutlineKeyboardDoubleArrowRight />
                <a href="###">giải thưởng 3</a>
              </li>
            </ul>
          </div>
        </div>

        <Pagination1
          currentPage={pagi}
          total={10}
          limit={2}
          setPagi={setPagi}
        />
        <div>
          <button onClick={handleButtonClick}>Mở cửa sổ</button>
          {showPopUp && (
            <PopUp
              heading="Thông báo"
              onClose={handleClosePopUp}
              content="this is content"
            ></PopUp>
          )}
        </div>
      </div>
    );
  };
};

export default Home;
