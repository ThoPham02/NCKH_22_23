import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Table from "react-bootstrap/Table";
import { useEffect, useRef, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import Accordion from "react-bootstrap/Accordion";
import InputGroup from "react-bootstrap/InputGroup";

import "./style.css";
import Card from "../../../components/Shares/Card";
import SubCard from "../../../components/Shares/Card/SubCard";
import {
  SearchWord,
  SearchDateFrom,
  SearchDateTo,
  SearchFaculty,
  SearchStatus,
  SearchDepartment,
} from "../../../components/Shares/Search";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import Loading from "../../../components/Shares/Loading";
import PaginationCustom from "../../../components/Shares/Pagination";
import Action from "../../../components/Shares/Action";
import { LIMIT } from "../../../const/const";
import Detail from "../../../components/Shares/Action/Detail";
import { currentEventSelector, topicSelector } from "../../../store/selectors";
import { convertDateToTimestamp } from "../../../utils/time";
import { fetchTopics } from "../../common/Topic/TopicSlice";
import { userSelector } from "../../../store/selectors";
import { createTopic, fetchCurrentEvent } from "./LectureSlice";
import Confirm from "../../../components/Shares/Confirm";

const Topic = () => {
  const [faculty, setFaculty] = useState(0);
  const [pagi, setPagi] = useState(1);
  const [confirmShow, setConfirmShow] = useState(false);
  const dispatch = useDispatch();
  const addRef = useRef(0);
  const nameRef = useRef("");
  const searchRef = useRef("");
  const statusRef = useRef(0);
  const dateFromRef = useRef("");
  const dateToRef = useRef("");
  const departmentRef = useRef(0);
  useEffect(() => {
    dispatch(
      fetchTopics({
        search: searchRef.current.value,
        departmentID: departmentRef.current.value,
        facultyID: faculty,
        status: statusRef.current.value,
        timeStart: dateToRef.current.value,
        timeEnd: dateToRef.current.value,
        limit: LIMIT,
        offset: (pagi - 1) * LIMIT,
      })
    );
    // eslint-disable-next-line
  }, [dispatch, pagi]);

  useEffect(() => {
    dispatch(fetchCurrentEvent());
  }, [dispatch]);

  const handleSubmitForm = (e) => {
    e.preventDefault();
    let search = searchRef.current.value;
    let status = statusRef.current.value;
    let dateFrom = dateFromRef.current.value;
    let dateTo = dateToRef.current.value;
    let department = departmentRef.current.value;

    setPagi(1);
    dispatch(
      fetchTopics({
        search: search,
        departmentID: department,
        facultyID: faculty,
        status: status,
        timeStart: convertDateToTimestamp(dateTo),
        timeEnd: convertDateToTimestamp(dateFrom),
        limit: LIMIT,
        offset: (pagi - 1) * LIMIT,
      })
    );
  };

  const handleSubmitAddTopic = (e) => {
    e.preventDefault();

    setConfirmShow(true);
  };

  const handleConfirmButton = () => {
    dispatch(
      createTopic({
        name: nameRef.current.value,
        departmentID: addRef.current.value * 1,
        eventID: current.id,
        lectureID: user.id,
      })
    );

    setConfirmShow(false);
  };

  const topic = useSelector(topicSelector);
  const current = useSelector(currentEventSelector);

  let isLoading = topic.status === "loading";
  const listTopic = topic.topics;
  let total = topic.total;
  const user = useSelector(userSelector);
  return (
    <div className="topic">
      <Card title="Danh sách đề tài">
        <Accordion defaultActiveKey="0">
          <Accordion.Item eventKey="0" style={{ margin: "0 20px" }}>
            <Accordion.Header className="add-topic-header">
              Thêm đề tài
            </Accordion.Header>
            <Accordion.Body className="add-topic-body">
              <form onSubmit={handleSubmitAddTopic}>
                <div className="row">
                  <div className="col-3">Tên đề tài: </div>
                  <div className="col-9">
                    <InputGroup>
                      <Form.Control ref={nameRef} />
                    </InputGroup>
                  </div>
                </div>
                <div className="row">
                  <div className="col-3">Bộ môn: </div>
                  <div className="col-9">
                    <SearchDepartment
                      faculty={user.faculty_id}
                      departmentRef={addRef}
                      style={{ width: "400px" }}
                      defaultValue={"Chọn bộ môn"}
                    />
                  </div>
                </div>
                <div className="row">
                  <div className="col-3">Đại hội: </div>
                  <div className="col-9">
                    {current.name ? current.name : ""}
                  </div>
                </div>
                <div className="row">
                  <div className="col-3">Giảng viên hướng dẫn:</div>
                  <div className="col-9">{user.name}</div>
                </div>
                <div className="row">
                  <div className="col-3"></div>
                  <div className="col-9">
                    <Button
                      variant="primary"
                      type="submit"
                      style={{ padding: "0 8px" }}
                    >
                      Xác nhận
                    </Button>
                  </div>
                </div>
              </form>

              <Confirm
                content={"Xác nhận đăng ký đề tài"}
                confirmShow={confirmShow}
                setConfirmShow={setConfirmShow}
                handleConfirmButton={handleConfirmButton}
              />
            </Accordion.Body>
          </Accordion.Item>

          <Accordion.Item eventKey="1" style={{ margin: "0 20px" }}>
            <Accordion.Header className="add-topic-header">
              Danh sách đề tài
            </Accordion.Header>
            <Accordion.Body className="add-topic-body">
              <SubCard title="Tìm kiếm">
                <Form className="search" onSubmit={handleSubmitForm}>
                  <SearchWord searchRef={searchRef}></SearchWord>
                  <SearchFaculty
                    faculty={faculty}
                    setFaculty={setFaculty}
                  ></SearchFaculty>
                  <SearchDepartment
                    departmentRef={departmentRef}
                    faculty={faculty}
                    defaultValue={"Tất cả bộ môn"}
                  ></SearchDepartment>
                  <SearchStatus statusRef={statusRef}></SearchStatus>
                  <SearchDateFrom dateFromRef={dateFromRef}></SearchDateFrom>
                  <SearchDateTo dateToRef={dateToRef}></SearchDateTo>
                  <Form.Group>
                    <Button
                      variant="primary"
                      type="submit"
                      className="search-submit"
                    >
                      {isLoading ? <Loading></Loading> : <></>}
                      Tìm kiếm
                    </Button>
                  </Form.Group>
                </Form>
              </SubCard>
              <SubCard>
                {total === 0 ? (
                  <EmptyListNoti title={"Không có đề tài nào!"} />
                ) : (
                  <div>
                    <Table bordered hover size="sm" className="topic-table">
                      <thead>
                        <tr>
                          <th>STT</th>
                          <th style={{ width: "200px" }}>Giảng viên</th>
                          <th>Liên hệ</th>
                          <th>Đề tài</th>
                          <th>Thao tác</th>
                        </tr>
                      </thead>
                      <tbody>
                        {listTopic.map((item, index) => {
                          return (
                            <tr key={index}>
                              <td>{(pagi - 1) * LIMIT + index + 1}</td>
                              <td>
                                {item.lectureInfo.degree +
                                  " " +
                                  item.lectureInfo.name}
                              </td>
                              <td style={{ textAlign: "center" }}>
                                {item.lectureInfo.email}
                                <br />
                                {item.lectureInfo.phone}
                              </td>
                              <td>{item.name}</td>
                              <td>
                                <Action
                                  todo={[
                                    <Detail
                                      name={"Xem chi tiết"}
                                      topic={item}
                                    />,
                                  ]}
                                />
                              </td>
                            </tr>
                          );
                        })}
                      </tbody>
                    </Table>

                    <PaginationCustom
                      setPagi={setPagi}
                      currentPage={pagi}
                      total={total}
                      limit={LIMIT}
                    />
                  </div>
                )}
              </SubCard>
            </Accordion.Body>
          </Accordion.Item>
        </Accordion>
      </Card>
    </div>
  );
};

export default Topic;
