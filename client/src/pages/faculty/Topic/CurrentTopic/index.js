import { Badge, Table } from "react-bootstrap";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";

import { FacultyCurrentTopicSearch } from "../../../../components/Shares/Search";
import { FacultyCurrentTopic, userSelector } from "../../../../store/selectors";
import EmptyListNoti from "../../../../components/Shares/EmptyListNoti";
import { getStatus } from "../../../../utils/getStatus";
import Action from "../../../../components/Shares/Action";
import Detail from "../../../../components/Shares/Action/Detail";
import { fetchFacultyCurentTopics, updateStatus } from "../FacultyTopicSlice";
import Confirm from "../../../../components/Shares/Confirm";

const CurrentTopic = () => {
    const dispatch = useDispatch()
    const currentTopic = useSelector(FacultyCurrentTopic)
    const facultyID = useSelector(userSelector).facultyID
    useEffect(() => {
        dispatch(fetchFacultyCurentTopics({ facultyID: facultyID }))
        // eslint-disable-next-line
    }, [dispatch])
    return (
        <div>
            <FacultyCurrentTopicSearch />
            {currentTopic.topics ?
                <>
                    <Table striped hover size="sm" >
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
                            {currentTopic.topics.map((item, index) => {
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
                                                item.status === 4 ?
                                                    <Confirm
                                                        title={"Duyệt đề tài"}
                                                        isAction={true}
                                                        action={updateStatus({ id: item.id, status: item.status * 2, facultyID: facultyID })}
                                                        content={"Xác nhận duyệt đề tài"}
                                                    /> : <></>,
                                                item.status === 4 ?
                                                    <Confirm
                                                        title={"Hủy đề tài"}
                                                        isAction={true}
                                                        action={updateStatus({ id: item.id, status: 1, facultyID: facultyID })}
                                                        content={"Xác nhận hủy đề tài"}
                                                    />
                                                    : <></>
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
        </div>
    )
}

export default CurrentTopic;