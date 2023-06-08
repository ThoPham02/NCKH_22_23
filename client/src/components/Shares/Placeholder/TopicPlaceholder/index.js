import Placeholder from "react-bootstrap/Placeholder";
import SubCard from "../../Card/SubCard";

const TopicDetailPlaceholder = () => {
  return (
    <div className="topic-info">
      <SubCard title={"Thông tin Chung: "}>
        <div className="row info-title">
          <div className="col-2">NCKH</div>
          <div className="col-10"><Placeholder xs={5} /></div>
        </div>
        <div className="row info-title">
          <div className="col-2">Khoa</div>
          <div className="col-10"><Placeholder xs={4} /></div>
        </div>
        <div className="row info-title">
          <div className="col-2">Bộ môn</div>
          <div className="col-10"><Placeholder xs={3} /></div>
        </div>
      </SubCard>
      <SubCard title={"Thông tin đề tài:"}>
        <div className=" row info-title">
          <div className="col-2">Tên đề tài</div>
          <div className="col-10"><Placeholder xs={8} /></div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Giảng viên hướng dẫn</div>
          <div className="col-10"><Placeholder xs={3} /></div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Liên hệ</div>
          <div className="col-10"><Placeholder xs={4} /></div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Trạng thái</div>
          <div className="col-10">
            <Placeholder xs={2} />
          </div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Thời gian thực hiện</div>
          <div className="col-10">
            <Placeholder xs={8} />
          </div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Số thành viên dự kiến</div>
          <div className="col-10">
            <Placeholder xs={1} />
          </div>
        </div>
        <div className=" row info-title">
          <div className="col-2">Nhóm sinh viên</div>
          <div className="col-10 list-student">
            <Placeholder xs={6} />
          </div>
        </div>
      </SubCard>
    </div>
  );
};


export default TopicDetailPlaceholder;