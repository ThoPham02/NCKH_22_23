const Faculties = require("../../models/faculties.models");
const { handleResponse } = require("../utils");

module.exports = {
  getFaculties: function (req, res) {
    const conditions = {};
    Faculties.GetFaculties(conditions, (err, data) =>
      handleResponse(res, err, data)
    );
  },

  createFaculties: function (req, res) {
    var data = req.body;
    Faculties.Create(data, (err, data) => handleResponse(res, err, data));
  },

  updateFaculties: function (req, res) {
    var id = req.params.id * 1;
    var body = req.body;
    var data = {
      id,
      ...body,
    };
    Faculties.Update(data, (err, data) => handleResponse(res, err, data));
  },

  deleteFaculties: function (req, res) {
    var id = req.params.id;
    Faculties.Delete(id, (err, data) => handleResponse(res, err, data));
  },
};
