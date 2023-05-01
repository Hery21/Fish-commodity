const express = require('express')
const aggregateController = require('./src/controllers/aggregate-controller.js');
const fetchController = require('./src/controllers/fetch-controller.js');
const authorizeJWT = require('./src/middlewares/authorize-jwt.js');
const app = express()

app.get('/aggregate', authorizeJWT, aggregateController);
app.get('/fetch', authorizeJWT, fetchController);

app.listen(3000, () => {
    console.log('Server running on port 3000');
})