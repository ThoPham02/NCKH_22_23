import Card from "../../../components/Shares/Card";
import SubCard from "../../../components/Shares/Card/SubCard";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";

const Mark = () => {
    return (
        <div>
            <Card title={"Danh sách đề tài được phân công"}>
                <SubCard title={"Tiểu ban: Công Nghệ Thông Tin 1"}></SubCard>
                <EmptyListNoti title={"Danh sách đề tài trống"}/>
            </Card>
        </div>
    )
}

export default Mark;