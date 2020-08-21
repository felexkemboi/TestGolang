package main

import ( 
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"encode/json"
	"fmt"
	"io/ioutil"
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

	func createPost(w http.ResponseWriter, r *http.Request){
		stmt, err := db.Prepare("INSERT INTO posts(title) VALUES(?)")

		if err != nil{
			panic(err.Error())
		}

		body,err := ioutil.ReadAll(r.Body)
		if err := nil(
			panic(err.Error())
		)

		keyVal := make(map[string]string)
		json.Unmarshal(body,&keyVal )

		_,err = stmt.Exec(title)
		if err != nil{
			panic(err.Error())
		}

		fmt.Fprint(w, "New post Created!")
	}

	func getPost(w http.ResponseWriter, r *http.Resquest){
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		result,err := db.Query("SELECT id,title FROM posts WHERE id = ?", params["id"])
	    if err != nil{
			panic(err.Error())
		}

		defer result.Close()

		var post post

		for result.Next(){
			err != nil{
				panic(err.Error())
			}
		}

		json.NewEncoder(w).NewEncoder(w).Encode(post)
	}

	router.HandleFunc("/posts", getPosts).Methods("GET")
  	router.HandleFunc("/posts", createPost).Methods("POST")
  	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
  	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
  	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}