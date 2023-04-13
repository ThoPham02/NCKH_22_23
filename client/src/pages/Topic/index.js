import "./topic.css"
import { Search } from "../../components/Shares/Search";
import TableTopic from "../../components/Shares/Table";
import Pagination1 from "../../components/Shares/Pagination";
import { useDispatch, useSelector } from "react-redux";
import { topicSelector } from "../../store/selectors";
import { useEffect, useState } from "react";
import { fetchTopic } from "./TopicSlice";

function Topic() {
  const dispatch = useDispatch();
  let topic = useSelector(topicSelector)

  const list = topic.topicList
  const [listTopic, setListTopic] = useState(list);
  const [pagi, setPagi] = useState(1)

  useEffect(() => {
    dispatch(fetchTopic())
  }, [dispatch]);

  useEffect(() => {
    setListTopic(list)
  }, [list]);

  const listHead = ["id", "name", "description", "description_url", "lecture_id", "faculity_id", "created_at"];
  const listKey = ["id", "name", "description", "description_url", "lecture_id", "faculity_id", "created_at"];
  
  return (
    <div className="topic">
      <Search />
      <TableTopic listHead={listHead} listKey={listKey} listItem={listTopic} />
      <Pagination1 limit={topic.limit} pagi={pagi} setPagi={setPagi} currentPage={topic.currentPage} total={topic.total}/>
    </div>
  );
}

export default Topic;
