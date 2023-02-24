const db = require("./config");

const Lecturers = function (lecturer) {
  (this.id = lecturer.id),
    (this.name = lecturer.name),
    (this.user_id = lecturer.user_id),
    (this.created_at = lecturer.created_at),
    (this.update_at = lecturer.update_at),
    (this.deleted_at = lecturer.deleted_at);
};

Lecturers.GetLectures = function (conditions, result) {
  var query = "SELECT * FROM lecturers Where deleted_at is null";

  if (conditions.user_id) {
    query += " AND user_id = " + conditions.user_id;
  }
  if (conditions.name) {
    query += ` AND name = '${conditions.name}%'`
  }

  db.query(query, function (err, rows) {
    if (err) {
      result(err, null);
      return;
    }
    result(null, rows);
  });
};

Lecturers.Create = function (data, result) {
  db.query("INSERT INTO lecturers SET ?", data, function (err, user) {
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

Lecturers.Delete = function (id, result) {
  db.query("UPDATE lecturers SET deleted_at = NOW() Where id = ?", id, function (err, _) {
    if (err) {
      result(err, null);
    } else {
      result(null , null);
    }
  });
};

Lecturers.Update = function (data, result) {
  db.query(
    "UPDATE lecturers SET update_at = NOW(), name = ?, user_id = ? WHERE id = ?",
    [data.name, data.user_id, data.id],
    function (err, _) {
      if (err) {
        result(err, null);
      } else {
        result(null, null);
      }
    }
  );
};

module.exports = Lecturers;
