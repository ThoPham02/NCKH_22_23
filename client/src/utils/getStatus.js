import { topicStatus } from "../const/const"

export const getStatus = (status) => {
    const result = topicStatus.find(item => item.id === status)
    if (result) {
        return result.name
    } else {
        return "Hủy"
    }
}