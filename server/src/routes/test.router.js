var express = require('express');
const router = express.Router();
const { testController, test2Controller } = require('../controllers/testController/test.controller');

router.get('/1', testController);
router.get('/2', test2Controller);

module.exports = router;