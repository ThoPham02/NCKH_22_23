import { useState } from "react";
import Placeholder from "react-bootstrap/Placeholder";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";

import "./style.css";
import Member from "./TopicMember";
import useApi from "../../../../hooks/useGetApi";
import FindMemberItem from "./FindMember";

const TopicRegistration = ({ name, topicRegis }) => {
  const [show, setShow] = useState(false);
  const [memberList, setMemberList] = useState([]);
  const [findName, setFindName] = useState("");
  let url = `/api/student?name=${findName}`;
  const { data, isLoading } = useApi(url);

  let StudetList = [];
  if (!isLoading && data != undefined && data.studentList != undefined) {
    StudetList = data.studentList;
  }

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);
  const handleSubmit = (e) => {
    e.preventDefault();
    
  };
  const handleInputChange = (e) => {
    setFindName(e.target.value);
  };

  return (
    <>
      <Button className="button" onClick={handleShow}>
        {name}
      </Button>

      <Modal show={show} onHide={handleClose} size="xl">
        <Modal.Header closeButton>
          <Modal.Title>Đăng ký đề tài</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <div className="topic-regis-body">
            <div className="row">
              <div className="col-3 topic-regis-title">Tên đề tài:</div>
              <div className="col-8">{topicRegis.name}</div>
            </div>
            <div className="row">
              <div className="col-3 topic-regis-title">
                Giảng viên hướng dẫn:
              </div>
              <div className="col-8">{topicRegis.lecture}</div>
            </div>
            <div className="row">
              <div className="col-3 topic-regis-title">Sự kiện:</div>
              <div className="col-8">
                Đại hội nghiên cứu khoa học lần thứ 36
              </div>
            </div>
            <div className="row">
              <div className="col-3 topic-regis-title">Các thành viên: </div>
              <div className="col-8 find-member">
                <form className="member-form" onSubmit={handleSubmit}>
                  <div className="members">
                    {memberList.map((item, index) => (
                      <Member
                        key={index}
                        id={item.id}
                        name={item.name}
                        memberList={memberList}
                        setMemberList={setMemberList}
                      />
                    ))}
                  </div>
                  <input
                    type="text"
                    placeholder="Tìm kiếm thành viên theo tên"
                    onChange={handleInputChange}
                    value={findName}
                  />
                </form>

                <div>
                  {findName === "" ? (
                    <></>
                  ) : (
                    <div className="find-member-list">
                      {isLoading ? (
                        <FindMemberItem
                          name={<Placeholder xs={4} />}
                          id={<Placeholder xs={6} />}
                        />
                      ) : StudetList.length === 0 ? (
                        <span>Không có sinh viên nào!</span>
                      ) : (
                        StudetList.map((item) => (
                          <FindMemberItem
                            key={item.id}
                            name={item.name}
                            id={item.id}
                            setMemberList={setMemberList}
                            setFindName={setFindName}
                          />
                        ))
                      )}
                    </div>
                  )}
                </div>
              </div>
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" type="submit">
            Xác Nhận
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
};

export default TopicRegistration;
