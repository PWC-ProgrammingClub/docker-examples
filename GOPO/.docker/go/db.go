package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var Connect *connect

type connect struct {
	user     string
	password string
	host     string
	db       string
	port     string
}

// Initialize the global connection object with the environment variables; validate they exist
func ConnectInit() {
	c := connect{}
	c.user = os.Getenv("DB_USER")
	strNotEmpty(c.user, "DB_USER not set!")
	c.password = os.Getenv("DB_PASSWORD")
	strNotEmpty(c.password, "DB_PASSWORD not set!")
	c.host = os.Getenv("DB_HOST")
	strNotEmpty(c.host, "DB_HOST not set!")
	c.db = os.Getenv("DB_DB")
	strNotEmpty(c.db, "DB_DB not set!")
	c.port = os.Getenv("DB_PORT")
	if len(c.port) > 0 {
		c.port = ":" + c.port
	}
	Connect = &c
}

// Get the connection string (private method of connection object)
func (c *connect) getStr() string {
	return "postgres://" + c.user + ":" + c.password + "@" + c.host + c.port + "/" + c.db
}

// Get a database connection (method of connection object)
func (c *connect) DB() *sql.DB {
	db, err := sql.Open("pgx", c.getStr())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
