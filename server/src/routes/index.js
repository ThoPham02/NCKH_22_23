var express = require('express');
var router = express.Router();

const test = require('./test.router')
const reference = require('./reference.routes.js')
router.use('/test', test)
router.use('/showListReference',reference)
module.exports = router;