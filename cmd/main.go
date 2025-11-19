package main

import (
	"ecom/cmd/api"
	"ecom/config"
	"ecom/db"
	"log"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Net: "tcp",
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		ParseTime: true,
		AllowNativePasswords: true,
	})

	if err != nil {
		log.Fatal(err)
	}
	
	initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Connected to database")
}