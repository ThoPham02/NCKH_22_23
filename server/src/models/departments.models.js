const db = require("./config");

const Users = function (user) {
  (this.id = user.id),
    (this.username = user.username),
    (this.password = user.password),
    (this.user_type_id = user.user_type_id),
    (this.created_at = user.created_at),
    (this.update_at = user.update_at),
    (this.deleted_at = user.deleted_at);
};

Users.GetUsers = function (conditions, result) {
  var query = "SELECT * FROM users Where deleted_at is null";

  if (conditions.user_type_id) {
    query += " AND user_type_id = " + conditions.user_type_id;
  }

  db.query(query, function (err, rows) {
    if (err) {
      result(err, null);
      return;
    }
    result(null, rows);
  });
};

Users.Create = function (data, result) {
  db.query("INSERT INTO users SET ?", data, function (err, user) {
    if (err) {
      if (err) {
        result(err, null);
      }
    } else {
      result(null, {
        id: user.insertId,
        ...data,
      });
    }
  });
};

Users.Delete = function (id, result) {
  db.query("UPDATE users SET deleted_at = NOW() Where id = ?", id, function (err, user) {
    if (err) {
      result(err, null);
    } else {
      result(null , null);
    }
  });
};

Users.Update = function (data, result) {
  db.query(
    "UPDATE users SET update_at = NOW(), username = ?, password = ?, user_type_id = ? WHERE id = ?",
    [data.username, data.password, data.user_type_id, data.id],
    function (err, _) {
      if (err) {
        result(err, null);
      } else {
        result(null, null);
      }
    }
  );
};

module.exports = Users;
