var express = require('express');
var router = express.Router();

const test = require('./test.router')
router.use('/test', test)

module.exports = router;