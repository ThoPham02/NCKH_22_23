import { useDispatch, useSelector } from "react-redux";
import { useEffect, useRef, useState } from "react";
import { Accordion } from "react-bootstrap";

import Card from "../../../components/Shares/Card";
import { AdminEventSelector, userSelector } from "../../../store/selectors";
import { fetchDoneEvents, fetchEvents } from "../../admin/DashBoard/EventSlice";
import "./style.css";
import TimeLine from "../../../components/Shares/TimeLine";
import getCurrentStage from "../../../utils/getCurrentStage";

const DashBoard = () => {
  const dispatch = useDispatch()
  const [stage, setStage] = useState()
  const [doneEvents, setDoneEvent] = useState([])
  const adminEvent = useSelector(AdminEventSelector)
  useEffect(() => {
    if (!adminEvent.current) {
      dispatch(fetchEvents())
    }
    if (!adminEvent.doneEvents) {
      dispatch(fetchDoneEvents())
    }
    
    if (adminEvent.current) {
      setStage(getCurrentStage(adminEvent.stages));
    }
    if (adminEvent.doneEvent) {
      setDoneEvent(adminEvent.doneEvent)
    }
    // eslint-disable-next-line
  }, [dispatch, adminEvent.current, adminEvent.doneEvent]);

  const descriptionRef = useRef()
  const timeStartRef = useRef()
  const timeEndRef = useRef()

  const [edit, setEdit] = useState(false)
  // eslint-disable-next-line
  const [show, setShow] = useState(false)
  const user = useSelector(userSelector)
  let isAdmin = user.role === 5

  return (
    <div className="admin_dash_board">
      <Card title={"NCKH đang diễn ra"}>
        {adminEvent.current && adminEvent.stages && stage ?
          <TimeLine descriptionRef={descriptionRef} timeEndRef={timeEndRef} timeStartRef={timeStartRef} setShow={setShow} setEdit={setEdit} edit={edit} data={adminEvent.current} isAdmin={isAdmin} />
          :
          <div className="no-current">
            <div className="bold margin_ver">Không có NCKH đang được thực hiện</div>
          </div>
        }
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
