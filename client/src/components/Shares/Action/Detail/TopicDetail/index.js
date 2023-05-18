import { convertTimestampToDateString } from "../../../../../utils/time";
import Loading from "../../../Loading";
import { topicStatus } from "../../../../../const/const";

const TopicDetail = (props) => {
  return (
    <div className="topic-info">
      <div className="topic-info-item row">
        <div className="col-2">Tên đề tài</div>
        <div className="col-10">{props.topicDetail.name}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Đại hội</div>
        <div className="col-10">{props.topicDetail.event}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Tiểu ban</div>
        <div className="col-10">
          {props.topicDetail.subcommittee
            ? props.topicDetail.subcommittee
            : "..."}
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Giảng viên hướng dẫn</div>
        <div className="col-10">{props.topic.lectureInfo.name}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Liên hệ</div>
        <div className="col-10">{props.topic.lectureInfo.phone}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Khoa</div>
        <div className="col-10">{props.topicDetail.faculty}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Bộ môn</div>
        <div className="col-10">{props.topicDetail.department}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Trạng thái</div>
        <div className="col-10">
          {topicStatus[props.topicDetail.status - 1]}
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Thời gian thực hiện</div>
        <div className="col-10">
          Từ {convertTimestampToDateString(props.topicDetail.timeStart)} đến{" "}
          {convertTimestampToDateString(props.topicDetail.timeEnd)}
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Nhóm sinh viên:</div>
        <div className="col-10 list-student">
          {props.topicDetail.listStudent ? (
            props.topicDetail.listStudent.map((item) => {
              return (
                <div key={item.id} className="student">
                  {item.name}
                </div>
              );
            })
          ) : (
            <span style={{fontWeight: "bold"}}>Chưa có ai đăng ký đề tài này</span>
          )}
          {props.status === "loading" ? <Loading /> : <></>}
        </div>
      </div>
    </div>
  );
};

export default TopicDetail;
