import { useSelector } from "react-redux"
import { useState } from "react"

import { AdminEventSelector } from "../../../store/selectors"
import EmptyListNoti from "../../../components/Shares/EmptyListNoti"
import SwitchCard from "../../../components/Shares/Card/SwitchCard"
import CurrentTopic from "./CurrentTopic"
import DoneTopic from "./DoneTopic"
import { SubTimeLine } from "../../../components/Shares/TimeLine"

const Topic = () => {
    const currentEvent = useSelector(AdminEventSelector)
    const [switchPage, setSwitchPage] = useState(false)

    console.log(currentEvent)

    return (
        <SwitchCard switchPage={switchPage} setSwitchPage={setSwitchPage}>
            {!switchPage
                ?
                currentEvent && currentEvent.current && currentEvent.current.stages && currentEvent.current.stages.length !== 0?
                    <>
                        <SubTimeLine data={currentEvent.current}/>
                        <CurrentTopic />
                    </>
                    :
                    <EmptyListNoti title={"Hiện không có đề tài nào đang được thực hiện"} />
                :
                <DoneTopic />
            }
        </SwitchCard >
    )
}

export default Topic;