import "./topic.css"
import { Search } from "../../components/Shares/Search";
import TableTopic from "../../components/Shares/Table";

function Topic() {
  const listHead = ["First Name", "Last Name", "UserName"];
  const listKey = ["first", "last", "name"];
  const listItem = [
    {
      first: "Mark",
      last: "Otto",
      name: "@mdo",
    },
    {
      first: "Mark",
      last: "Otto",
      name: "@mdo",
    },
    {
      first: "Mark",
      last: "Otto",
      name: "@mdo",
    },
  ];

  return (
    <div className="topic">
      <Search />
      <TableTopic listHead={listHead} listKey={listKey} listItem={listItem} />
    </div>
  );
}

export default Topic;
