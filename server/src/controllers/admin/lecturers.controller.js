const Lecturers = require("../../models/lecturers.models.js");
const { handleResponse } = require("../utils");

module.exports = {
  getLecturers: function (req, res) {
    const conditions = req.query
    Lecturers.GetLectures(conditions, (err, data) => handleResponse(res, err, data))
  },

  createLecturers: function (req, res) {
    var data = req.body;
    Lecturers.Create(data, (err, data) => handleResponse(res, err, data));
  },

  updateLecturers: function (req, res) {
    var id = req.params.id * 1;
    var body = req.body;
    var data = {
      id,
      ...body,
    };
    Lecturers.Update(data, (err, data) => handleResponse(res, err, data));
  },

  deleteLecturers: function (req, res) {
    var id = req.params.id;
    Lecturers.Delete(id, (err, data) => handleResponse(res, err, data));
  },
};
