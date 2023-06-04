import { useDispatch, useSelector } from "react-redux";
import Card from "../../../components/Shares/Card";
import "./style.css"
import { useEffect } from "react";
import { fetchStudentMyTopic } from "./StudentMyTopicSlice";
import { StudentMyTopicCurrent, StudentMyTopicDone, userSelector } from "../../../store/selectors";
import { Table } from "react-bootstrap";
import { getStatus } from "../../../utils/getStatus";
import Action from "../../../components/Shares/Action";
import Detail from "../../../components/Shares/Action/Detail";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import SubCard from "../../../components/Shares/Card/SubCard";
import TopicDetail from "../../../components/Shares/Action/Detail/TopicDetail";

const MyTopic = () => {
    const dispatch = useDispatch()
    const user = useSelector(userSelector)
    useEffect(() => {
        dispatch(fetchStudentMyTopic({ studentID: user.id }))
    }, [dispatch, user])
    const current = useSelector(StudentMyTopicCurrent)
    const done = useSelector(StudentMyTopicDone)
    return (
        <Card title={"Đề tài của tôi"}>
            {current.topics || done.topics ?
                <>
                    <SubCard title={"NCKH đang diễn ra"}>
                        {current.topics ?
                            <TopicDetail data={current.topics[0]} />
                            : 
                            <EmptyListNoti title={
                                "Bạn chưa đăng ký tham gia đề tài nào"
                            }/>}
                    </SubCard>
                    {done.topics ?
                        <SubCard title={"NCKH đã hoàn thành"}>
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
                                        done.topics.map((item, index) => {
                                            return (
                                                <tr key={index}>
                                                    <td style={{ textAlign: "center" }}>{index + 1}</td>
                                                    <td>{item.lectureInfo.degree + "." + item.lectureInfo.name}</td>
                                                    <td style={{ textAlign: "center" }}>{item.lectureInfo.email}<br />{item.lectureInfo.phone}</td>
                                                    <td>{item.name}</td>
                                                    <td style={{ textAlign: "center" }}>{getStatus(item.status)}</td>
                                                    <td><Action todo={[
                                                        <Detail name="Chi tiết" topicIn={item} />,
                                                    ]} /></td>
                                                </tr>
                                            )
                                        })
                                    }
                                </tbody>
                            </Table>
                        </SubCard>
                        : <></>
                    }
                </>
                :
                <EmptyListNoti title={"Bạn chưa tham gia nghiên cứu khoa học nào"} />
            }
        </Card >
    )
}

export default MyTopic;