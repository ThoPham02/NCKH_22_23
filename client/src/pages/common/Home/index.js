import { MdOutlineKeyboardDoubleArrowRight } from "react-icons/md";
import { useEffect, useState } from "react";

import Card from "../../../components/Shares/Card";
import { image1, image2, image3 } from "./imgHome";
import { listNotifications, instructions } from "./data";
import "./style.css";

const Home = () => {
  const images = [image1, image2, image3];
  const [curent, setCurent] = useState(0);
  useEffect(() => {
    const intervalId = setInterval(() => {
      setCurent((curent + 1) % images.length);
    }, 4000);
    return () => clearInterval(intervalId);
  });
  return (
    <div className="home">
      <Card title="Thông Báo">
        <div  className="notification">
          <div className="notification-img">
            <img src={images[curent]} alt="..."/>
          </div>
          <div className="notification-content">
            {listNotifications.map((notification, index) => {
              return (
                <div key={index}>
                    <MdOutlineKeyboardDoubleArrowRight/>
                  <a href={notification.url}>
                    {notification.name}
                  </a>
                </div>
              );
            })}
          </div>
        </div>
      </Card>

      <Card title="Hướng dẫn">
      <div className="notification-content">
            {instructions.map((notification, index) => {
              return (
                <div key={index}>
                    <MdOutlineKeyboardDoubleArrowRight />
                  <a href={notification.url}>
                    {notification.name}
                  </a>
                </div>
              );
            })}
          </div>
      </Card>
    </div>
  );
};

export default Home;
