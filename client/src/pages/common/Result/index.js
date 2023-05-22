import Accordion from "react-bootstrap/Accordion";
import Button from "react-bootstrap/Button";
import Badge from "react-bootstrap/Badge";
import { useEffect, useRef, useState } from "react";
import Form from "react-bootstrap/Form";
import { useDispatch } from "react-redux";

import Card from "../../../components/Shares/Card";
import "./style.css";
import { data, total } from "./data";
import PaginationCustom from "../../../components/Shares/Pagination";
import { LIMIT } from "../../../const/const";
import SubCard from "../../../components/Shares/Card/SubCard";
import {
  SearchDepartment,
  SearchFaculty,
  SearchWord,
} from "../../../components/Shares/Search";
import Loading from "../../../components/Shares/Loading";
import { SortByMark } from "../../../utils/sort";
import { Link } from "react-router-dom";

const Result = () => {
  const [faculty, setFaculty] = useState(0);
  const [pagi, setPagi] = useState(1);
  const dispatch = useDispatch();
  const searchRef = useRef("");
  const departmentRef = useRef(0);
  useEffect(() => {
    // eslint-disable-next-line
  }, [dispatch, pagi]);

  const handleSubmitForm = (e) => {
    e.preventDefault();
  };
  let isLoading = false;
  const sortData = SortByMark(data);
  return (
    <Card title={"Đánh giá kết quả thực hiện đề tài"}>
      <SubCard title={"Tìm kiếm"}>
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
          <Form.Group>
            <Button variant="primary" type="submit" className="search-submit">
              {isLoading ? <Loading></Loading> : <></>}
              Tìm kiếm
            </Button>
          </Form.Group>
        </Form>
      </SubCard>
      <SubCard title={"Danh sách đề tài"}>
        <Accordion>
          {sortData.map((item, index) => {
            return (
              <Accordion.Item eventKey={index} key={index}>
                <Accordion.Header >
                  <span style={{ fontWeight: "bold", width: "85%" }}>
                    {item.topic_name}
                  </span>

                  <Badge
                    bg="danger"
                    style={{ margin: "0 0 0 24px", width: "50px" }}
                  >
                    {item.mark + "đ"}
                  </Badge>
                </Accordion.Header>
                <Accordion.Body>
                  {item.topic_mark.map((mark, index) => {
                    return (
                      <div className="row" style={{margin: "0 0 12px 0"}}>
                        <div className="row topic-mark" key={index}>
                          <div className="col-4">
                            <span>Giảng viên chấm: </span>
                            {mark.lecture}
                          </div>
                          <div className="col-3">
                            <Link to={"#"}>Biên bản nghiệm thu</Link>
                          </div>
                          <div className="col-2">
                            <span>Điểm:</span>
                            <Badge bg="danger" style={{ margin: "0 0 0 24px" }}>
                              {mark.point + "đ"}
                            </Badge>
                          </div>
                        </div>
                        <div className="row">
                          <div className="col-10">Nhận xét: {mark.comment}</div>
                        </div>
                      </div>
                    );
                  })}
                </Accordion.Body>
              </Accordion.Item>
            );
          })}
        </Accordion>
      </SubCard>

      <PaginationCustom
        setPagi={setPagi}
        currentPage={pagi}
        total={total}
        limit={LIMIT}
        style={{ marginTop: "24px" }}
      />
    </Card>
  );
};

export default Result;
