import { useSelector } from "react-redux";
import { Button, FloatingLabel, Form } from "react-bootstrap";
import { useEffect, useState } from "react";
import { AiTwotoneNotification } from "react-icons/ai"

import "./style.css"
import CustomeProgress from "./CustomeProgress";
import { convertTimestampToDateString } from "../../../utils/time";
import Confirm from "../Confirm";
import { cancelEvent, updateStage } from "../../../pages/admin/DashBoard/EventSlice";
import { roleSelector, userSelector } from "../../../store/selectors";
import { AdminEventSelector } from "../../../store/selectors";
import { convertDateToTimestamp } from "../../../utils/time";
import Loading from "../Loading";

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
    useEffect(() => {
        setStage(data.stages[now])
    }, [data, now])
    const [edit, setEdit] = useState(false)

    let isEdit = useSelector(userSelector).role === 5 && data.isCurrent === 1

    let isLoading = useSelector(AdminEventSelector).status === 'loading'

    const handleResetState = () => {
        setEdit(false)
        setDescription("")
        setTimeEnd("")
        setTimeStart("")
    }

    return (
        <div>
            <div className="row">
                <div className="col-2 detail-head">NCKH:</div>
                <div className="col-6">{data.name}</div>
                {
                    isEdit ?
                        <div className="col-4">
                            <Confirm
                                title={"Kết thúc NCKH"}
                                content={"Xác nhận kết thúc NCKH hiện tại!"}
                                isLoading={isLoading}
                                action={cancelEvent({ id: data.id })}
                                variant={"danger"}
                            />
                        </div> : <></>
                }
            </div>
            <div className="row">
                <div className="col-2 detail-head">Năm học: </div>
                <div className="col-6">{data.schoolYear}</div>
            </div>
            <CustomeProgress setStage={setStage} data={data.stages} stage={stage} />
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
                            <FloatingLabel controlId="floatingTextarea2" label="Mô tả chung về giai đoạn">
                                <Form.Control
                                    as="textarea"
                                    placeholder="Mô tả chung"
                                    style={{ height: '100px' }}
                                    value={description}
                                    onChange={(e) => setDescription(e.target.value)}
                                />
                            </FloatingLabel>
                            :
                            stage.description ? stage.description : "Chưa có mô tả"}
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
                {isEdit
                    ?
                    <div className="row detail-item">
                        <div className="col-2 detail-head"></div>
                        <div className="col-8" style={{ fontWeight: "bold", marginTop: "12px" }}>
                            {edit ?
                                <div>
                                    <Button style={{ marginRight: "12px" }} onClick={() => setEdit(false)}>Hủy</Button>
                                    <Confirm
                                        title={"Xác nhận"}
                                        content={"Xác nhận chỉnh sửa giai đoạn"}
                                        isLoading={isLoading}
                                        action={updateStage({
                                            stageID: stage.id,
                                            description: description,
                                            timeStart: convertDateToTimestamp(timeStart),
                                            timeEnd: convertDateToTimestamp(timeEnd)
                                        })}
                                        onClick={handleResetState}
                                    />
                                </div>
                                : <Button onClick={() => setEdit(true)}>Chỉnh sửa</Button>}
                        </div>
                        {isLoading ? <Loading /> : <></>}
                    </div>
                    : <></>}
            </div>
        </div>
    )
}

export const SubTimeLine = (props) => {
    const { data } = props

    let current = new Date().getTime();
    let now = 0
    for (let i = 0; i < data.stages.length; i++) {
        if (data.stages[i].timeStart !== 0 && data.stages[i].timeStart <= current) {
            now = i
        }
    }
    const [stage, setStage] = useState(data.stages[now])
    useEffect(() => {
        setStage(data.stages[now])
    }, [data, now])

    const role = useSelector(roleSelector)

    const notiSuggest = role === 3 && stage.name === "Đề Xuất"
    const notiRegis =  role === 3 && stage.name === "Đăng ký"

    const handleSuggestClick = () => {}
    return (
        <>
            <CustomeProgress setStage={setStage} data={data.stages} stage={stage} />
            {notiSuggest ?
                <div>
                    <span style={{ fontWeight: "bold", color: "#00ace9" }}>
                        <AiTwotoneNotification />Note: 
                    </span>
                    Công việc cần thực hiện trong Giai đoạn đề xuất - <Button onClick={handleSuggestClick}>Duyệt đề xuất</Button>
                </div> : <></>}
            {notiRegis ?
                <div>
                    <span style={{ fontWeight: "bold", color: "#00ace9" }}>
                        <AiTwotoneNotification />Note: 
                    </span>
                    Công việc cần thực hiện trong Giai đoạn đăng ký - <Button onClick={handleSuggestClick}>Duyệt đăng ký</Button>
                </div>
                : <></>}
        </>
    )
}


export default TimeLine;