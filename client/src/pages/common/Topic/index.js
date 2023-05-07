import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Table from "react-bootstrap/Table";
import { useRef, useState } from "react";

import "./style.css";
import Card from "../../../components/Shares/Card";
import SubCard from "../../../components/Shares/Card/SubCard";
import {
  SearchWord,
  SearchDateFrom,
  SearchDateTo,
  SearchFaculty,
  SearchStatus,
} from "../../../components/Shares/Search";
import useApi from "../../../hooks/useGetApi";

import "./style.css";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import Loading from "../../../components/Shares/Loading";
import PaginationCustom from "../../../components/Shares/Pagination";
import Action from "../../../components/Shares/Action";
import { LIMIT } from "../../../const/const";
import TopicInfo from "../../../components/Shares/TopicInfo";
import Detail from "../../../components/Shares/Action/Detail";

const Topic = () => {
  const searchRef = useRef("");
  const facultyRef = useRef(0);
  const statusRef = useRef(0);
  const dateFromRef = useRef("");
  const dateToRef = useRef("");
  const [pagi, setPagi] = useState(1);
  const [url, setUrl] = useState(
    `/api/topic?limit=${LIMIT}&offset=${pagi - 1}`
  );

  const { data, isLoading } = useApi(url);
  let listTopic = [];
  if (data && data.total !== 0) {
    listTopic = data.listTopics;
  }

  const handleSubmitForm = (e) => {
    e.preventDefault();
    let search = searchRef.current.value;
    let status = statusRef.current.value;
    let dateFrom = dateFromRef.current.value;
    let dateTo = dateToRef.current.value;
    let pagiNum = pagi - 1;
    setUrl(
      `/api/topic?limit=${LIMIT}&offset=${pagiNum}&search=${search}&status=${status}&dateFrom=${dateFrom}&dateTo=${dateTo}`
    );
  };

  return (
    <div className="topic">
      <Card title="Danh sách đề tài">
        <SubCard title="Tìm kiếm">
          <Form className="search" onSubmit={handleSubmitForm}>
            <SearchWord searchRef={searchRef}></SearchWord>
            <SearchFaculty facultyRef={facultyRef}></SearchFaculty>
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

        <SubCard title={"Danh sách"}>
          {listTopic.length === 0 ? (
            <EmptyListNoti title={"Không có đề tài nào!"} />
          ) : (
            <div>
              <Table bordered hover size="sm">
                <thead>
                  <tr>
                    <th>Mã đề tài</th>
                    <th>Thông tin đề tài</th>
                    <th>Thao tác</th>
                  </tr>
                </thead>
                <tbody>
                  {listTopic.map((item, index) => {
                    return (
                      <tr key={index}>
                        <td>{item.id}</td>
                        <td>
                          <TopicInfo
                            name={item.name}
                            status={item.status}
                            lecture={item.lecture}
                            dateTo={item.timeStart}
                            dateFrom={item.timeEnd}
                            fileUrl={item.resultUrl}
                          />
                        </td>
                        <td>
                          <Action todo={[<Detail name={"Xem chi tiết"} topicID={item.id}/>]} />
                        </td>
                      </tr>
                    );
                  })}
                </tbody>
              </Table>

              <PaginationCustom
                setPagi={setPagi}
                currentPage={pagi}
                total={data.total}
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
