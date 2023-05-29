import { useDispatch, useSelector } from "react-redux";
import { useEffect, useRef, useState } from "react";
import { Accordion, Button } from "react-bootstrap";

import Card from "../../../components/Shares/Card";
import { AdminEventSelector } from "../../../store/selectors";
import { createEvent, fetchEvents } from "./EventSlice";
import "./style.css";
import TimeLine from "../../../components/Shares/TimeLine";
import Confirm from "../../../components/Shares/Confirm";

const DashBoard = () => {
  const dispatch = useDispatch()
  const adminEvent = useSelector(AdminEventSelector)

  useEffect(() => {
    dispatch(fetchEvents())
  }, [dispatch])


  let currentExists = adminEvent.current && adminEvent.current.stages
  let isLoading = adminEvent.status === 'loading'

  const nameRef = useRef()
  const yearRef = useRef()
  const [name, setName] = useState("")
  const [year, setYear] = useState("")
  const [create, setCreate] = useState(false)

  return (
    <div className="admin_dash_board">
      <Card title={"NCKH đang diễn ra"}>

        {currentExists ?
          <TimeLine data={adminEvent.current} />
          :
          <>
            {create
              ?
              <div className="create-event">
                <div className="row margin_ver">
                  <div className="col-2 detail-head">NCKH:</div>
                  <div className="col-6">
                    <input
                      type="text"
                      placeholder="Tên NCKH"
                      style={{ width: "400px" }}
                      className="input"
                      value={name}
                      onChange={(e) => setName(e.target.value)}
                    />
                  </div>
                </div>
                <div className="row margin_ver">
                  <div className="col-2 detail-head">Năm học: </div>
                  <div className="col-6">
                    <input
                      type="text"
                      placeholder="Năm học"
                      className="input"
                      value={year}
                      onChange={(e) => setYear(e.target.value)}
                    />
                  </div>
                </div>
                <div className="row margin_ver">
                  <div className="col-2 detail-head"></div>
                  <div className="col-6">
                    <Button style={{ marginRight: "12px" }} onClick={() => setCreate(false)}>Hủy</Button>
                    <Confirm
                      title={"Xác nhận"}
                      content={"Xác nhận tạo mới nghiên cứu khoa học"}
                      isLoading={isLoading}
                      action={createEvent({ name: name, schoolYear: year })}
                    />
                  </div>
                </div>
                <div>
                </div>
              </div>
              :
              <div className="no-current">
                <div className="bold margin_ver">Không có NCKH đang được thực hiện</div>
                <Button onClick={() => setCreate(true)}>Bắt đầu NCKH mới</Button>
              </div>
            }
          </>
        }
      </Card>

      {/* <Card title={"NCKH đã thực hiện"}>
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
			</Card> */}
    </div >
  );
};

export default DashBoard;
