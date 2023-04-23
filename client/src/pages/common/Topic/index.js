import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";

import "./style.css";
import Card from "../../../components/Shares/Card";
import SubCard from "../../../components/Shares/Card/SubCard";
import SearchWord from "../../../components/Shares/Search/Word";
import {
  SearchDateFrom,
  SearchDateTo,
  SearchDepartment,
  SearchFaculty,
  SearchStatus,
} from "../../../components/Shares/Search";
import { useState } from "react";

const Topic = () => {
  const [filter, setFilter] = useState({
    departmentId: 0,
    facultyId: 0,
    word: "",
    status: 0,
    dateTo: "",
    dateFrom: "",
  });

  return (
    <div className="topic">
      <Card title="Danh sách đề tài">
        <SubCard title="Tìm kiếm">
          <Form className="search">
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
            <Form.Group>
              <Button variant="primary" type="submit" className="search-submit">
                Tìm kiếm
              </Button>
            </Form.Group>
          </Form>
        </SubCard>
      </Card>
    </div>
  );
};

export default Topic;
