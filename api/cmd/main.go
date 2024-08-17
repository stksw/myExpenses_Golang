package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	println("hello world")
	dsn := fmt.Sprintf(
		"host=postgres user=postgres password=pass dbname=myExpenses port=5432 sslmode=disable TimeZone=Asia/Tokyo")
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to init database: ", err)
	}
	log.Default().Println("success to connect db !!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	err = http.ListenAndServe(":8880", nil)
	if err != nil {
		log.Fatal("failed to serve: ", err)
	}
	log.Default().Println("Server started on port: 8880")
}
