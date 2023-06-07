import { useEffect, useState } from "react";
import { Button, Form, Modal } from "react-bootstrap";
import { TbFilePlus } from "react-icons/tb";
import Confirm from "../../Confirm";
import SubCard from "../../Card/SubCard";
import { useDispatch, useSelector } from "react-redux";
import { AdminEventSelector, itemSelector, userSelector } from "../../../../store/selectors";
import { ItemAction, fetchTeacher } from "./itemSlice";
import { addSubcommittee } from "../../../../pages/faculty/Report/facultySubcommitteeSlice";
import { getDegree } from "../../../../utils/getDegree";

const AddSubcommittee = () => {
    const [show, setShow] = useState(false);
    const [name, setName] = useState("")
    const dispatch = useDispatch()
    const user = useSelector(userSelector)
    useEffect(() => {
        dispatch(fetchTeacher({ facultyID: user.faculty_id }))
    }, [dispatch, user])
    const ListMem = useSelector(itemSelector)
    const event = useSelector(AdminEventSelector).current

    const handleShow = () => {
        setShow(true);
    }
    const handleAddItem = () => {
        dispatch(ItemAction.addItem())
    }
    const handleClose = () => {
        dispatch(ItemAction.resetItem())
        setName("")
        setShow(false)
    }
    return (
        <>
            <Button onClick={handleShow}>
                <TbFilePlus />
                Thêm tiểu ban
            </Button>

            <Modal show={show} onHide={handleClose} size="lg" centered >
                <Modal.Header closeButton>
                    <Modal.Title>Thêm Tiểu Ban</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form>
                        <Form.Group style={{ margin: "0 20px" }}>
                            <Form.Label>Tên tiểu ban</Form.Label>
                            <Form.Control type="text" value={name} onChange={(e) => setName(e.target.value)} />
                        </Form.Group>
                        <SubCard title={"Thành viên hội đồng nghiệm thu:"}>
                            {
                                ListMem.item.map((item, index) => {
                                    return <InputMem item={item} key={index} />
                                })
                            }
                            {ListMem.item.length < 5 ?
                                <Button onClick={handleAddItem} style={{
                                    padding: "0 12px",
                                    margin: "0 20px 0 0",
                                    float: "right",
                                }}>
                                    Thêm
                                </Button> : <></>}
                        </SubCard>
                    </Form>
                </Modal.Body>
                <Modal.Footer className="modal-footer">
                    <Button variant="secondary" onClick={handleClose}>
                        Hủy
                    </Button>
                    <Confirm
                        title="Đăng ký"
                        content="Xác nhận thêm mới tiểu ban"
                        action={addSubcommittee({
                            name: name,
                            listLectures: ListMem.item.map(item => {
                                const { lectureID, role } = item;
                                return { lectureID: lectureID * 1, role: role * 1 };
                            }),
                            facultyID: user.faculty_id,
                            eventID: event.id,
                        })}
                        onClick={() => setShow(false)}
                        isLoading={ListMem.status === "loading"}
                    />
                </Modal.Footer>
            </Modal>
        </>
    )
}

const InputMem = ({ item }) => {
    const [role, setRole] = useState(1)
    const [teacher, setTeacher] = useState(0)
    const ListMem = useSelector(itemSelector)
    const dispatch = useDispatch()
    const handleChangle = () => {
        console.log(teacher)
        dispatch(ItemAction.changeItem({
            index: item.index,
            lectureID: teacher,
            lectureName: "",
            role: role,
        }))
    }
    return (
        <Form.Group style={{ display: "flex", margin: "20px 0" }} onChange={handleChangle}>
            <Form.Select value={teacher} onChange={(e) => setTeacher(e.target.value)} style={{ marginLeft: "20px" }}>
                <option value={0}>Chọn giảng viên</option>
                {ListMem.teacher ?
                    ListMem.teacher.map((item) => {
                        return <option value={item.id} key={item.id}>{getDegree(item.degree) + "." + item.name}</option>
                    })
                    : <></>}
            </Form.Select>
            <Form.Select value={role} onChange={(e) => setRole(e.target.value)} style={{ width: "200px", margin: "0 20px" }}>
                <option value="1">Chủ tịch hội đồng</option>
                <option value="2">Ủy viên</option>
                <option value="3">Thư ký</option>
            </Form.Select>
        </Form.Group>
    )
}

export default AddSubcommittee;