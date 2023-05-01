# Fish Commodity Project

## Prerequisites

- Go
- Docker
- MySQL
- Node.js
- Express.js

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/your-repo.git

2. Set up the MySQL database by running the db.sql script:
   mysql -u <username> -p < db.sql

3. Install Node.js dependencies:
   npm install
   
4. Build the Docker image:
   docker build -t FISH-COMMODITY .
   
5. Run the Docker container
   docker run -p 8080:8080 -p 3000:3000 -e MYSQL_HOST=<mysql_host> -e MYSQL_USER=<mysql_user> -e MYSQL_PASSWORD=<mysql_password> my-go-project

6. Open a web browser and go to http://localhost:3000 to see the Express.js application, and go to http://localhost:8080 to see the Go application.

## Usage
The Go application is a simple web server that listens on port 8080. You can access it by going to http://localhost:8080 in a web browser.

The Express.js application is a simple web server that listens on port 3000. You can access it by going to http://localhost:3000 in a web browser.

The API endpoints for the Go application are:

POST /register - receives json with new user info, create the user in DB, and generate password.
POST /login - Outputs JWT Token if the credentials given is true.
GET /validate - validate the token and outputs the decoded token.

The API endpoints for the Express.js application are:

GET /fetch - get the data from the link and create a new column by converting price to USD and pass it to the DB.
GET /aggregate - role admin can aggregate the data from previously fetched data.

## Docker Configuration
The Docker image is configured to expose ports 8080 and 3000, and set the environment variables APP_ENV, MYSQL_HOST, MYSQL_USER, and MYSQL_PASSWORD.
