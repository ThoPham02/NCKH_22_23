var express = require('express');
const router = express.Router();

const { getUser, createUser, updateUser, deleteUser } = require('../controllers/admin/users.controller');
router.get('/users', getUser);
router.post('/users', getUser);
router.put('/users/:id', getUser);
router.delete('/users/:id', getUser);



module.exports = router;