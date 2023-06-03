import { useState } from "react";
import { Form } from "react-bootstrap";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import { TbFilePlus } from "react-icons/tb"
import Confirm from "../../Confirm";
import { addTopic } from "../../../../pages/common/Topic/CommonTopicSlice";
import { useSelector } from "react-redux";
import { userSelector } from "../../../../store/selectors";

const Suggest = () => {
    const user = useSelector(userSelector)

    const [show, setShow] = useState(false);
    const [num, setNum] = useState(5);
    const [name, setName] = useState("")
    const [description, setDescription] = useState("")

    const handleClose = () => setShow(false);
    const handleClick = () => {

    }
    return (
        <>
            <Button onClick={() => setShow(true)}>
                <TbFilePlus />
                Thêm đề xuất
            </Button>

            <Modal show={show} onHide={handleClose}>
                <Modal.Header closeButton>
                    <Modal.Title>Đề xuất đề tài</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form>
                        <Form.Group className="mb-3" >
                            <Form.Label>Tên đề tài</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="Nhập tên đề tài"
                                autoFocus
                                value={name}
                                onChange={e => setName(e.target.value)}
                            />
                        </Form.Group>
                        <Form.Group className="mb-3">
                            <Form.Label>Số lượng sinh viên thực hiện dự tính:</Form.Label>
                            <Form.Select value={num} onChange={e => setNum(e.target.value)}>
                                <option value="5">5</option>
                                <option value="4">4</option>
                                <option value="3">3</option>
                                <option value="2">2</option>
                            </Form.Select>
                        </Form.Group>
                        <Form.Group className="mb-3">
                            <Form.Label>Mô tả:</Form.Label>
                            <Form.Control as="textarea" rows={4} placeholder="Mô tả thêm về đề tài" value={description} onChange={(e) => setDescription(e.target.value)} />
                        </Form.Group>
                    </Form>
                </Modal.Body>
                <Modal.Footer className="modal-footer">
                    <Button variant="secondary" onClick={handleClose}>
                        Hủy
                    </Button>
                    <Confirm
                        title="Đăng ký"
                        content="Xác nhận đề xuất đề tài"
                        action={addTopic({ name: name, lectureID: user.id, departmentID: user.department_id, estimateStudent: num, description: description })}
                        onClick={handleClick}
                    />
                </Modal.Footer>
            </Modal>
        </>
    );
};

export default Suggest;
