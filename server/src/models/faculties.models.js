const db = require("./config");

const Faculties = function (faculties) {
  (this.id = faculties.id),
    (this.name = faculties.name),
    (this.user_id = faculties.user_id),
    (this.created_at = faculties.created_at),
    (this.update_at = faculties.update_at),
    (this.deleted_at = faculties.deleted_at);
};

Faculties.GetFaculties = function (conditions, result) {
  var query = "SELECT * FROM faculties Where deleted_at is null";

  db.query(query, function (err, rows) {
    if (err) {
      result(err, null);
      return;
    }
    result(null, rows);
  });
};

Faculties.Create = function (data, result) {
  db.query("INSERT INTO faculties SET ?", data, function (err, user) {
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

Faculties.Delete = function (id, result) {
  db.query(
    "UPDATE faculties SET deleted_at = NOW() Where id = ?",
    id,
    function (err, _) {
      if (err) {
        result(err, null);
      } else {
        result(null, null);
      }
    }
  );
};

Faculties.Update = function (data, result) {
  db.query(
    "UPDATE faculties SET update_at = NOW(), name = ? WHERE id = ?", [data.name, data.id],
    function (err, _) {
      if (err) {
        result(err, null);
      } else {
        result(null, null);
      }
    }
  );
};

module.exports = Faculties;
