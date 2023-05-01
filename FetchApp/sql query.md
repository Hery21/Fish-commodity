CREATE TABLE prices (
  id INT NOT NULL AUTO_INCREMENT,
  uuid VARCHAR(255) NOT NULL,
  komoditas VARCHAR(255),
  area_provinsi VARCHAR(255),
  size VARCHAR(255),
  price INT NOT NULL,
  price_usd DECIMAL(10, 2),
  tgl_parsed TIMESTAMP,
  PRIMARY KEY (id)
);