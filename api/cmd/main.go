package main

import (
	"log"
	"myExpenses/config"
	"myExpenses/database"
	"net/http"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	r := config.Router(*database.DB)
	err := http.ListenAndServe(":8880", r)
	if err != nil {
		log.Fatal("failed to serve: ", err)
	}

}
