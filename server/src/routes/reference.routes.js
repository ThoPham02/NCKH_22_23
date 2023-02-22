var express = require('express');
const router = express.Router();
const referenceRoutes = require('../controllers/reference.controller.js')

router.get('/', referenceRoutes.get_list );

module.exports = router;