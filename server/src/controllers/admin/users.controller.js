const Users = require('../../models/users/users.models')
// const express = require('express')
// const app = express()
// const bodyParser = require('body-parser');
// app.use(bodyParser.urlencoded({extended: false}));
module.exports = {
    getUser: function(req, res) {
        const conditions = {}
        Users.GetUsers(conditions, function(result) {
            if (result.status === 400) {
                res.status(400).json({
                    message: "Bad request"
                })
                return
            }
            res.status(200).json({
                message: "OK",
                data: result.data
            })
        })
    },
    createUser: function (req, res) {
        var data = req.body;
        Users.create(data, function(res2){         
        res.send({result: res2})
    })
    },
    updateUser: function (req, res) {
        var id = req.params.id * 1;   
        var body = req.body; 
        var data = {
            id, ...body
        }
        Users.update(data, function(res2){
        res.send({result: res2})
    })
    },
    deleteUser: function (req, res) {
        var id = req.params.id;   
        Users.remove(id, function(res2){
        res.send({result: res2})
    })
    }
}