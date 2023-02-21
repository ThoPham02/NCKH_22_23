const db = require('../config.js')

const Reference = function(us){
    this.id = us.id;
    this.username = us.username;
    this.password = us.password;
    this.user_type_id = us.user_type_id;
    this.created_at = us.created_at;
    this.update_at = us.update_at
    this.deleted_at = us.deleted_at;
}

 Reference.get_all = function(result) {
    db.query("SELECT * FROM reference",function(err,rf){
        if(err){
            result(err);
            return;
        }
        else {
            result(rf);
        }
    })
}

module.exports = Reference;