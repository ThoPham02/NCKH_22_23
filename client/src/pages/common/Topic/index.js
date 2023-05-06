import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Table from "react-bootstrap/Table";
import { useState } from "react";

import "./style.css";
import Card from "../../../components/Shares/Card";
import SubCard from "../../../components/Shares/Card/SubCard";
import {
  SearchWord,
  SearchDateFrom,
  SearchDateTo,
  SearchDepartment,
  SearchFaculty,
  SearchStatus,
} from "../../../components/Shares/Search";
import useApi from "../../../hooks/useGetApi";

import "./style.css";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import Loading from "../../../components/Shares/Loading";
import PaginationCustom from "../../../components/Shares/Pagination";
import Action from "../../../components/Shares/Action";

const Topic = () => {
  const [filter, setFilter] = useState({
    departmentId: 0,
    facultyId: 0,
    word: "",
    status: 0,
    dateTo: "",
    dateFrom: "",
    limit: 10,
  });
  const [pagi, setPagi] = useState(1);

  let url = `/api/topic?search=${filter.word}`;
  const { data, isLoading, error, fetchData } = useApi(url);
  let listTopic = [];
  if (data && data.total !== 0) {
    listTopic = data.listTopics;
  }

  return (
    <div className="topic">
      <Card title="Danh sách đề tài">
        <SubCard title="Tìm kiếm">
          <Form className="search">
            {isLoading ? <Loading></Loading> : <></>}
            <SearchWord value={filter.word} setFilter={setFilter}></SearchWord>
            <SearchDepartment
              faculty={filter.facultyId}
              value={filter.departmentId}
              setFilter={setFilter}
            ></SearchDepartment>
            <SearchFaculty
              value={filter.facultyId}
              setFilter={setFilter}
            ></SearchFaculty>
            <SearchStatus
              value={filter.status}
              setFilter={setFilter}
            ></SearchStatus>
            <SearchDateFrom
              value={filter.dateFrom}
              setFilter={setFilter}
            ></SearchDateFrom>
            <SearchDateTo
              value={filter.dateTo}
              setFilter={setFilter}
            ></SearchDateTo>
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
                        <td></td>
                        <td>
                          <Action
                            todo={[
                              { name: "Xem chi tiết", href: "#" },
                              { name: "Đăng ký", href: "#" },
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
                total={data.total}
                limit={filter.limit}
              />
            </div>
          )}
        </SubCard>
      </Card>
    </div>
  );
};

export default Topic;
