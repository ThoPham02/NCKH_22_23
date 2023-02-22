var express = require('express');
const router = express.Router();

const { getUser, createUser, updateUser, deleteUser } = require('../controllers/admin/users.controller');
router.get('/users', getUser);

router.post('/createUsers', createUser);

router.put('/update/:id', updateUser);

router.delete('/delete/:id', deleteUser);



module.exports = router;