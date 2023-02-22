const db = require('../config')

const Users = function(user) {
    this.id = user.id,
    this.name = user.name,
    this.password = user.password,
    this.user_type = user.user_type,
    this.created_at = user.created_at,
    this.update_at = user.update_at,
    this.deleted_at = user.deleted_at
}

Users.GetUsers = function(conditions, result) {
    const query = "SELECT * FROM users"
    db.query(query, function(err, rows) {
        if (err) {
            result({
                status: 400,
                error: err
            })
            return
        }
        result({
            status: 200,
            data: rows
        })
    })
}

module.exports = Users;