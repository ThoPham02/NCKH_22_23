const Department = require("../../models/departments.models");
const { handleResponse } = require("../utils");

module.exports = {
  getDepartment: function (req, res) {
    var conditions = req.query;
    Department.GetDepartments(conditions, (err, data) =>
      handleResponse(res, err, data)
    );
  },

  createDepartment: function (req, res) {
    var data = req.body;
    Department.Create(data, (err, data) => handleResponse(res, err, data));
  },

  updateDepartment: function (req, res) {
    var id = req.params.id * 1;
    var body = req.body;
    var data = {
      id,
      ...body,
    };
    Department.Update(data, (err, data) => handleResponse(res, err, data));
  },

  deleteDepartment: function (req, res) {
    var id = req.params.id;
    Department.Delete(id, (err, data) => handleResponse(res, err, data));
  },
};
