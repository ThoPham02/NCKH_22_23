import { useEffect, useRef } from "react";
import Card from "../../../components/Shares/Card";
import { TopicSearch } from "../../../components/Shares/Search";
import { useDispatch, useSelector } from "react-redux";
import { Table } from "react-bootstrap";
import { LectureTopicSelector } from "../../../store/selectors";

const Topic = () => {
  const dispatch = useDispatch();
  const searchRef = useRef("");
  const statusRef = useRef(0);
  const dateFromRef = useRef("");
  const dateToRef = useRef("");
  const departmentRef = useRef(0);

  const lectureTopic = useSelector(LectureTopicSelector)
	// useEffect(() => {
	// 	if (!lectureTopic.topics) {
	// 		dispatch()
	// 	}
	// 	// eslint-disable-next-line
	// }, [dispatch])

	// useEffect(() => {
	// 	if (lectureTopic) {
			
	// 	}
	// 	// eslint-disable-next-line
	// }, []);

  const handleSubmitForm = (e) => {
    e.preventDefault();
  }

  return (
    <>
      <Card title={"Danh sách đề tài"}>
        <TopicSearch
          handleSubmitForm={handleSubmitForm}
          searchRef={searchRef}
          statusRef={statusRef}
          dateFromRef={dateFromRef}
          dateToRef={dateToRef}
          departmentRef={departmentRef}
        />
        <Table bordered hover size="sm" className="topic-table">

        </Table>
      </Card>
    </>
  )
};

export default Topic;
