import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Table from "react-bootstrap/Table";
import { useState, useRef } from "react";

import "./style.css";
import Card from "../../../components/Shares/Card";
import SubCard from "../../../components/Shares/Card/SubCard";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import { SearchFaculty, SearchWord } from "../../../components/Shares/Search";
import useApi from "../../../hooks/useGetApi";
import Loading from "../../../components/Shares/Loading";
import PaginationCustom from "../../../components/Shares/Pagination";
import { LIMIT } from "../../../const/const";
import ActionRedirect from "../../../components/Shares/ActionRedirect";

const TopicRegis = () => {
  const searchRef = useRef("");
  const facultyRef = useRef(0);
  const [pagi, setPagi] = useState(1);
  const [url, setUrl] = useState(`api/topic-registation?limit=${LIMIT}&offset=${pagi - 1}&status=2`);
  const { data, isLoading } = useApi(url);

  var listTopic = [];
  if (data && data.total !== 0) {
    listTopic = data.topicRegistrations;
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    setUrl(`/api/topic-registation?search=${searchRef.current.value}&facultyId=${facultyRef.current.value}&limit=${LIMIT}&offset=${pagi - 1}&status=2`)
  };

  return (
    <div className="topic-regis">
      <Card title="Danh sách đề tài đề xuất">
        <SubCard title="Tìm kiếm">
          <Form className="search" onSubmit={handleSubmit}>
            <SearchWord searchRef={searchRef}></SearchWord>
            <SearchFaculty facultyRef={facultyRef}></SearchFaculty>
            <Form.Group>
              <Button variant="primary" type="submit" className="search-submit">
                {isLoading ? <Loading></Loading> : <></>}
                Tìm kiếm
              </Button>
            </Form.Group>
          </Form>
        </SubCard>
        <SubCard title="Danh sách">
          {listTopic.length === 0 ? (
            <EmptyListNoti title={"Không có đề tài nào!"} />
          ) : (
            <div>
              <Table bordered hover size="sm">
                <thead>
                  <tr>
                    <th>Mã đề xuất</th>
                    <th>Người hướng dẫn</th>
                    <th>Số điện thoại - Email</th>
                    <th>Tên đề tài đề xuất</th>
                    <th>Thao tác</th>
                  </tr>
                </thead>
                <tbody>
                  {listTopic.map((item, index) => {
                    return (
                      <tr key={index}>
                        <td>{item.id}</td>
                        <td>{item.lecture}</td>
                        <td>
                          {item.phone}
                          <br />
                          {item.email}
                        </td>
                        <td>{item.name}</td>
                        <td>
                          <ActionRedirect
                            todo={[
                              { name: "Đăng ký", href: "/login" },
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
                limit={LIMIT}
              />
            </div>
          )}
        </SubCard>
      </Card>
    </div>
  );
};

export default TopicRegis;
