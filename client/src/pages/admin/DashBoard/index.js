// import {
//   FaUserAlt,
//   FaUserGraduate,
//   FaUsers,
//   FaHouseUser,
// } from "react-icons/fa";
// import { MdTopic } from "react-icons/md";

import { useDispatch, useSelector } from "react-redux";
import { useEffect, useRef, useState } from "react";
import { Accordion, Button } from "react-bootstrap";

import Card from "../../../components/Shares/Card";
import { AdminEventSelector, userSelector } from "../../../store/selectors";
import { cancelEvent, createEvent, fetchDoneEvents, fetchEvents, updateStage } from "./EventSlice";
import "./style.css";
import { convertDateToTimestamp } from "../../../utils/time";
import Confirm from "../../../components/Shares/Confirm";
import TimeLine from "../../../components/Shares/TimeLine";
import getCurrentStage from "../../../utils/getCurrentStage";

const DashBoard = () => {
	const dispatch = useDispatch()
	const [stage, setStage] = useState() 
	const [doneEvents, setDoneEvent] = useState([])
	const adminEvent = useSelector(AdminEventSelector)
	
	useEffect(() => {
		if (!adminEvent.current.stages) {
			dispatch(fetchEvents())
		}
		if (!adminEvent.doneEvents) {
			dispatch(fetchDoneEvents())
		}
	}, [dispatch, adminEvent.current.stages, adminEvent.doneEvents])
	
	useEffect(() => {
		if (adminEvent.current?.stages) {
			setStage(getCurrentStage(adminEvent.current.stages));
		}
		if (adminEvent?.doneEvent) {
			setDoneEvent(adminEvent.doneEvent)
		}
		// eslint-disable-next-line
	}, [adminEvent.current.stages, adminEvent.doneEvent]);

	const [create, setCreate] = useState(false)
	const descriptionRef = useRef()
	const timeStartRef = useRef()
	const timeEndRef = useRef()
	const nameRef = useRef()
	const yearRef = useRef()
	const [edit, setEdit] = useState(false)
	const [show, setShow] = useState(false)
	const loading = adminEvent.status === "loading"
	const handleConfirmButton = () => {
		if (edit) {
			const description = descriptionRef.current.value
			const timeStart = convertDateToTimestamp(timeStartRef.current.value)
			const timeEnd = convertDateToTimestamp(timeEndRef.current.value)
			dispatch(updateStage({ stageID: stage.id, description: description, timeStart: timeStart, timeEnd: timeEnd }))
			setEdit(false)
			if (!loading) {
				setShow(false)
			}
		}
		else if (create) {
			const name = nameRef.current.value
			const schoolYear = yearRef.current.value
			if (name === "" || schoolYear === "") {

			} else {
				dispatch(createEvent({ name: name, schoolYear: schoolYear }))
			}
			setEdit(false)
		} else {
			dispatch(cancelEvent({ id: adminEvent.current.id }))
		}
		setShow(false)
		setEdit(false)
		setCreate(false)
	}

	const user = useSelector(userSelector)
	let isAdmin = user.role === 5

	return (
		<div className="admin_dash_board">
			<Card title={"NCKH đang diễn ra"}>
				{adminEvent.current && adminEvent.current.stages && stage ?
					<TimeLine descriptionRef={descriptionRef} timeEndRef={timeEndRef} timeStartRef={timeStartRef} setShow={setShow} setEdit={setEdit} edit={edit} data={adminEvent.current} isAdmin={isAdmin} />
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
				<Confirm confirmShow={show} setConfirmShow={setShow} handleConfirmButton={handleConfirmButton} content={"Xác nhận thao tác"} loading={loading}></Confirm>
			</Card>

			<Card title={"NCKH đã thực hiện"}>
				<Accordion defaultActiveKey={0}>
					{doneEvents.map((item, index) => {
						return (
							<Accordion.Item key={item.id} eventKey={index}>
								<Accordion.Header style={{}}>
									{item.name}
								</Accordion.Header>
								<Accordion.Body>
									<TimeLine data={item} />
								</Accordion.Body>
							</Accordion.Item>
						)
					})}
				</Accordion>
			</Card>
		</div >
	);
};

export default DashBoard;
