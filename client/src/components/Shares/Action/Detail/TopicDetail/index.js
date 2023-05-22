import { useSelector } from "react-redux";
import { convertTimestampToDateString } from "../../../../../utils/time";
import Loading from "../../../Loading";
import { topicStatus } from "../../../../../const/const";
import { userSelector } from "../../../../../store/selectors";
import { Button } from "react-bootstrap";

const TopicDetail = (props) => {
  const user = useSelector(userSelector);
  const isAction =
    user.role === 1 &&
    props.topicDetail.listStudent &&
    props.topicDetail.listStudent.length <= 5;
  const isCancel = props.topicDetail.listStudent &&  props.topicDetail.listStudent.find(
    (item) => item.id === user.id
  );

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
            <span></span>
          )}
          {isAction ? (
            isCancel ? (
              <Button className="student" onClick={props.handleCancelButton}>
                Hủy đăng ký
              </Button>
            ) : (
              <Button className="student" onClick={props.handleRegisButton}>
                Đăng ký
              </Button>
            )
          ) : (
            <></>
          )}
          {props.status === "loading" ? <Loading /> : <></>}
        </div>
      </div>
    </div>
  );
};

export default TopicDetail;
