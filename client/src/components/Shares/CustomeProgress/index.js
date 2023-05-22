import ProgressBar from "react-bootstrap/ProgressBar";
import { useSelector } from "react-redux";
import { stageSelector } from "../../../store/selectors";

import "./style.css";

const CustomeProgress = ({ currentStage, setCurrentStage }) => {
  let now = 0;

  const data = useSelector(stageSelector)

  return (
    <div className="custom_progress">
      {data.map((item, index) => {
        if (item.id === currentStage) {
          now = (index / (data.length - 1)) * 100;
        }
        return (
          <div
            key={index}
            className="progress_item"
            style={{
              left: `calc(${(index / (data.length - 1)) * 95}% - 42px)`,
              cursor: "pointer",
            }}
            onClick={() => setCurrentStage(item.id)}
          >
            <div className="drop-shape"></div>
            <div className="progress_name">{item.name}</div>
          </div>
        );
      })}
      <ProgressBar
        now={now}
        label={`${now}%`}
        visuallyHidden
        style={{ width: "95%", position: "absolute" }}
      />
    </div>
  );
};

export default CustomeProgress;
