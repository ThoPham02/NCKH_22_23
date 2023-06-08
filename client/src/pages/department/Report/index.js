import { Table } from "react-bootstrap";
import Card from "../../../components/Shares/Card";

const Report = () => {

    return (
        <Card title={"Danh sách đề tài nghiệm thu cấp bộ môn"}>
            <Table striped hover size="sm">
                <thead>
                    <tr>
                        <th>STT</th>
                        <th>Đề tài</th>
                        <th>Mô tả báo cáo</th>
                        <th>Link đính kèm</th>
                        <th>Thao tác</th>
                    </tr>
                </thead>
            </Table>
        </Card>
    )
}

export default Report;