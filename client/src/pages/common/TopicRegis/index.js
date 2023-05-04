import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { useState } from "react";

import "./style.css";
import Card from "../../../components/Shares/Card";
import SubCard from "../../../components/Shares/Card/SubCard";
import { SearchFaculty, SearchWord } from "../../../components/Shares/Search";
import useApi from "../../../hooks/useGetApi";
import Loading from "../../../components/Shares/Loading";

const TopicRegis = () => {
  const [filter, setFilter] = useState({
    facultyId: 0,
    word: "",
    limit: 20,
    offset: 0,
  });
  let url = `/api/topic-registation?search=${filter.word}&facultyId=${filter.facultyId}&limit=${filter.limit}&offset=${filter.offset}`;
  const { data, isLoading, error, fetchData } = useApi(url);

  const handleSubmit = (e) => {
    e.preventDefault();
    fetchData();
  };

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

        <SubCard title="Danh sách"></SubCard>
      </Card>
    </div>
  );
};

export default TopicRegis;
