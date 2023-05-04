import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { useState } from "react";

import "./style.css";
import Card from "../../../components/Shares/Card";
import SubCard from "../../../components/Shares/Card/SubCard";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import { SearchFaculty, SearchWord } from "../../../components/Shares/Search";
import useApi from "../../../hooks/useGetApi";
import Loading from "../../../components/Shares/Loading";
import PaginationCustom from "../../../components/Shares/Pagination";

const TopicRegis = () => {
  const [filter, setFilter] = useState({
    facultyId: 0,
    word: "",
    limit: 10
  });
  const [pagi, setPagi] = useState(0)
  let url = `/api/topic-registation?search=${filter.word}&facultyId=${filter.facultyId}&limit=${filter.limit}&offset=${pagi}`;
  const { data, isLoading, error, fetchData } = useApi(url);

  const handleSubmit = (e) => {
    e.preventDefault();
    fetchData();
  };

  var listTopic = [];

  if (data && data.total !== 0) {
    console.log(data.total)
    listTopic = data.topicRegistrations;
  }

  return (
    <div className="topic-regis">
      <Card title="Danh sách đề tài đề xuất">
        <SubCard title="Tìm kiếm">
          <Form className="search" onSubmit={handleSubmit}>
            <SearchWord value={filter.word} setFilter={setFilter}></SearchWord>
            <SearchFaculty
              value={filter.facultyId}
              setFilter={setFilter}
            ></SearchFaculty>
            <Form.Group>
              <Button variant="primary" type="submit" className="search-submit">
                {isLoading ? <Loading></Loading> : <></>}
                Tìm kiếm
              </Button>
            </Form.Group>
          </Form>
        </SubCard>
        <SubCard title="Danh sách">
          {listTopic.length == 0 ? (
            <EmptyListNoti title={"Không có đề tài nào!"} />
          ) : (
            <div>
              <table className="table">
                <thead className="table-head">
                  <tr className="table-tr">
                    <th>STT</th>
                    <th>Người hướng dẫn</th>
                    <th>Khoa</th>
                    <th>Số điện thoại liên hệ Email</th>
                    <th>Tên đề tài đề xuất</th>
                  </tr>
                </thead>
                <tbody>
                  {listTopic.map((item, index) => {
                    return (
                      <tr key={index}>
                        <td>{index + 1}</td>
                        <td>{item.lecture}</td>
                        <td>{item.faculty}</td>
                        <td>
                          {item.phone}
                          <br />
                          {item.email}
                        </td>
                        <td>{item.name}</td>
                      </tr>
                    );
                  })}
                </tbody>
              </table>

              <PaginationCustom setPagi={setPagi} currentPage={pagi} total={data.total} limit={filter.limit}/>
            </div>
          )}
        </SubCard>
      </Card>
    </div>
  );
};

export default TopicRegis;
