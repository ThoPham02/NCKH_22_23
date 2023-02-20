const db = require('../config.js')

const Reference = function(rf){
    this.id = rf.id;
    this.reference_url = rf.id;
    this.name = rf.name;
    this.created_at = rf.reated_at;
    this.update_at = rf.update_at;
    this.deleted_at = rf.deleted_at;
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

module.exposts = Reference;