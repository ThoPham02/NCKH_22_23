import { useState } from "react";
import { Button, Form, Modal } from "react-bootstrap";

import Confirm from "../../Confirm";
import { addReport } from "../../../../pages/lecture/MyTopic/LectureMyTopicSlice";

const AddReport = ({id}) => {
    const [show, setShow] = useState(false);
    const [name, setName] = useState("");

    return (
        <>
            <Button onClick={(() => setShow(true))} className="button">
                Báo cáo
            </Button>

            <Modal show={show} onHide={() => setShow(false)} size="lg" centered >
                <Modal.Header closeButton>
                    <Modal.Title>Thêm Báo Cáo</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form>
                        <Form.Group style={{ margin: "0 20px" }}>
                            <Form.Label>Mô tả</Form.Label>
                            <Form.Control as="textarea" rows={3} value={name} onChange={(e) => setName(e.target.value)} />
                        </Form.Group>
                        <Form.Group style={{ margin: "0 20px" }}>
                            <Form.Label>File đính kèm</Form.Label><br/>
                            <input type="file" />
                        </Form.Group>
                    </Form>
                </Modal.Body>
                <Modal.Footer className="modal-footer">
                    <Button variant="secondary" onClick={() => setShow(false)}>
                        Hủy
                    </Button>
                    <Confirm
                        title="Thêm báo cáo"
                        content="Xác nhận nộp báo cáo"
                        action={addReport({
                            id: id,
                            description: name,
                            url: "",
                            stageID: 0
                        })}
                        onClick={() => setShow(false)}
                        isLoading={false}
                    />
                </Modal.Footer>
            </Modal>
        </>
    )
}

export default AddReport;