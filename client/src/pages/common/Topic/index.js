import { Table } from "react-bootstrap";
import { useDispatch, useSelector } from "react-redux";
import { useEffect, useRef, useState } from "react";

import "./style.css"
import Card from "../../../components/Shares/Card";
import { TopicSearch } from "../../../components/Shares/Search";
import { fetchTopics } from "./CommonTopicSlice";
import { CommonTopicSelector } from "../../../store/selectors";
import Action from "../../../components/Shares/Action";
import { getStatus } from "../../../utils/getStatus";
import PaginationCustom from "../../../components/Shares/Pagination";
import { LIMIT } from "../../../const/const";
import Detail from "../../../components/Shares/Action/Detail";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";

const Topic = () => {
    const dispatch = useDispatch();
    const searchRef = useRef("");
    const statusRef = useRef(0);
    const dateFromRef = useRef("");
    const dateToRef = useRef("");
    const departmentRef = useRef(0);
    const eventRef = useRef(0);

    const topics = useSelector(CommonTopicSelector)
    const [list, setList] = useState([])
    const [pagi, setPagi] = useState(1)
    useEffect(() => {
        if (!topics.topics) {
            dispatch(fetchTopics({
                limit: LIMIT,
                offset: (pagi - 1) * LIMIT,
            }))
        }
        // eslint-disable-next-line
    }, [dispatch, topics.topics, pagi])
    useEffect(() => {
        if (topics.topics) {
            setList(topics.topics)
        }
        // eslint-disable-next-line
    }, [topics.topics])
    const [faculty, setFaculty] = useState(0)
    const handleSubmitForm = (e) => {
        e.preventDefault();
        setPagi(1)
        dispatch(fetchTopics({
            status: statusRef.current.value,
            search: searchRef.current.value,
            timeStart: dateFromRef.current.value,
            timeEnd: dateToRef.current.value,
            event: eventRef.current.value,
            departmentID: departmentRef.current.value,
            facultyID: faculty,
            limit: LIMIT,
            offset: (pagi - 1) * LIMIT
        }))
        setList(topics.topics)
    }

    return (
        <>
            <Card title={"Danh sách đề tài"}>
                <TopicSearch
                    handleSubmitForm={handleSubmitForm}
                    searchRef={searchRef}
                    statusRef={statusRef}
                    dateFromRef={dateFromRef}
                    faculty={faculty}
                    setFaculty={setFaculty}
                    dateToRef={dateToRef}
                    departmentRef={departmentRef}
                    eventRef={eventRef}
                />
                {list.length === 0 ?
                    <EmptyListNoti title={"Không có đề tài nào "} /> :
                    <>
                        <Table bordered hover size="sm" className="topic-table">
                            <thead>
                                <tr>
                                    <th style={{ width: "60px" }}>STT</th>
                                    <th style={{ width: "150px" }}>Giảng viên</th>
                                    <th style={{ width: "260px" }}>Liên hệ</th>
                                    <th>Đề tài</th>
                                    <th style={{ width: "120px" }}>Trạng thái</th>
                                    <th style={{ width: "60px" }}>Thao tác</th>
                                </tr>
                            </thead>
                            <tbody>
                                {list.map((item, index) => {
                                    return (
                                        <tr key={item.id}>
                                            <td style={{ textAlign: "center" }}>{index + 1}</td>
                                            <td>{item.lectureInfo.degree + "." + item.lectureInfo.name}</td>
                                            <td style={{ textAlign: "center" }}>{item.lectureInfo.email}<br />{item.lectureInfo.phone}</td>
                                            <td>{item.name}</td>
                                            <td style={{ textAlign: "center" }}>{getStatus(item.status)}</td>
                                            <td><Action todo={[
                                                < Detail name="Xem chi tiết" topicIn={item} />,
                                            ]} /></td>
                                        </tr>
                                    )
                                })}
                            </tbody>
                        </Table>
                        <PaginationCustom setPagi={setPagi} pagi={pagi} total={topics.total} limit={LIMIT} />
                    </>
                }
            </Card>
        </>
    )
};

export default Topic;
