import { useSelector } from "react-redux";
import Card from "../../../../components/Shares/Card";
import { SchoolReportSearch } from "../../../../components/Shares/Search";
import { userSelector } from "../../../../store/selectors";
import { FacultySubcommittee } from "../../../../store/selectors";
import EmptyListNoti from "../../../../components/Shares/EmptyListNoti";
import { Table } from "react-bootstrap";
import Action from "../../../../components/Shares/Action";
import Detail from "../../../../components/Shares/Action/Detail";

const SchoolReport = () => {
    const user = useSelector(userSelector)
    const topic = useSelector(FacultySubcommittee)

    return (
        <Card title={"Danh sách đề tài báo cáo cấp trường"}>
            <SchoolReportSearch facultyID={user.faculty_id} />
            {topic.school && topic.school.length > 0 ?
                <Table striped hover size="sm" >
                    <thead>
                        <tr>
                            <th>STT</th>
                            <th>Đề tài</th>
                            <th>Tiểu ban</th>
                            <th>Thao tác</th>
                        </tr>
                    </thead>
                    <tbody>
                        {topic.school.map((item, index) => {
                            return (
                                <tr key={item.id}>
                                    <td style={{ textAlign: "center" }}>{index + 1}</td>
                                    <td>{item.name}</td>
                                    <td style={{ textAlign: "center" }}>{item.subcommitteeID.name ? item.subcommitteeID.name : "Chưa có"}</td>
                                    <td>
                                        <Action todo={[<Detail name={"Xem chi tiết"} topicIn={item} />]} />
                                    </td>
                                </tr>
                            )
                        })}
                    </tbody>
                </Table> :
                <EmptyListNoti title={"Danh sách đề tài trống"} />}
        </Card>
    )
}

export default SchoolReport;