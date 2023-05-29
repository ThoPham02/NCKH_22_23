import { useSelector } from "react-redux";
import { Button } from "react-bootstrap";
import { useRef, useState } from "react";

import "./style.css"
import CustomeProgress from "./CustomeProgress";
import { convertTimestampToDateString } from "../../../utils/time";
import Confirm from "../Confirm";
import { cancelEvent, updateStage } from "../../../pages/admin/DashBoard/EventSlice";
import { userSelector } from "../../../store/selectors";
import { AdminEventSelector } from "../../../store/selectors";
import { convertDateToTimestamp } from "../../../utils/time";

const TimeLine = (props) => {
    const { data } = props

    const [description, setDescription] = useState("")
    const [timeStart, setTimeStart] = useState("")
    const [timeEnd, setTimeEnd] = useState("")

    let current = new Date().getTime();
    let now = 0
    for (let i = 0; i < data.stages.length; i++) {
        if (data.stages[i].timeStart !== 0 && data.stages[i].timeStart <= current) {
            now = i
        }
    }
    const [stage, setStage] = useState(data.stages[now])
    const [edit, setEdit] = useState(false)
    let isCurrent = useSelector(userSelector).role === 5 && data.isCurrent

    let isLoading = useSelector(AdminEventSelector).status === 'loading'

    return (
        <div>
            <div className="row">
                <div className="col-2 detail-head">NCKH:</div>
                <div className="col-6">{data.name}</div>
                {
                    isCurrent ?
                        <div className="col-4">
                            <Confirm
                                title={"Kết thúc NCKH"}
                                content={"Xác nhận hủy  NCKH hiện tại!"}
                                isLoading={isLoading}
                                action={cancelEvent({ id: data.id })}
                            />
                        </div> : <></>
                }
            </div>
            <div className="row">
                <div className="col-2 detail-head">Năm học: </div>
                <div className="col-6">{data.schoolYear}</div>
            </div>
            <CustomeProgress setStage={setStage} data={data.stages} />
            <div className="stage-detail">
                <div style={{ fontWeight: "bold" }}>Thông tin giai đoạn:</div>
                <div className="row detail-item">
                    <div className="col-2 detail-head">Giai đoạn:</div>
                    <div className="col-8">{stage.name}</div>
                </div>
                <div className="row detail-item">
                    <div className="col-2 detail-head">Mô tả:</div>
                    <div className="col-8">
                        {edit ?
                            <input
                                type="text"
                                placeholder="Mô tả chung về giai đoạn"
                                style={{ width: "600px" }}
                                className="input"
                                value={description}
                                onChange={(e) => setDescription(e.target.value)}
                            /> : <>{stage.description}</>}
                    </div>
                </div>
                <div className="row detail-item">
                    <div className="col-2 detail-head">Thời gian:</div>
                    <div className="col-8">
                        {"Từ "}
                        {edit ?
                            <input
                                type="date"
                                value={timeStart}
                                onChange={(e) => setTimeStart(e.target.value)}
                            /> : <>{stage.timeStart === 0 ? " ... " : convertTimestampToDateString(stage.timeStart)}</>}
                        {" đến "}
                        {edit ?
                            <input
                                type="date"
                                value={timeEnd}
                                onChange={(e) => setTimeEnd(e.target.value)}
                            /> : <>{stage.timeEnd === 0 ? " ... " : convertTimestampToDateString(stage.timeEnd)}</>}
                    </div>
                </div>
                <div className="row detail-item">
                    <div className="col-2 detail-head"></div>
                    <div className="col-8" style={{ fontWeight: "bold" }}>
                        {
                            // eslint-disable-next-line
                            edit ? <input type="file" /> : <a href="#">Tài liệu đính kèm</a>
                        }
                    </div>
                </div>
                {isCurrent
                    ?
                    <div className="row detail-item">
                        <div className="col-2 detail-head"></div>
                        <div className="col-8" style={{ fontWeight: "bold", marginTop: "12px" }}>
                            {edit ?
                                <div>
                                    <Button style={{ marginRight: "12px" }} onClick={() => setEdit(false)}>Hủy</Button>
                                    <Confirm
                                        title={"Xác nhận"}
                                        content={"Xác nhận thao tác"}
                                        isLoading={isLoading}
                                        action={updateStage({
                                            stageID: stage.id,
                                            description: description,
                                            timeStart: convertDateToTimestamp(timeStart),
                                            timeEnd: convertDateToTimestamp(timeEnd)
                                        })}
                                        onClick={() => setEdit(false)}
                                    />
                                </div>
                                : <Button onClick={() => setEdit(true)}>Chỉnh sửa</Button>}
                        </div>
                    </div>
                    : <></>}
            </div>
        </div>
    )
}


export default TimeLine;