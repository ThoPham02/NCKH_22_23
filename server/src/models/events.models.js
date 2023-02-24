const db = require('./config')

const Users = function(user) {
    this.id = user.id,
    this.username = user.username,
    this.password = user.password,
    this.user_type_id = user.user_type_id,
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

Users.create = function(data, result) {
    db.query("INSERT INTO users SET ?", data, function(err,user){
        if(err){
            if(err){
                result(err);
            }
        }
        else {
            result({
                id: user.insertId,
                ...data,
            });
        }
    })
}

Users.create = function(data, result) {
    db.query("INSERT INTO users SET ?", data, function(err,user){
        if(err){
            if(err){
                result(err);
            }
        }
        else {
            result({
                id: user.insertId,
                ...data,
            });
        }
    })
}

Users.remove = function(id, result) {
    db.query("DELETE FROM users WHERE id = ?", id, function(err,user){
        if(err){
            result(err);
        }
        else {
            result("đã xóa "+ id);
        }
    })
}

Users.update = function(data, result) {
    db.query("UPDATE users SET username = ?, password = ?, user_type_id = ? WHERE id = ?",
    [
        data.username,
        data.password,
        data.user_type_id,
        data.id
    ],
    function(err,user){
        if(err){
            result(err);
        }
        else {
            result("dữ liệu đã được update");
        }
    })
}

module.exports = Users;