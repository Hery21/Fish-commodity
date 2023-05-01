const express = require('express')
const mysql = require('mysql2')
const app = express()

const connection = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: 'super-password',
    database: 'fish_commodity',
});

function authorizeJWT(req, res, next) {
    const authHeader = req.headers.authorization;
    const token = authHeader && authHeader.split(' ')[1];
  
    if (token == null) {
        return res.status(401).json({ message: 'Unauthorized' });
    }
  
    jwt.verify(token, secretKey, function(err, decoded) {
        if (err) {
            return res.status(401).json({ message: 'Unauthorized' });
        }
  
        req.decoded = decoded;
  
        next();
    });
}

app.get('/fetch', authorizeJWT, (req, res) => {
    https.get('https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list', (response) => {
        let data = ''

        response.on('data', (chunk)=> {
            data += chunk
        })

        response.on('end', function() {
            const jsonData = JSON.parse(data);
        
            let index = 0;
        
            function insertNextRow() {
              if (index < jsonData.length) {
                convertToUSD(jsonData[index].price, function(priceUsd) {
                  connection.execute(
                    'INSERT INTO prices (uuid, komoditas, area_provinsi, size, price, price_usd, tgl_parsed) VALUES (?, ?, ?, ?, ?, ?, ?)',
                    [jsonData[index].uuid, jsonData[index].komoditas, jsonData[index].area_provinsi, jsonData[index].size, jsonData[index].price, priceUsd, new Date(jsonData[index].tgl_parsed)],
                    function(error) {
                      if (error) {
                        console.error(error);
                      } else {
                        index++;
                        insertNextRow();
                        if (index === jsonData.length) {
                            console.log('Data inserted successfully');
                            connection.close();
                            res.status(200).send('Data inserted successfully');
                        }
                      }
                    }
                  );
                });
              } else {
                console.log('Data inserted successfully');
                connection.close();
              }
            }
        
            insertNextRow();
          });
    }).on('error', (error) => {
        console.error(error)
        res.status(500).send('Error fetching data')
    })
})