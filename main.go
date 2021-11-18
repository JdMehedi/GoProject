package main

import (
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"main.go/handler"
)

func main(){

	var schema = `
	CREATE TABLE IF NOT EXISTS categories (
		id serial,
		name text,

		primary key(id)
	);`

	db, err := sqlx.Connect("postgres", "user=postgres password=Passw0rd dbname=test sslmode=disable")
    if err != nil {
        log.Fatalln(err)
    }

	db.MustExec(schema)

	h:=handler.New(db)

 http.HandleFunc("/", h.Home)
 http.HandleFunc("/categories/create", h.CreateCategory)
 http.HandleFunc("/categories/store", h.StoreCategory)
 http.HandleFunc("/categories/edit/", h.EditCategory)
 http.HandleFunc("/categories/update/", h.UpdateCategory)
 http.HandleFunc("/categories/delete/", h.DeleteCategory)
 log.Println("Server starting ...........")

 if err := http.ListenAndServe("127.0.0.1:3000",nil); err !=nil{
	log.Fatal(err)
  }

}