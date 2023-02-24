const db = require("./config");

const Department = function (department) {
  (this.id = department.id),
    (this.name = department.name),
    (this.faculty_id = department.faculty_id),
    (this.user_id = department.user_id),
    (this.created_at = department.created_at),
    (this.update_at = department.update_at),
    (this.deleted_at = department.deleted_at);
};

Department.GetDepartments = function (conditions, result) {
  var query = "SELECT * FROM departments Where deleted_at is null";

  if (conditions.name) {
    query += `AND name like '${conditions.name}';`
  }
  if (conditions.faculty_id) {
    query += " AND faculty_id = " + conditions.faculty_id;
  }
  if (conditions.user_id) {
    query += " AND user_id = " + conditions.user_id;
  }

  db.query(query, function (err, rows) {
    if (err) {
      result(err, null);
      return;
    }
    result(null, rows);
  });
};

Department.Create = function (data, result) {
  db.query("INSERT INTO departments SET ?", data, function (err, user) {
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

Department.Delete = function (id, result) {
  db.query("UPDATE departments SET deleted_at = NOW() Where id = ?", id, function (err, _) {
    if (err) {
      result(err, null);
    } else {
      result(null , null);
    }
  });
};

Department.Update = function (data, result) {
  db.query(
    "UPDATE departments SET update_at = NOW(), name = ?, faculty_id = ?, user_id = ? WHERE id = ?",
    [data.name, data.faculty_id, data.user_id, data.id],
    function (err, _) {
      if (err) {
        result(err, null);
      } else {
        result(null, null);
      }
    }
  );
};

module.exports = Department;
