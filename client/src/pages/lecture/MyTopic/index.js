import { useEffect } from "react";
import Card from "../../../components/Shares/Card";
import "./style.css"
import { useDispatch, useSelector } from "react-redux";
import { fetchTopics } from "../../common/Topic/CommonTopicSlice";
import { userSelector } from "../../../store/selectors";
import { CommonTopicSelector } from "../../../store/selectors";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import Suggest from "../../../components/Shares/Action/Suggest";
import { Table } from "react-bootstrap";
import { getStatus } from "../../../utils/getStatus";
import Action from "../../../components/Shares/Action";
import Detail from "../../../components/Shares/Action/Detail";

const MyTopic = () => {
    const dispatch = useDispatch()
    const user = useSelector(userSelector)
    useEffect(() => {
        dispatch(fetchTopics({ userID: user.id }))
    }, [dispatch, user.id])

    const topicSelector = useSelector(CommonTopicSelector)

    return (
        <div className="lecture_topic">
            <Card title={"Đề tài của tôi"}>
                <Suggest />
                {topicSelector.topics ?
                    <div>
                        <Table>
                            <thead>
                                <tr>
                                    <th style={{ width: "60px" }}>STT</th>
                                    <th >Đề Tài</th>
                                    <th style={{ width: "200px"}}>Nhóm sinh viên</th>
                                    <th style={{ width: "60px" }}>Trạng Thái</th>
                                    <th style={{ width: "60px" }}>Thao Tác</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    topicSelector.topics.map((item, index) => {
                                        return (
                                            <tr key={index}>
                                                <td style={{ textAlign: "center" }}>{index + 1}</td>
                                                <td>{item.name}</td>
                                                <td>{item.name}</td>
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