const Users = require('../../models/users/users.models')

module.exports = {
    getUser: function(req, res) {
        const conditions = {}
        Users.GetUsers(conditions, function(result) {
            if (typeof(result) === 'string') {
                res.status(400).json({
                    message: "Bad request"
                })
                return
            }
            res.status(200).json({
                message: "OK",
                data: result
            })
        })
    }
}