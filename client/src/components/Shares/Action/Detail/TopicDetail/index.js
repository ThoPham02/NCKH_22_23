
import "./style.css"
import { getStatus } from "../../../../../utils/getStatus";
import { convertTimestampToDateString } from "../../../../../utils/time";
import { useDepartmentFaculty } from "../../../../../hooks/useDepartmentFaculty";
import SubCard from "../../../Card/SubCard";
import { Badge } from "react-bootstrap";

const TopicDetail = (props) => {
  // eslint-disable-next-line
  const { status, topic, event, subcommittee, reports, marks, listStudent } = props.data

  const { department, faculty } = useDepartmentFaculty(topic.departmentID)
  // const isLoading = status === 'loading'
  const isRegis = true
  const isCancel = true
  return (
    <div className="topic-info">
      <SubCard title={"Thông tin Chung: "}>
        <div className="row info-title">
          <div className="col-2">NCKH</div>
          <div className="col-10">{event.name}</div>
        </div>
        <div className="row info-title">
          <div className="col-2">Khoa</div>
          <div className="col-10">{faculty}</div>
        </div>
        <div className="row info-title">
          <div className="col-2">Bộ môn</div>
          <div className="col-10">{department}</div>
        </div>
      </SubCard>
      <SubCard title={"Thông tin đề tài:"}>
        <div className=" row info-title">
          <div className="col-2">Tên đề tài</div>
          <div className="col-10">{topic.name}</div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Giảng viên hướng dẫn</div>
          <div className="col-10">{topic.lectureInfo.degree + "." + topic.lectureInfo.name}</div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Liên hệ</div>
          <div className="col-10">{topic.lectureInfo.phone}</div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Trạng thái</div>
          <div className="col-10">
            {getStatus(topic.status)}
          </div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Thời gian thực hiện</div>
          <div className="col-10">
            Từ {convertTimestampToDateString(topic.timeStart)} đến{" "}
            {convertTimestampToDateString(topic.timeEnd)}
          </div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Số thành viên dự kiến</div>
          <div className="col-10">
            <Badge bg="success" style={{marginBottom: "4px"}}>
              {topic.estimateStudent}
            </Badge>
          </div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Nhóm sinh viên</div>
          <div className="col-10 list-student">
            {listStudent ? (
              listStudent.map((item) => {
                return (
                  <div key={item.studentID} className="student">
                    {item.name}
                  </div>
                );
              })
            ) : <></>}
            {isRegis ? <></> : <></>}
            {isCancel ? <></> : <></>}
          </div>
        </div>
      </SubCard>
      <SubCard title={"Nghiệm thu:"}>
        <div className=" row info-title">
          <div className="col-2">Tiểu ban</div>
          <div className="col-10">
            {subcommittee.name
              ? subcommittee.name
              : "..."}
          </div>
        </div>
      </SubCard>

      <SubCard title={"Báo cáo:"}>

      </SubCard>
      <SubCard title={"Kết quả:"}>

      </SubCard>
    </div>
  );
};

export default TopicDetail;
