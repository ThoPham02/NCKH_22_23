import { Accordion, Table } from "react-bootstrap";
import { useSelector } from "react-redux";

import { CurrentTopicsSearch } from "../../../components/Shares/Search";
import "./style.css"
import { CommonTopicSelector } from "../../../store/selectors";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import { getStatus } from "../../../utils/getStatus";
import Action from "../../../components/Shares/Action";
import Detail from "../../../components/Shares/Action/Detail";
import PaginationCustom from "../../../components/Shares/Pagination";
import { LIMIT } from "../../../const/const";
import { useState } from "react";

const Topic = () => {
    const commonTopics = useSelector(CommonTopicSelector)

    const [pagi, setPagi] = useState(1)
    const [pagi2, setPagi2] = useState(1)

    return (
        <div className="topic">

            <Accordion defaultActiveKey="0">
                <Accordion.Item eventKey="0">
                    <Accordion.Header >
                        <div className="accordion-header">Nghiên Cứu Khoa Học sinh viên đang diễn ra</div>
                    </Accordion.Header>
                    <Accordion.Body>
                        <CurrentTopicsSearch offset={LIMIT * (pagi - 1)} limit={LIMIT} />

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
                                                        <td style={{ textAlign: "center" }}>{index + 1}</td>
                                                        <td>{item.lectureInfo.degree + "." + item.lectureInfo.name}</td>
                                                        <td style={{ textAlign: "center" }}>{item.lectureInfo.email}<br />{item.lectureInfo.phone}</td>
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
                                <PaginationCustom setPagi={setPagi} pagi={pagi} total={commonTopics.total} limit={LIMIT} />
                            </>
                            : <EmptyListNoti title={"Không có đề tài nào"} />}
                    </Accordion.Body>
                </Accordion.Item>
                <Accordion.Item eventKey="1">
                    <Accordion.Header>
                    <div className="accordion-header">Nghiên Cứu Khoa Học sinh viên đã kết thúc</div>
                    </Accordion.Header>
                    <Accordion.Body>
                        <CurrentTopicsSearch offset={LIMIT * (pagi2 - 1)} limit={LIMIT} />
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
                                                        <td style={{ textAlign: "center" }}>{index + 1}</td>
                                                        <td>{item.lectureInfo.degree + "." + item.lectureInfo.name}</td>
                                                        <td style={{ textAlign: "center" }}>{item.lectureInfo.email}<br />{item.lectureInfo.phone}</td>
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
                                <PaginationCustom setPagi={setPagi2} pagi={pagi2} total={commonTopics.total} limit={LIMIT} />
                            </>
                            : <EmptyListNoti title={"Không có đề tài nào"} />}
                    </Accordion.Body>
                </Accordion.Item>
            </Accordion>
        </div>
    )
};

export default Topic;
