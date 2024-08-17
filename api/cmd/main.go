package main

import (
	"log"
	"myExpenses/database"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	err := http.ListenAndServe(":8880", nil)
	if err != nil {
		log.Fatal("failed to serve: ", err)
	}
	log.Default().Println("Server started on port: 8880")
}
