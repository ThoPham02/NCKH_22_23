import { useDispatch, useSelector } from "react-redux";
import { Accordion, Table } from "react-bootstrap";
import { useEffect } from "react";

import "./style.css"
import Card from "../../../components/Shares/Card";
import { fetchSubcommittee } from "./facultySubcommitteeSlice";
import { FacultySubcommittee, userSelector } from "../../../store/selectors";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import AddSubcommittee from "../../../components/Shares/Action/AddSubcommittee";
import SubCard from "../../../components/Shares/Card/SubCard";
import { GroupRole } from "../../../const/const";
import { FacultySubcommitteeSearch } from "../../../components/Shares/Search";
import AddTopicsToSub from "../../../components/Shares/Action/AddTopicsToSub";

const Report = () => {
    const dispatch = useDispatch()
    useEffect(() => {
        dispatch(fetchSubcommittee())
    }, [dispatch])

    const report = useSelector(FacultySubcommittee)
    const user = useSelector(userSelector)

    return (
        <div>
            <Card title={"Danh sách tiểu ban"}>
                <AddSubcommittee />
                <div style={{ marginTop: "12px" }}>
                    {report.subcommittee ?
                        <Accordion>
                            {report.subcommittee.map((item, index) => {
                                return (
                                    <>
                                        <AddTopicsToSub subcommittee={item} />
                                        <Accordion.Item eventKey={index} key={item.id} className="add-">
                                            <Accordion.Header>
                                                <span className="cus-accordion-header">{item.name}</span>
                                            </Accordion.Header>
                                            <Accordion.Body>
                                                <SubCard title="Hội đồng nghiệm thu:">
                                                    <Table>
                                                        <thead>
                                                            <tr>
                                                                <th style={{ width: "60px" }}>STT</th>
                                                                <th style={{ width: "200px" }}>Giảng Viên</th>
                                                                <th >Liên Hệ</th>
                                                                <th style={{ width: "200px" }}>Vai trò</th>
                                                            </tr>
                                                        </thead>
                                                        <tbody >
                                                            {item.groups.map((group, index) => {
                                                                return (
                                                                    <tr>
                                                                        <td style={{ textAlign: "center" }}>{index + 1}</td>
                                                                        <td >{group.lecture.name}</td>
                                                                        <td style={{ textAlign: "center" }}>{group.lecture.email} <br /> {group.lecture.phone}</td>
                                                                        <td style={{ textAlign: "center" }}>{GroupRole[group.role - 1]}</td>
                                                                    </tr>
                                                                )
                                                            })}

                                                        </tbody>
                                                    </Table>
                                                </SubCard>
                                            </Accordion.Body>
                                        </Accordion.Item>
                                    </>
                                )
                            })}
                        </Accordion>
                        :
                        <EmptyListNoti title={"Chưa có tiểu ban nào"} />}
                </div>
            </Card>
            <Card title={"Danh sách đề tài"}>
                <FacultySubcommitteeSearch facultyID={user.faculty_id} />
                {report.topics ?
                    <Table >
                        <thead>
                            <tr>
                                <th>STT</th>
                                <th>Đề tài</th>
                            </tr>
                        </thead>
                        <tbody>
                            {report.topics.map((item, index) => {
                                return (
                                    <tr>
                                        <td style={{ textAlign: "center" }}>{index + 1}</td>
                                        <td>{item.name}</td>
                                    </tr>
                                )
                            })}
                        </tbody>
                    </Table>
                    : <EmptyListNoti title={"Danh sách đề tài trống!"} />}
            </Card>
        </div>
    )
}

export default Report;