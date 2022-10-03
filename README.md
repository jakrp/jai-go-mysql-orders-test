# jai-go-mysql-orders-test

This is for Zerologix coding test

You need to have Go 1.16+ version
Also a MySQL 5.1+ server 

Clone the project

Inside the project you will find these directories - cmd, pkg and a compressed file mysql-db.zip

The compressed file mysql-db.zip must be extracted. It has a database file of .sql format. Add this database to mysql server. 

Inside the pkg directory, you have routes, controller, model and config. Edit the config/app.go file and add the mysql database configuration details inside Connect() function. 

This is a REST-API based application. The pkg/routes/order-routes.go file has the API endpoints.

To start the project 
> cd cmd/main

After inside the directory cmd/main, run this command
> go run main.go
