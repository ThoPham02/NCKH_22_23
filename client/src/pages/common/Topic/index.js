import { Table } from "react-bootstrap";
import { useSelector } from "react-redux";
import { useState } from "react";

import { CurrentTopicsSearch, TopicSearch } from "../../../components/Shares/Search";
import "./style.css"
import { CommonCurrentTopicSelector, CommonDoneTopicSelector, userSelector } from "../../../store/selectors";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import { getStatus } from "../../../utils/getStatus";
import Action from "../../../components/Shares/Action";
import Detail from "../../../components/Shares/Action/Detail";
import PaginationCustom from "../../../components/Shares/Pagination";
import { LIMIT } from "../../../const/const";
import SwitchCard from "../../../components/Shares/Card/SwitchCard";
import Confirm from "../../../components/Shares/Confirm";
import { registationTopic } from "./CommonTopicSlice";
const Topic = () => {
    const commonTopics = useSelector(CommonCurrentTopicSelector)
    const doneTopic = useSelector(CommonDoneTopicSelector)
    const user = useSelector(userSelector)

    const [pagi, setPagi] = useState(1)
    const [pagi2, setPagi2] = useState(1)
    const [switchPage, setSwitchPage] = useState(false)

    return (
        <div className="topic">
            <SwitchCard switchPage={switchPage} setSwitchPage={setSwitchPage}>
                {!switchPage ?
                    (
                        <><CurrentTopicsSearch offset={LIMIT * (pagi - 1)} limit={LIMIT} />
                            {commonTopics.topics ?
                                <>
                                    <Table striped hover size="sm" >
                                        <thead>
                                            <tr>
                                                <th style={{ width: "60px" }}>STT</th>
                                                <th style={{ width: "150px" }}>Giảng Viên</th>
                                                <th style={{ width: "320px" }}>Liên Hệ</th>
                                                <th >Đề Tài</th>
                                                <th style={{ width: "60px" }}>Trạng Thái</th>
                                                <th style={{ width: "60px" }}>Thao Tác</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            {
                                                commonTopics.topics.map((item, index) => {
                                                    return (
                                                        <tr key={index}>
                                                            <td style={{ textAlign: "center" }}>{index + 1 + LIMIT * (pagi - 1)}</td>
                                                            <td>{item.lectureInfo.degree + "." + item.lectureInfo.name}</td>
                                                            <td style={{ textAlign: "center" }}>{item.lectureInfo.email}<br />{item.lectureInfo.phone}</td>
                                                            <td>{item.name}</td>
                                                            <td style={{ textAlign: "center" }}>{getStatus(item.status)}</td>
                                                            <td><Action todo={[
                                                                <Detail name="Chi tiết" topicIn={item} />,
                                                                item.status === 4 ?
                                                                    <Confirm
                                                                        title={"Đăng ký"}
                                                                        action={registationTopic({ topicID: item.id, studentID: user.id })}
                                                                        content={"Xác nhận đăng ký đề tài"}
                                                                        isAction={true} />
                                                                    : <></>
                                                            ]} /></td>
                                                        </tr>
                                                    )
                                                })
                                            }
                                        </tbody>
                                    </Table>
                                    <PaginationCustom setPagi={setPagi} pagi={pagi} total={commonTopics.total} limit={LIMIT} />
                                </>
                                : <EmptyListNoti title={"Không có đề tài nào"} />}
                        </>)
                    :
                    (
                        <>
                            <TopicSearch pagi={pagi2} setPagi={setPagi2} isLoading={doneTopic.status === "loading"} />
                            {doneTopic.topics ?
                                <>
                                    <Table striped hover size="sm" >
                                        <thead>
                                            <tr>
                                                <th style={{ width: "60px" }}>STT</th>
                                                <th style={{ width: "150px" }}>Giảng Viên</th>
                                                <th >Đề Tài</th>
                                                <th style={{ width: "60px" }}>Trạng Thái</th>
                                                <th style={{ width: "60px" }}>Thao Tác</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            {
                                                doneTopic.topics.map((item, index) => {
                                                    return (
                                                        <tr key={index}>
                                                            <td style={{ textAlign: "center" }}>{index + 1 + LIMIT * (pagi2 - 1)}</td>
                                                            <td>{item.lectureInfo.degree + "." + item.lectureInfo.name}</td>
                                                            <td>{item.name}</td>
                                                            <td style={{ textAlign: "center" }}>{getStatus(item.status)}</td>
                                                            <td>
                                                                <Action todo={[
                                                                    <Detail name="Chi tiết" topicIn={item} />
                                                                ]} />
                                                            </td>
                                                        </tr>
                                                    )
                                                })
                                            }
                                        </tbody>
                                    </Table>
                                    <PaginationCustom setPagi={setPagi2} pagi={pagi2} total={commonTopics.total} limit={LIMIT} />
                                </>
                                : <EmptyListNoti title={"Không có đề tài nào"} />}
                        </>
                    )
                }

            </SwitchCard>
        </div>
    )
};

export default Topic;
