import ProgressBar from "react-bootstrap/ProgressBar";

import "./style.css";

const CustomeProgress = ({ stage, setStage, data }) => {
  let now = 0
  let current = new Date().getTime();
  for (let i = 0; i < data.length; i++) {
    if (data[i].timeStart !== 0 && data[i].timeStart <= current) {
      now = i
    }
  }
  let persent = (now / (data.length - 1)) * 100;
  return (
    <div className="custom_progress">
      <div className="curentStage" style={{ left: `calc(${(data.indexOf(stage) / (data.length - 1)) * 95}% - 5px)` }}></div>
      {data.map((item, index) => {
        return (
          <div
            key={index}
            className="progress_item"
            style={{ left: `calc(${(index / (data.length - 1)) * 95}% - 42px)` }}
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
