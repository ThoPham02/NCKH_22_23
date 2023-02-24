const Users = require("../../models/users.models");
const { handleResponse } = require("../utils");

module.exports = {
  getUser: function (req, res) {
    var user_type_id = req.query.user_type_id;
    const conditions = {
      user_type_id: user_type_id,
    };
    Users.GetUsers(conditions, (err, data) => handleResponse(res, err, data))
  },

  createUser: function (req, res) {
    var data = req.body;
    Users.Create(data, (err, data) => handleResponse(res, err, data));
  },

  updateUser: function (req, res) {
    var id = req.params.id * 1;
    var body = req.body;
    var data = {
      id,
      ...body,
    };
    Users.Update(data, (err, data) => handleResponse(res, err, data));
  },

  deleteUser: function (req, res) {
    var id = req.params.id;
    Users.Delete(id, (err, data) => handleResponse(res, err, data));
  },
};
