import { useSelector } from "react-redux";
import { convertTimestampToDateString } from "../../../../../utils/time";
import Loading from "../../../Loading";
import { userSelector } from "../../../../../store/selectors";
import { Button } from "react-bootstrap";
import { getStatus } from "../../../../../utils/getStatus";
import { useDepartmentFaculty } from "../../../../../hooks/useDepartmentFaculty";

const TopicDetail = (props) => {
  const { topic, event, subcommittee, reports, marks, listStudent, isLoading, handleCancelButton, handleRegisButton } = props
  const user = useSelector(userSelector);
  const isAction =
    user.role === 0 || (user.role === 1 && listStudent.length <= topic.estimateStudent);
  const isCancel = !user.role === 0 && listStudent && listStudent.find(
    (item) => item.id === user.id
  );

  const { department, faculty } = useDepartmentFaculty(topic.departmentID)
  console.log(marks)
  console.log(reports)

  return (
    <div className="topic-info">
      <div className="topic-info-item row">
        <div className="col-2">Tên đề tài</div>
        <div className="col-10">{topic.name}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Đại hội</div>
        <div className="col-10">{event.name}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Tiểu ban</div>
        <div className="col-10">
          {subcommittee.name
            ? subcommittee.name
            : "..."}
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Giảng viên hướng dẫn</div>
        <div className="col-10">{topic.lectureInfo.name}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Liên hệ</div>
        <div className="col-10">{topic.lectureInfo.phone}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Khoa</div>
        <div className="col-10">{faculty}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Bộ môn</div>
        <div className="col-10">{department}</div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Trạng thái</div>
        <div className="col-10">
          {getStatus(topic.status)}
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Thời gian thực hiện</div>
        <div className="col-10">
          Từ {convertTimestampToDateString(topic.timeStart)} đến{" "}
          {convertTimestampToDateString(topic.timeEnd)}
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Nhóm sinh viên:</div>
        <div className="col-10 list-student">
          {listStudent ? (
            listStudent.map((item) => {
              return (
                <div key={item.studentID} className="student">
                  {item.name}
                </div>
              );
            })
          ) : (
            <span></span>
          )}
          {isAction ? (
            isCancel ? (
              <Button className="student" onClick={handleCancelButton} variant="danger">
                Hủy đăng ký
              </Button>
            ) : (
              <Button className="student" onClick={handleRegisButton} variant="success">
                Đăng ký
              </Button>
            )
          ) : (
            <></>
          )}
          {isLoading ? <Loading /> : <></>}
        </div>
      </div>
    </div>
  );
};

export default TopicDetail;
