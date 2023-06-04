import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

import { FacultyDoneTopicSearch } from "../../../../components/Shares/Search";
import { FacultyDoneTopic, userSelector } from "../../../../store/selectors";
import { Badge, Table } from "react-bootstrap";
import Action from "../../../../components/Shares/Action";
import Detail from "../../../../components/Shares/Action/Detail";
import { getStatus } from "../../../../utils/getStatus";
import EmptyListNoti from "../../../../components/Shares/EmptyListNoti";
import { fetchFacultyDoneTopics } from "../FacultyTopicSlice";

const DoneTopic = () => {
    const dispatch = useDispatch()
    const doneTopic = useSelector(FacultyDoneTopic)
    const facultyID = useSelector(userSelector).faculty_id
    useEffect(() => {
        dispatch(fetchFacultyDoneTopics({ facultyID: facultyID }))
        // eslint-disable-next-line
    }, [dispatch])

    return (
        <>
            <FacultyDoneTopicSearch />
            {doneTopic.topics ?
                <>
                    <Table >
                        <thead>
                            <tr>
                                <th style={{ width: "60px" }}>STT</th>
                                <th style={{ width: "150px" }}>Giảng Viên</th>
                                <th >Đề Tài</th>
                                <th style={{ width: "60px" }}>SL Sinh Viên</th>
                                <th style={{ width: "60px" }}>Trạng Thái</th>
                                <th style={{ width: "60px" }}>Thao Tác</th>
                            </tr>
                        </thead>
                        <tbody>
                            {doneTopic.topics.map((item, index) => {
                                return (
                                    <tr key={item.id}>
                                        <td style={{ textAlign: "center" }}>{index + 1}</td>
                                        <td >{item.lectureInfo.degree + "." + item.lectureInfo.name}</td>
                                        <td >{item.name}</td>
                                        <td style={{ textAlign: "center" }}>
                                            <Badge>
                                                {item.estimateStudent}
                                            </Badge>
                                        </td>
                                        <td style={{ textAlign: "center" }}>{getStatus(item.status)}</td>
                                        <td style={{ textAlign: "center" }}>
                                            <Action todo={[
                                                <Detail name="Chi tiết" topicIn={item} />,
                                            ]} />
                                        </td>
                                    </tr>
                                )
                            })}
                        </tbody>
                    </Table>
                </>

                :
                <EmptyListNoti title={"Danh sách đề tài trống"} />
            }
        </>
    )
}

export default DoneTopic;