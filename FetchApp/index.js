const express = require('express')
const app = express()

app.get('/fetch', (req, res) => {
    res.send('fetched')
})