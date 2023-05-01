const express = require('express')
const mysql = require('mysql2')
const app = express()

const connection = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: 'super-password',
    database: 'testExpress',
});

app.get('/fetch', (req, res) => {
    res.send('fetched')
})