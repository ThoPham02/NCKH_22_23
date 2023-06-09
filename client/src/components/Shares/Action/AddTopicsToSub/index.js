import { useEffect, useState } from "react";
import { Button, Form, Modal, Table } from "react-bootstrap";
import { TbFilePlus } from "react-icons/tb";
import Confirm from "../../Confirm";
import { useDispatch, useSelector } from "react-redux";
import { FacultySubcommittee } from "../../../../store/selectors";
import { addTopicsToSub, fetchTopicsBySubcommittee } from "../../../../pages/faculty/Report/facultySubcommitteeSlice";
import EmptyListNoti from "../../EmptyListNoti";

const AddTopicsToSub = ({ subcommittee }) => {
    const [show, setShow] = useState(false);
    const dispatch = useDispatch()
    const handleShow = () => {
        dispatch(fetchTopicsBySubcommittee({subcommitteeID: subcommittee.id}))
        setShow(true);
    }
    const report = useSelector(FacultySubcommittee)
    console.log(report)
    const [select, setSelect] = useState([])
    return (
        <>
            <Button onClick={handleShow} style={{
                position: "absolute",
                zIndex: "10",
                right: "80px",
                padding: "0 12px",
                marginTop: "14px"
            }}>
                <TbFilePlus />
                Thêm đề tài
            </Button>

            <Modal
                show={show}
                onHide={() => setShow(false)}
                size="lg"
                centered
                backdrop="static"
                keyboard={false}>
                <Modal.Header closeButton>
                    <Modal.Title>Chọn đề tài phù hợp với tiểu ban</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <div className="row" style={{ fontWeight: "bold" }}>
                        <div className="col-2 ">Tiểu ban: </div>
                        <div className="col-4">{subcommittee.name}</div>
                    </div>
                    {report.topics && report.topics.length !== 0 && report.topics.filter(item => item.subcommitteeID.id === 0).length !== 0?
                        <TableTopic select={select} setSelect={setSelect} topic={report.topics.filter(item => item.subcommitteeID.id === 0)}/>
                        :
                        <EmptyListNoti title={"Danh sách đề tài phù hợp trống!"} />}
                </Modal.Body>
                <Modal.Footer className="modal-footer">
                    <Button variant="secondary" onClick={() => setShow(false)}>
                        Hủy
                    </Button>
                    <Confirm
                        title="Xác Nhận"
                        content="Xác nhận thêm đè tài vào tiểu ban"
                        action={addTopicsToSub({
                            id: subcommittee.id,
                            listTopicID: select,
                            facultyID: subcommittee.facultyID
                        })}
                        onClick={{}}
                        isLoading={false}
                    />
                </Modal.Footer>
            </Modal>
        </>
    )
}

const TableTopic = ({select, setSelect, topic}) => {
    const [selectAll, setSelectAll] = useState(false)
    const handleClick = (value) => {
        if (select.includes(value)) {
            setSelect(select.filter((item) => item !== value));
        } else {
            setSelect([...select, value]);
        }
    }
    useEffect(() => {
        if (select.length === topic.length) {
            setSelectAll(true)
        } else {
            setSelectAll(false)
        }
    }, [select, topic])
     const handleClickAll = () => {
        if (selectAll) {
            setSelect([])
            setSelectAll(false)
        }
        else {
            setSelectAll(true)
            console.log(topic.map(item => item.id))
            setSelect(topic.map(item => item.id))
        }
    }
    return (
        <Table striped hover size="sm">
            <thead>
                <tr>
                    <th style={{ width: "100px" }}>Chọn tất cả
                        <Form.Check
                            type={"checkbox"}
                            checked={selectAll}
                            onClick={handleClickAll}
                        />
                    </th>
                    <th>Đề tài<br /></th>
                </tr>
            </thead>
            <tbody>
                {topic.map(item => {
                    return (
                        <tr>
                            <td style={{ textAlign: "center" }}><Form.Check
                                type={"checkbox"}
                                id={`default-checkbox`}
                                checked={select.includes(item.id)}
                                onClick={() => handleClick(item.id)}
                            />
                            </td>
                            <td>{item.name}</td>
                        </tr>
                    )
                })}
            </tbody>
        </Table>
    )
}

export default AddTopicsToSub;