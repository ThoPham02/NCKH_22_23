var express = require('express');
var router = express.Router();

router.use('/admin', require('./admin.router.js'))


module.exports = router;