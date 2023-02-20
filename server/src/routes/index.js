var express = require('express');
var router = express.Router();

router.use('/showListReference', require('./reference.routes.js'))


module.exports = router;