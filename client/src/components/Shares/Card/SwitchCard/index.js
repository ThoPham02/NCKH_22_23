import { Form } from "react-bootstrap"

const SwitchCard = (props) => {
    const { switchPage, setSwitchPage, children } = props
    const handleChange = () => {
        setSwitchPage(!switchPage)
    }
    return (
        <div className="card container-fluid">
            <div className="card-header" >
                <div className={switchPage ? "card-switch" : ""}>NCKH đang diễn ra</div>
                <Form style={{marginLeft: "10px"}}>
                    <Form.Check
                        type="switch"
                        id="custom-switch"
                        onChange={handleChange}
                    />
                </Form>
                <div className={!switchPage ? "card-switch" : ""}>NCKH đã hoàn thành</div>
            </div>
            <div style={{ padding: "24px" }}>{children}</div>
        </div>
    )
}

export default SwitchCard;