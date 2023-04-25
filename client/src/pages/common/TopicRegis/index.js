import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { useState } from "react";

import "./style.css";
import Card from "../../../components/Shares/Card";
import SubCard from "../../../components/Shares/Card/SubCard";
import SearchWord from "../../../components/Shares/Search/Word";
import {
  SearchFaculty,
} from "../../../components/Shares/Search";
import { listDataTopicRegis } from "./data";
import TableTopic from "../../../components/Shares/Table";

const TopicRegis = () => {
  const [filter, setFilter] = useState({
    faculty: "",
    word: "",
    limit: 20,
    offset: 0,
  });


  const handleSubmit = () => { };

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
                Tìm kiếm
              </Button>
            </Form.Group>
          </Form>
        </SubCard>

        <SubCard title="Danh sách">
          <TableTopic
            listHead={['STT', 'Khoa','Người hướng dẫn', 'Học hàm/Học vị', 'Số điện thoại liên hệ','Email', 'Tên đề tài', 'Thao tác']}
            listItem={listDataTopicRegis}
            listKey={['faculty', 'lecture', 'rank','phone','email','name']}
          />
        </SubCard>
      </Card>
    </div>
  );
};

export default TopicRegis;
