import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Table from "react-bootstrap/Table";
import { useEffect, useRef, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

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
import { topicSelector } from "../../../store/selectors";
import { convertDateToTimestamp } from "../../../utils/time";
import { fetchTopics } from "../../common/Topic/TopicSlice";

const Topic = () => {
  const [faculty, setFaculty] = useState(0);
  const [pagi, setPagi] = useState(1);
  const dispatch = useDispatch();
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

  const handleSubmitForm = (e) => {
    e.preventDefault();
    let search = searchRef.current.value;
    let status = statusRef.current.value;
    let dateFrom = dateFromRef.current.value;
    let dateTo = dateToRef.current.value;
    let department = departmentRef.current.value;

    console.log(dateTo, convertDateToTimestamp(dateTo));

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

  const topic = useSelector(topicSelector);

  let isLoading = topic.status === "loading";
  const listTopic = topic.topics;
  let total = topic.total;

  return (
    <div className="topic">
      <Card title="Danh sách đề tài">
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
              <Button variant="primary" type="submit" className="search-submit">
                {isLoading ? <Loading></Loading> : <></>}
                Tìm kiếm
              </Button>
            </Form.Group>
          </Form>
        </SubCard>

        <SubCard title={"Danh sách đề tài"}>
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
                              <Detail name={"Xem chi tiết"} topic={item} />,
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
      </Card>
    </div>
  );
};

export default Topic;
