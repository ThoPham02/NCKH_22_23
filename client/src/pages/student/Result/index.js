import { useState } from "react"
import SwitchCard from "../../../components/Shares/Card/SwitchCard"
import "./style.css"

const Result = () => {
    const [switchPage, setSwitchPage] = useState(false)
    return (
        <SwitchCard switchPage={switchPage} setSwitchPage={setSwitchPage}>
            {switchPage ?
            <>
                
            </>
            : 
            <>
            </>
            }
        </SwitchCard>
    )
}

export default Result