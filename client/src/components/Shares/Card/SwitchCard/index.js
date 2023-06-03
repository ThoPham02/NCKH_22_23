import { Form } from "react-bootstrap"

const SwitchCard = (props) => {
    const { buttonSwitch, setButtonSwitch, title1, title2, children } = props
    const handleChange = () => {
        setButtonSwitch(!buttonSwitch)
    }
    return (
        <div className="card container-fluid">
            <div className="card-header" >
                <div className={buttonSwitch ? "card-switch" : ""}>{title1}</div>
                <Form style={{marginLeft: "10px"}}>
                    <Form.Check
                        type="switch"
                        id="custom-switch"
                        onChange={handleChange}
                    />
                </Form>
                <div className={!buttonSwitch ? "card-switch" : ""}>{title2}</div>
            </div>
            <div style={{ padding: "24px" }}>{children}</div>
        </div>
    )
}

export default SwitchCard;