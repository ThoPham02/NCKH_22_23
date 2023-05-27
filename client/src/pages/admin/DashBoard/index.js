// import {
//   FaUserAlt,
//   FaUserGraduate,
//   FaUsers,
//   FaHouseUser,
// } from "react-icons/fa";
// import { MdTopic } from "react-icons/md";

import { useDispatch, useSelector } from "react-redux";
import { useEffect, useRef, useState } from "react";
import { Button } from "react-bootstrap";

import Card from "../../../components/Shares/Card";
import { AdminEventSelector } from "../../../store/selectors";
import { createEvent, fetchEvents, updateStage } from "./EventSlice";
import "./style.css";
import CustomeProgress from "../../../components/Shares/CustomeProgress";
import { convertDateToTimestamp, convertTimestampToDateString } from "../../../utils/time";
import Confirm from "../../../components/Shares/Confirm";

const DashBoard = () => {
  const dispatch = useDispatch()
  useEffect(() => {
    dispatch(fetchEvents())
  }, [dispatch])
  const adminEvent = useSelector(AdminEventSelector)
  const [stage, setStage] = useState(adminEvent.current ? adminEvent.current.stages[0] : 0)

  const [edit, setEdit] = useState(false)
  const descriptionRef = useRef()
  const timeStartRef = useRef()
  const timeEndRef = useRef()
  const nameRef = useRef()
  const yearRef = useRef()
  const [show, setShow] = useState(false)
  const handleConfirmButton = () => {
    if (edit) {
      const description = descriptionRef.current.value
      const timeStart = convertDateToTimestamp(timeStartRef.current.value)
      const timeEnd = convertDateToTimestamp(timeEndRef.current.value)
      dispatch(updateStage({ stageID: stage, description: description, timeStart: timeStart, timeEnd: timeEnd }))
      setEdit(false)
    }
    if (create) {
      const name = nameRef.current.value
      const schoolYear = yearRef.current.value
      if (name === "" || schoolYear === "") {

      } else {
        dispatch(createEvent({ name: name, schoolYear: schoolYear }))
      }
      setEdit(false)
    }
    setShow(false)
  }

  const [create, setCreate] = useState(false)

  return (
    <div className="admin_dash_board">
      <Card title={"NCKH đang diễn ra"}>
        {adminEvent.result && adminEvent.result.code === 0 ?
          <div>
            <div className="row">
              <div className="col-2 detail-head">NCKH:</div>
              <div className="col-6">{adminEvent.current.name}</div>
            </div>
            <div className="row">
              <div className="col-2 detail-head">Năm học: </div>
              <div className="col-6">{adminEvent.current.schoolYear}</div>
            </div>
            <CustomeProgress stage={stage} setStage={setStage} data={adminEvent.current.stages} />
            <div className="stage-detail">
              <div style={{ fontWeight: "bold" }}>Thông tin giai đoạn:</div>
              <div className="row detail-item">
                <div className="col-2 detail-head">Giai đoạn:</div>
                <div className="col-8">{stage.name}</div>
              </div>
              <div className="row detail-item">
                <div className="col-2 detail-head">Mô tả:</div>
                <div className="col-8">{edit ? <input type="text" placeholder="Mô tả chung về giai đoạn" style={{ width: "600px" }} className="input" ref={descriptionRef} /> : <>{stage.description}</>}</div>
              </div>
              <div className="row detail-item">
                <div className="col-2 detail-head">Thời gian:</div>
                <div className="col-8">Từ {edit ? <input type="date" ref={timeStartRef} /> : <>{stage.timeStart === 0 ? " ... " : convertTimestampToDateString(stage.timeStart)}</>} đến {edit ? <input type="date" ref={timeEndRef} /> : <>{stage.timeEnd === 0 ? " ... " : convertTimestampToDateString(stage.timeEnd)}</>}</div>
              </div>
              <div className="row detail-item">
                <div className="col-2 detail-head"></div>
                <div className="col-8" style={{ fontWeight: "bold" }}>{edit ? <input type="file" /> : <>
                  {
                    // eslint-disable-next-line
                    stage.url ? <a href="#">Tài liệu đính kèm</a> : <></>
                  }</>}</div>
              </div>
              <div className="row detail-item">
                <div className="col-2 detail-head"></div>
                <div className="col-8" style={{ fontWeight: "bold", marginTop: "12px" }}>
                  {edit ? <div><Button style={{ marginRight: "12px" }} onClick={() => setEdit(false)}>Hủy</Button> <Button onClick={() => setShow(true)}>Xác nhận</Button></div> : <Button onClick={() => setEdit(true)}>Chỉnh sửa</Button>}
                </div>
              </div>
            </div>
          </div>
          :
          <>
            {create
              ?
              <div className="create-event">
                <div className="row margin_ver">
                  <div className="col-2 detail-head">NCKH:</div>
                  <div className="col-6"><input type="text" placeholder="Tên NCKH" style={{ width: "400px" }} className="input" ref={nameRef} /></div>
                </div>
                <div className="row margin_ver">
                  <div className="col-2 detail-head">Năm học: </div>
                  <div className="col-6"><input type="text" placeholder="Năm học" className="input" ref={yearRef} /></div>
                </div>
                <div className="row margin_ver">
                  <div className="col-2 detail-head"></div>
                  <div className="col-6">
                    <Button style={{ marginRight: "12px" }} onClick={() => setCreate(false)}>Hủy</Button> <Button onClick={() => setShow(true)}>Xác nhận</Button>
                  </div>
                </div>
                <div>
                </div>
              </div>
              :
              <div className="no-current">
                <div className="bold margin_ver">Không có NCKH đang được thực hiện</div>
                <Button onClick={() => setCreate(true)}>Bắt đầu NCKH mới</Button>
              </div>}
          </>
        }
        <Confirm confirmShow={show} setConfirmShow={setShow} handleConfirmButton={handleConfirmButton} content={"Xác nhận thao tác"}></Confirm>
      </Card>

      <Card title={"NCKH đã thực hiện"}>
        
      </Card>
    </div >
  );
};

export default DashBoard;
