const express = require('express');
const morgan = require('morgan');
const { default: helmet } = require('helmet');
const compression = require('compression');
const app = express();

//add middleware
app.use(morgan("dev"))
app.use(helmet())
app.use(compression())

//add db
const db = require('./models/config');

//add routes
app.get('/', (req, res, next) => {
    res.status(200).json({
        message: "Test app"
    })
})

//add handle error

module.exports = app;