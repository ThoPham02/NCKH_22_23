const express = require('express');
const morgan = require('morgan');
const { default: helmet } = require('helmet');
const compression = require('compression');
const app = express();
const bodyParser = require('body-parser')

//add middleware
app.use(bodyParser.urlencoded({extended: false}));
app.use(bodyParser.json());
app.use(morgan("dev"))
app.use(helmet())
app.use(compression())

//add db

//add routes
const router = require('./routes');
app.use(router)

//add handle error

module.exports = app;