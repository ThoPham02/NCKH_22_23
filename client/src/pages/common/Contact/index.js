import Card from "../../../components/Shares/Card";
import "./style.css";

const Contact = () => {
  return (
    <Card title="Thông tin liên hệ">
      <div className="contact-content">
        <h4>Mọi chi tiết xin gửi về:</h4>
        <p>Phòng Khoa học Công nghệ (P307, nhà C 12 tầng)</p>
        <p>Điện thoại: <br/>0243 8386437 </p>
        <p>Email: khoahoccongnghe@humg.edu.vn</p>
      </div>
    </Card>
  );
};

export default Contact;
