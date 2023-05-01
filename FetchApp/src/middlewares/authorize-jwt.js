const jwt = require('jsonwebtoken');
const config = require('../config')

function authorizeJWT(req, res, next) {
  const authHeader = req.headers.authorization;
  const token = authHeader && authHeader.split(' ')[1];

  if (token == null) {
      return res.status(401).json({ message: 'Unauthorized' });
  }

//   console.log('===================', token, config.jwt.JWTSecret)

  jwt.verify(token, config.jwt.JWTSecret, function(err, decoded) {
      if (err) {
          return res.status(401).json({ message: 'Unauthorized' });
      }
    //   console.log("passssssss")

      req.decoded = decoded;

      next();
  });
}

module.exports = authorizeJWT;
