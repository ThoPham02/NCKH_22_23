import { Button } from "react-bootstrap";
import "./style.css"
import CustomeProgress from "./CustomeProgress";
import { useState } from "react";
import { convertTimestampToDateString } from "../../../utils/time";

const TimeLine = (props) => {
    const { data, setShow, edit, setEdit, descriptionRef, timeStartRef, timeEndRef, isAdmin } = props

    let now = 0
    let current = new Date().getTime();
    for (let i = 0; i < data.stages.length; i++) {
        if (data.stages[i].timeStart !== 0 && data.stages[i].timeStart <= current) {
            now = i
        }
    }
    const [stage, setStage] = useState(data.stages[now])

    let isCurrent = isAdmin && data.isCurrent

    return (
        <div>
            <div className="row">
                <div className="col-2 detail-head">NCKH:</div>
                <div className="col-6">{data.name}</div>
                {isCurrent ? <div className="col-4"><Button onClick={() => setShow(true)} variant="danger">Kết thúc NCKH</Button></div> : <></>}
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
                        {edit ? <input type="text" placeholder="Mô tả chung về giai đoạn" style={{ width: "600px" }} className="input" ref={descriptionRef} /> : <>{stage.description}</>}
                    </div>
                </div>
                <div className="row detail-item">
                    <div className="col-2 detail-head">Thời gian:</div>
                    <div className="col-8">
                        Từ 
                        {edit ? <input type="date" ref={timeStartRef} /> : <>{stage.timeStart === 0 ? " ... " : convertTimestampToDateString(stage.timeStart)}</>}
                        đến
                        {edit ? <input type="date" ref={timeEndRef} /> : <>{stage.timeEnd === 0 ? " ... " : convertTimestampToDateString(stage.timeEnd)}</>}
                    </div>
                </div>
                <div className="row detail-item">
                    <div className="col-2 detail-head"></div>
                    <div className="col-8" style={{ fontWeight: "bold" }}>{edit ? <input type="file" /> : <>
                        {
                            // eslint-disable-next-line
                            stage.url ? <a href="#">Tài liệu đính kèm</a> : <></>
                        }</>}</div>
                </div>
                {isCurrent
                    ?
                    <div className="row detail-item">
                        <div className="col-2 detail-head"></div>
                        <div className="col-8" style={{ fontWeight: "bold", marginTop: "12px" }}>
                            {edit ? <div><Button style={{ marginRight: "12px" }} onClick={() => setEdit(false)}>Hủy</Button> <Button onClick={() => setShow(true)}>Xác nhận</Button></div> : <Button onClick={() => setEdit(true)}>Chỉnh sửa</Button>}
                        </div>
                    </div>
                    : <></>}
            </div>
        </div>
    )
}


export default TimeLine;