var express = require('express');
const router = express.Router();
const { testController } = require('../controllers/testController/test.controller');

router.get('/users', testController);
router.post('/users', testController);
router.put('/users/:id', testController);
router.delete('/users/:id', testController);

module.exports = router;