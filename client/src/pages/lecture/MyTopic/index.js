import { useEffect } from "react";
import Card from "../../../components/Shares/Card";
import "./style.css"
import { useDispatch, useSelector } from "react-redux";
import { fetchTopics } from "../../common/Topic/CommonTopicSlice";
import { userSelector } from "../../../store/selectors";
import { CommonTopicSelector } from "../../../store/selectors";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import Suggest from "../../../components/Shares/Action/Suggest";

const MyTopic = () => {
    const dispatch = useDispatch()
    const user = useSelector(userSelector)
    useEffect(() => {
        dispatch(fetchTopics({ userID: user.id }))
    }, [dispatch, user.id])

    const topicSelector = useSelector(CommonTopicSelector)

    return (
        <div className="lecture_topic">
            <Card title={"Đề tài của tôi"}>
                <Suggest />
                {topicSelector.topics ?
                    <div></div> :
                    <EmptyListNoti title={"Bạn chưa có đề tài nào!"} />}
            </Card>
        </div>
    )
}

export default MyTopic;