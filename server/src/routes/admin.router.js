var express = require('express');
const { getUser } = require('../controllers/admin/users.controller');
const router = express.Router();
const { testController } = require('../controllers/testController/test.controller');

router.get('/users', getUser);
router.post('/users', testController);
router.put('/users/:id', testController);
router.delete('/users/:id', testController);

module.exports = router;