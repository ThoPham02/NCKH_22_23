const Reference = require('../models/reference/reference.models.js')
console.log(Reference)
exports.get_list = function(req, res) {
    Reference.get_all(function(rf){
        res.send({result:rf});
    })
}
