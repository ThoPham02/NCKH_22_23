import { useDispatch, useSelector } from "react-redux";
import { Accordion, Table } from "react-bootstrap";
import { useEffect } from "react";

import "./style.css"
import Card from "../../../components/Shares/Card";
import { fetchSubcommittee } from "./facultySubcommitteeSlice";
import { FacultySubcommittee } from "../../../store/selectors";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import AddSubcommittee from "../../../components/Shares/Action/AddSubcommittee";
import SubCard from "../../../components/Shares/Card/SubCard";
import { GroupRole } from "../../../const/const";

const Report = () => {
    const dispatch = useDispatch()
    useEffect(() => {
        dispatch(fetchSubcommittee())
    }, [dispatch])

    const report = useSelector(FacultySubcommittee)
    return (
        <div>
            <Card title={"Danh sách tiểu ban"}>
                <AddSubcommittee />
                <div style={{ marginTop: "12px" }}>
                    {report.subcommittee ?
                        <Accordion>
                            {report.subcommittee.map((item) => {
                                return (
                                    <Accordion.Item eventKey="0" key={item.id} className="add-">
                                        <Accordion.Header>
                                            <span className="cus-accordion-header">{item.name}</span>
                                        </Accordion.Header>
                                        <Accordion.Body>
                                            <SubCard title="Hội đồng nghiệm thu:">
                                                <Table>
                                                    <thead>
                                                        <tr>
                                                            <th style={{ width: "60px" }}>STT</th>
                                                            <th style={{ width: "150px" }}>Giảng Viên</th>
                                                            <th >Liên Hệ</th>
                                                            <th style={{ width: "60px" }}>Vai trò</th>
                                                        </tr>
                                                    </thead>
                                                </Table>
                                                <tbody >
                                                    {item.groups.map((group, index) => {
                                                        return (
                                                            <tr>
                                                                <td >{index + 1}</td>
                                                                <td >{group.lecture.name}</td>
                                                                <td style={{textAlign:"center"}}>{group.lecture.email} <br /> {group.lecture.phone}</td>
                                                                <td >{GroupRole[group.role-1]}</td>
                                                            </tr>
                                                        )
                                                    })}

                                                </tbody>
                                            </SubCard>
                                        </Accordion.Body>
                                    </Accordion.Item>
                                )
                            })}
                        </Accordion>
                        :
                        <EmptyListNoti title={"Chưa có tiểu ban nào"} />}
                </div>
            </Card>
            <Card title={"Danh sách đề tài"}>

            </Card>
        </div>
    )
}

export default Report;