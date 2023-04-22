import { useEffect, useState } from "react";
// import { AiOutlineLeft, AiOutlineRight } from "react-icons/ai";
import { MdOutlineKeyboardDoubleArrowRight } from "react-icons/md";
import { image1, image2, image3 } from "./imgHome";

import Block from "../../components/Shares/Block";
import "./home.css";

const listNotifications = [
  {
    name: "Trao bằng bổ sung cho sinh viên hoàn thành đầu ra đợt thi 12/2022",
    url: "#",
  },
  {
    name: "Huỷ nhóm học phần Tiếng Anh 3 và Thị giác máy",
    url: "#",
  },
  {
    name: "Đăng ký bổ sung khối lượng học tập học kỳ II năm học 2022-2023",
    url: "#",
  },
  {
    name: "Nghỉ học Nhân dịp kỉ niệm 92 năm ngày thành lập Đoàn TNCS Hồ Chí Minh",
    url: "#",
  },
  {
    name: "Triển khai kế hoạch thi khảo sát năng lực tiếng Anh bắt buộc đối với sinh viên hệ chính quy K65 (Đợt 2)",
    url: "#",
  },
];

const Home = () => {
  const images = [image1, image2, image3];
  const [currentImg, setCurrentImg] = useState(0);
  useEffect(() => {
    const intervalId = setInterval(() => {
      setCurrentImg((currentImg + 1) % images.length);
    }, 4000);
    return () => clearInterval(intervalId);
  });
  return (
    <Block title="Thông báo" className="block-noti">
      <div className="slide">
        <img src={images[currentImg]} alt="..." />
      </div>

      <div className="noti">
        {listNotifications.map((notification, index) => (
          <div>
            <a key={index} href={notification.url}>
              <MdOutlineKeyboardDoubleArrowRight />
              {notification.name}
            </a>
          </div>
        ))}
      </div>
    </Block>
  );

  // const [imageIndex, setImageIndex] = useState(0);

  // const prevImg = () => {
  //   setImageIndex(imageIndex === 0 ? images.length - 1 : imageIndex - 1);
  // };
  // const nextImg = () => {
  //   setImageIndex(imageIndex === images.length - 1 ? 0 : imageIndex + 1);

  //   const [pagi, setPagi] = useState(1);
  //   const [showPopUp, setShowPopUp] = useState(false);
  //   const handleButtonClick = () => {
  //     setShowPopUp(true);
  //   };
  //   const handleClosePopUp = () => {
  //     setShowPopUp(false);
  //   };

  //   return (
  //     <div className="home">
  //       <div>
  //         <p className="name__head">Trang chủ</p>
  //       </div>
  //       <div className="home__child">
  //         <div className="newIMG">
  //           <AiOutlineLeft className="prevImg dtImg" onClick={prevImg} />
  //           <img className="img--new" src={images[imageIndex]} alt="###" />
  //           <AiOutlineRight className="nextImg dtImg" onClick={nextImg} />
  //         </div>
  //         <div className="home__notification">
  //           <p>Thông báo</p>
  //         </div>
  //         <div className="home__prize">
  //           <p>giải thưởng</p>
  //           <ul>
  //             <li>
  //               <MdOutlineKeyboardDoubleArrowRight />
  //               <a href="###">giải thưởng 1</a>
  //             </li>
  //             <li>
  //               <MdOutlineKeyboardDoubleArrowRight />
  //               <a href="###">giải thưởng 2</a>
  //             </li>
  //             <li>
  //               <MdOutlineKeyboardDoubleArrowRight />
  //               <a href="###">giải thưởng 3</a>
  //             </li>
  //           </ul>
  //         </div>
  //       </div>

  //       <Pagination1
  //         currentPage={pagi}
  //         total={10}
  //         limit={2}
  //         setPagi={setPagi}
  //       />
  //       <div>
  //         <button onClick={handleButtonClick}>Mở cửa sổ</button>
  //         {showPopUp && (
  //           <PopUp
  //             heading="Thông báo"
  //             onClose={handleClosePopUp}
  //             content="this is content"
  //           ></PopUp>
  //         )}
  //       </div>
  //     </div>
  //   );
  // };
};

export default Home;
