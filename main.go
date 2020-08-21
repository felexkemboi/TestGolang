package main

import ( 
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"io/ioutil"
	    "net/http"

)


var db *sql.DB
var err error


type Post struct {
	ID string `json:"id"`
	Title string `json:"title"`
}

func main(){
	db,err = sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/Go1")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router :=mux.NewRouter()


	router.HandleFunc("/posts", getPosts).Methods("GET")
  	router.HandleFunc("/posts", createPost).Methods("POST")
  	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
  	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
  	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func createPost(w http.ResponseWriter, r *http.Request){
	stmt, err := db.Prepare("INSERT INTO posts(title) VALUES(?)")

	if err != nil{
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	   panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body,&keyVal )
	title := keyVal["title"]

	_,err = stmt.Exec(title)
	if err != nil{
		panic(err.Error())
	}

	fmt.Fprint(w, "New post Created!")
}

func getPost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result,err := db.Query("SELECT id,title FROM posts WHERE id = ?", params["id"])
	if err != nil{
		panic(err.Error())
		}

	defer result.Close()

	var post Post

	for result.Next() {
    	err := result.Scan(&post.ID, &post.Title)
    	if err != nil {
      		panic(err.Error())
    	}
  	}

	json.NewEncoder(w).Encode(post)
}

func getPosts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var posts []Post

	result ,err := db.Query("SELECT id,title FROM posts")
	if err != nil{
		panic(err.Error())
	}

	defer result.Close()

	for result.Next(){
		var post Post
		err :=result.Scan(&post.ID, &post.Title)
		if err != nil{
			panic(err.Error())
		}

	posts = append(posts,post)

	}
  	json.NewEncoder(w).Encode(posts)
}


func updatePost(w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)

	stmt ,err := db.Prepare("UPDATE posts SET title = ? WHERE id = ?" )
	if err != nil{
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil{
		panic(err.Error())
	}

	keyVal :=make(map[string] string)
	json.Unmarshal(body,&keyVal)
	newTitle := keyVal["title"]

	_,err = stmt.Exec(newTitle,params["id"])

	if err != nil{
		panic(err.Error())
	}

	fmt.Fprint(w, "Post with ID = %s was updated", params["id"])
}

func deletePost(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)

	stmt,err := db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil{
		panic(err.Error())
	}

	_,err = stmt.Exec(params["id"])

	if err != nil{
		panic(err.Error())
	}	
	fmt.Fprint(w, "Post with ID = %s was deleted", params["id"])	
}