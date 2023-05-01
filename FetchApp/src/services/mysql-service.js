const mysql = require('mysql2');
const config = require('../config.js');

const connection = mysql.createConnection(config.database);

module.exports = connection;