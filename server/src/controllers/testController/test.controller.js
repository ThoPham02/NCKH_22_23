module.exports = {
    testController:(req, res) => {
        res.status(200).json({
            message: "Success"
        })
    },
    test2Controller: (req, res) => {
        res.status(200).json({
            message: "Success2"
        })
    }
}