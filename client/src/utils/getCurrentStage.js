const getCurrentStage = (data) => {
    let now = 0
    let current = new Date().getTime();
    for (let i = 0; i < data.length; i++) {
        if (data[i].timeStart !== 0 && data[i].timeStart <= current) {
            now = i
        }
    }
    return data[now]
}
export default getCurrentStage;