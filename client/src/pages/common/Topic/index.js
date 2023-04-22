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

const Topic = () => {
  return (
    <div className="topic">
      <Card title="Danh sách đề tài">
        <SubCard title="Tìm kiếm">
          <Form className="search">
            <SearchWord></SearchWord>
            <SearchDepartment></SearchDepartment>
            <SearchFaculty></SearchFaculty>
            <SearchStatus></SearchStatus>
            <SearchDateFrom></SearchDateFrom>
            <SearchDateTo></SearchDateTo>
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
