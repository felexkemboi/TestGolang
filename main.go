package main

import ( 
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main(){
	db,err = sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/Go1")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router :=mux NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
  	router.HandleFunc("/posts", createPost).Methods("POST")
  	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
  	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
  	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
  	
	http.ListenAndServe(":8000", router)
}