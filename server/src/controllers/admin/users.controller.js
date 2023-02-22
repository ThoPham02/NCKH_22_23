const Users = require('../../models/users/users.models')

module.exports = {
    getUser: function(req, res) {
        const conditions = {}
        Users.GetUsers(conditions, function(result) {
            if (result.status === 400) {
                res.status(400).json({
                    message: "Bad request"
                })
                return
            }
            res.status(200).json({
                message: "OK",
                data: result.data
            })
        })
    },
    createUser: function (req, res) {

    },
    updateUser: function (req, res) {

    },
    deleteUser: function (req, res) {

    }
}