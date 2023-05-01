const mysqlService = require('../services/mysql-service.js');

function aggregateController(req, res) {
  const { role } = req.decoded['user'];

    if (role !== 'admin') {
        return res.status(401).json({ message: 'Unauthorized' });
    }

    mysqlService.connect(function(err) {
        if (err) {
          console.error('Error connecting to database:', err);
          res.status(500).json({ message: 'Internal Server Error' });
          return;
        }
    
        const query = `
          SELECT
            area_provinsi,
            YEARWEEK(tgl_parsed) AS week,
            MIN(price) AS min_price,
            MAX(price) AS max_price,
            AVG(price) AS avg_price,
            MIN(size) AS min_size,
            MAX(size) AS max_size,
            AVG(size) AS avg_size
          FROM prices
          GROUP BY area_provinsi, week
          ORDER BY area_provinsi, week;
        `;
    
        mysqlService.query(query, function(error, results, fields) {
          if (error) {
            console.error('Error executing query:', error);
            res.status(500).json({ message: 'Internal Server Error' });
            return;
          }
    
          res.json(results);
        });
    
        mysqlService.end();
    });
}

module.exports = aggregateController;
