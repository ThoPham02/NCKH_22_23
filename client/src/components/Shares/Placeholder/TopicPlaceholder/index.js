import Placeholder from "react-bootstrap/Placeholder";

const TopicDetailPlaceholder = () => {
  return (
    <div className="topic-info">
      <div className="topic-info-item row">
        <div className="col-2">Tên đề tài</div>
        <div className="col-10">
          <Placeholder xs={8} />
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Đại hội</div>
        <div className="col-10">
          <Placeholder xs={6} />
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Tiểu ban</div>
        <div className="col-10">
          <Placeholder xs={5} />
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Giảng viên hướng dẫn</div>
        <div className="col-10">
          <Placeholder xs={3} />
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Liên hệ</div>
        <div className="col-10">
          <Placeholder xs={4} />
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Khoa</div>
        <div className="col-10">
          <Placeholder xs={3} />
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Bộ môn</div>
        <div className="col-10">
          <Placeholder xs={3} />
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Trạng thái</div>
        <div className="col-10">
          <Placeholder xs={2} />
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Thời gian thực hiện</div>
        <div className="col-10">
          <Placeholder xs={6} />
        </div>
      </div>
      <div className="topic-info-item row">
        <div className="col-2">Nhóm sinh viên</div>
        <div className="col-10">
          <Placeholder xs={8} />
        </div>
      </div>
    </div>
  );
};


export default TopicDetailPlaceholder;