export const getDegree = (degree) => {
    const degrees = [{ id: 1, name: "SV" }, { id: 2, name: "Ths"}, {id: 4, name: "TS"}, {id: 8, name: "PGS"}, {id: 16, name: "GS"}]
    const item = degrees.filter(item => item.id === degree)
    return item[0].name
}