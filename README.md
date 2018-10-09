# golang
Golang webserver CRUD using Iris and mySQL

Installation:

1. Download golang and mysql on the machine.
2. Install iris and mysql packages
  $ go get -u github.com/go-sql-driver/mysql
  $ go get -u github.com/kataras/iris
3. Create database using
  mysql> create database dbname;
4. Create table using
  mysql> use dbname;
  mysql> create table user(
   Name VARCHAR(100) NOT NULL,
   Description VARCHAR(100) NOT NULL,
  );
5. Now run the server
  $ go run Webserver.go
