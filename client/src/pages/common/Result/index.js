import { useState } from "react";
import SwitchCard from "../../../components/Shares/Card/SwitchCard";
import { ResultSearch } from "../../../components/Shares/Search";
import { useSelector } from "react-redux";
import { ResultTopicSelector } from "../../../store/selectors";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import { Badge, Table } from "react-bootstrap";
import Action from "../../../components/Shares/Action";
import Detail from "../../../components/Shares/Action/Detail";
import PaginationCustom from "../../../components/Shares/Pagination";
import { LIMIT } from "../../../const/const";

const Result = () => {
  const [switchPage, setSwitchPage] = useState(false)
  const [pagi, setPagi] = useState(1)
  const [pagi2, setPagi2] = useState(1)
  const topics = useSelector(ResultTopicSelector)

  return (
    <SwitchCard setSwitchPage={setSwitchPage} switchPage={switchPage}>
      {switchPage
        ?
        <>
          <ResultSearch facultyID={0} isCurrent={false} limit={20} offset={20 * (pagi - 1)} />
          {topics.current && topics.done.length !== 0 ?
            <>
              <Table striped hover size="sm">
                <thead>
                  <tr>
                    <th style={{ width: "60px" }}>STT</th>
                    <th>Đề tài</th>
                    <th style={{ width: "200px" }}>Tiểu ban</th>
                    <th style={{ width: "100px" }}>Điểm số</th>
                    <th style={{ width: "100px" }}>Thao tác</th>
                  </tr>
                </thead>
                <tbody>
                  {topics.done.map((item, index) => {
                    return (
                      <tr key={item.id}>
                        <td style={{ textAlign: "center" }}>{index + 1 + LIMIT * (pagi2 - 1)}</td>
                        <td>{item.name}</td>
                        <td style={{ textAlign: "center" }}>{item.subcommitteeID.name ? item.subcommitteeID.name : "Chưa có"}</td>
                        <td style={{ textAlign: "center" }}><Badge bg="danger">{item.mark + "đ"}</Badge></td>
                        <td><Action todo={[
                          <Detail name={"Xem chi tiết"} topicIn={item} />
                        ]} /></td>
                      </tr>
                    )
                  })}
                </tbody>
              </Table>
              <PaginationCustom pagi={pagi} setPagi={setPagi} limit={20} total={topics.total} />
            </>
            :
            <EmptyListNoti title={"Danh sách đề tài trống"} />
          }
        </>
        :
        <>
          <ResultSearch isCurrent={true} limit={20} offset={20 * (pagi2 - 1)} />
          {topics.current && topics.current.length !== 0 ?
            <>
              <Table striped hover size="sm">
                <thead>
                  <tr>
                    <th style={{ width: "60px" }}>STT</th>
                    <th>Đề tài</th>
                    <th style={{ width: "200px" }}>Tiểu ban</th>
                    <th style={{ width: "100px" }}>Điểm số</th>
                    <th style={{ width: "100px" }}>Thao tác</th>
                  </tr>
                </thead>
                <tbody>
                  {topics.current.map((item, index) => {
                    return (
                      <tr key={item.id}>
                        <td style={{ textAlign: "center" }}>{index + 1 + LIMIT * (pagi2 - 1)}</td>
                        <td>{item.name}</td>
                        <td style={{ textAlign: "center" }}>{item.subcommitteeID.name ? item.subcommitteeID.name : "Chưa có"}</td>
                        <td style={{ textAlign: "center" }}><Badge bg="danger">{item.mark + "đ"}</Badge></td>
                        <td><Action todo={[
                          <Detail name={"Xem chi tiết"} topicIn={item} />
                        ]} /></td>
                      </tr>
                    )
                  })}
                </tbody>
              </Table>
              <PaginationCustom pagi={pagi2} setPagi={setPagi2} limit={20} total={topics.total} />
            </>
            :
            <EmptyListNoti title={"Danh sách đề tài trống"} />
          }
        </>
      }
    </SwitchCard>
  )
}

export default Result;