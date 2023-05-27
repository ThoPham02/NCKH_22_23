import ProgressBar from "react-bootstrap/ProgressBar";

import "./style.css";
import { current } from "@reduxjs/toolkit";

const CustomeProgress = ({ stage, setStage, data }) => {
  let persent = 0;
  setStage()

  let now = new Date().getSeconds()
  const currentStage = data.find(item => item.timeStart <= now && item.timeEnd >= now) || data[0]


  return (
    <div className="custom_progress">
      <div className="curentStage" style={{}}></div>
      {data.map((item, index) => {
        if (item.id === stage.id) {
          persent = (index / (data.length - 1)) * 100;
        }
        return (
          <div
            key={index}
            className="progress_item"
            style={{left: `calc(${(index / (data.length - 1)) * 95}% - 42px)`}}
            onClick={() => setStage(item)}
          >
            <div className="drop-shape"></div>
            <div className="progress_name">{item.name}</div>
          </div>
        );
      })}
      <ProgressBar
        now={persent}
        label={`${persent}%`}
        visuallyHidden
        style={{ width: "95%", position: "absolute" }}
      />
    </div>
  );
};

export default CustomeProgress;
