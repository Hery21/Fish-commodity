const https = require('https');
const mysqlService = require('../services/mysql-service');
const exchangeService = require('../services/exchange-service.js');

function fetchController(req, res) {
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
                  exchangeService.convertToUSD(jsonData[index].price, function(priceUsd) {
                    mysqlService.execute(
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
                              mysqlService.close();
                              res.status(200).send('Data inserted successfully');
                          }
                        }
                      }
                    );
                  });
                } else {
                  console.log('Data inserted successfully');
                  mysqlService.close();
                }
              }
          
              insertNextRow();
            });
      }).on('error', (error) => {
          console.error(error)
          res.status(500).send('Error fetching data')
      })
}

module.exports = fetchController;
