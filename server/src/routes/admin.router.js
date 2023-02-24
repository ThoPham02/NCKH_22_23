var express = require('express');
const router = express.Router();

const { getUser, createUser, updateUser, deleteUser } = require('../controllers/admin/users.controller');
router.get('/user', getUser);
router.post('/user', createUser);
router.put('/user/:id', updateUser);
router.delete('/user/:id', deleteUser);

const { getFaculties, createFaculties, updateFaculties, deleteFaculties } = require('../controllers/admin/faculties.controller');
router.get('/faculty', getFaculties);
router.post('/faculty', createFaculties);
router.put('/faculty/:id', updateFaculties);
router.delete('/faculty/:id', deleteFaculties);


const { getDepartment, createDepartment, updateDepartment, deleteDepartment } = require('../controllers/admin/departments.controller');
router.get('/department', getDepartment);
router.post('/department', createDepartment);
router.put('/department/:id', updateDepartment);
router.delete('/department/:id', deleteDepartment);

// const { getStudents, createStudents, updateStudents, deleteStudents } = require('../controllers/admin/students.controller');
router.get('/student', getFaculties);
router.post('/student', createFaculties);
router.put('/student/:id', updateFaculties);
router.delete('/student/:id', deleteFaculties);

// const { getLectures, createLectures, updateLectures, deleteLectures } = require('../controllers/admin/lectures.controller');
router.get('/lecture', getFaculties);
router.post('/lecture', createFaculties);
router.put('/lecture/:id', updateFaculties);
router.delete('/lecture/:id', deleteFaculties);

// const { getEvents, createEvents, updateEvents, deleteEvents } = require('../controllers/admin/events.controller');
router.get('/events', getFaculties);
router.post('/events', createFaculties);
router.put('/events/:id', updateFaculties);
router.delete('/events/:id', deleteFaculties);

// const { getGroup, createGroup, updateGroup, deleteGroup } = require('../controllers/admin/groups.controller');
router.get('/group', getFaculties);
router.post('/group', createFaculties);
router.put('/group/:id', updateFaculties);
router.delete('/group/:id', deleteFaculties);

// const { getStudentGruop, createStudentGruop, updateStudentGruop, deleteStudentGruop } = require('../controllers/admin/student_groups.controller');
router.get('/student_group', getFaculties);
router.post('/student_group', createFaculties);
router.put('/student_group/:id', updateFaculties);
router.delete('/student_group/:id', deleteFaculties);

// const { getTopics, createTopics, updateTopics, deleteTopics } = require('../controllers/admin/topics.controller');
router.get('/topic', getFaculties);
router.post('/topic', createFaculties);
router.put('/topic/:id', updateFaculties);
router.delete('/topic/:id', deleteFaculties);

// const { getReferences, createReferences, updateReferences, deleteReferences } = require('../controllers/admin/references.controller');
router.get('/reference', getFaculties);
router.post('/reference', createFaculties);
router.put('/reference/:id', updateFaculties);
router.delete('/reference/:id', deleteFaculties);

module.exports = router;