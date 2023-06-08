import { useEffect } from "react";
import Card from "../../../components/Shares/Card";
import "./style.css"
import { useDispatch, useSelector } from "react-redux";
import { LectureMyTopicSelector, userSelector } from "../../../store/selectors";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import Suggest from "../../../components/Shares/Action/Suggest";
import { Badge, Table } from "react-bootstrap";
import { getStatus } from "../../../utils/getStatus";
import Action from "../../../components/Shares/Action";
import Detail from "../../../components/Shares/Action/Detail";
import { fetchLectureTopic } from "./LectureMyTopicSlice";
import AddReport from "../../../components/Shares/Action/AddReport";

const MyTopic = () => {
    const dispatch = useDispatch()
    const user = useSelector(userSelector)
    useEffect(() => {
        dispatch(fetchLectureTopic({ userID: user.id }))
    }, [dispatch, user.id])

    const topicSelector = useSelector(LectureMyTopicSelector)

    return (
        <div className="lecture_topic">
            <Card title={"NCKH  đang diễn ra"}>
                <Suggest />
                {topicSelector.topics ?
                    <div>
                        <Table striped hover size="sm" >
                            <thead>
                                <tr>
                                    <th style={{ width: "60px" }}>STT</th>
                                    <th >Đề Tài</th>
                                    <th style={{ width: "180px" }}>Số lượng SV dự kiến</th>
                                    <th style={{ width: "200px" }}>Trạng Thái</th>
                                    <th style={{ width: "60px" }}>Thao Tác</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    topicSelector.topics.filter(i => i.eventId.isCurrent === 1).map((item, index) => {
                                        return (
                                            <tr key={index}>
                                                <td style={{ textAlign: "center" }}>{index + 1}</td>
                                                <td>{item.name}</td>
                                                <td style={{ textAlign: "center" }}>
                                                    <Badge >{item.estimateStudent}</Badge>
                                                </td>
                                                <td style={{ textAlign: "center" }}>{getStatus(item.status)}</td>
                                                <td><Action todo={[
                                                    <Detail name="Chi tiết" topicIn={item} />,
                                                    item.status === 16 ? <AddReport id={item.id}/> : <></>
                                                ]} /></td>
                                            </tr>
                                        )
                                    })
                                }
                            </tbody>
                        </Table>
                    </div> :
                    <EmptyListNoti title={"Bạn chưa có đề tài nào!"} />}
            </Card>
            <Card title={"NCKH đã hoàn thành"}>
                {topicSelector.topics && topicSelector.topics.filter(i => i.eventId.isCurrent === 2).length !== 0?
                    <div>
                        <Table striped hover size="sm" >
                            <thead>
                                <tr>
                                    <th style={{ width: "60px" }}>STT</th>
                                    <th >Đề Tài</th>
                                    <th style={{ width: "180px" }}>Số lượng SV dự kiến</th>
                                    <th style={{ width: "200px" }}>Trạng Thái</th>
                                    <th style={{ width: "60px" }}>Thao Tác</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    topicSelector.topics.filter(i => i.eventId.isCurrent === 2).map((item, index) => {
                                        return (
                                            <tr key={index}>
                                                <td style={{ textAlign: "center" }}>{index + 1}</td>
                                                <td>{item.name}</td>
                                                <td style={{ textAlign: "center" }}>
                                                    <Badge >{item.estimateStudent}</Badge>
                                                </td>
                                                <td style={{ textAlign: "center" }}>{getStatus(item.status)}</td>
                                                <td><Action todo={[
                                                    <Detail name="Chi tiết" topicIn={item} />
                                                ]} /></td>
                                            </tr>
                                        )
                                    })
                                }
                            </tbody>
                        </Table>
                    </div> :
                    <EmptyListNoti title={"Bạn chưa có đề tài nào!"} />}
            </Card>
        </div>
    )
}

export default MyTopic;