import { useState } from "react";
import { Button, Form, Modal } from "react-bootstrap";
import { TbFilePlus } from "react-icons/tb";
import Confirm from "../../Confirm";

const AddSubcommittee = () => {
    const [show, setShow] = useState(false);
    const handleShow = () => {
        setShow(true);
    }
    return (
        <>
            <Button onClick={handleShow}>
                <TbFilePlus />
                Thêm tiểu ban
            </Button>

            <Modal show={show} onHide={() => setShow(false)} size="xl" centered>
                <Modal.Header closeButton>
                    <Modal.Title>Thêm Tiểu Ban</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form>

                    </Form>
                </Modal.Body>
                <Modal.Footer className="modal-footer">
                    <Button variant="secondary" onClick={() => setShow(false)}>
                        Hủy
                    </Button>
                    <Confirm
                        title="Đăng ký"
                        content="Xác nhận thêm mới tiểu ban"
                        action={{}}
                        onClick={{}}
                        isLoading={false}
                    />
                </Modal.Footer>
            </Modal>
        </>
    )
}

export default AddSubcommittee;