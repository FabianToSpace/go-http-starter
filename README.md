[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/license/mit)

# Go-HTTP-Starter

This is a simple starter setup for a Golang HTTP Server application. Besides a sample setup for a HTTP Server, this starter enables the use of PostgreSQL with migrations too.


## Getting started
Run the Server using ```go run .``` or build the server using ```go build```. Besides that, you can build a docker container using the Dockerfile in this repository.

## Add HTTP Endpoints
To add HTTP Endpoints to the server, you can register a new HTTP handler in the ```registerRoutes()``` function. Take the ```healthHandler``` as a starting point. 

## Add migrations
To add Database Migrations, you have to add a migration SQL file, that follows the scheme ```<number>_description.up.sql```.\
Migrations are automatically applied to the Database Server on the server startup.