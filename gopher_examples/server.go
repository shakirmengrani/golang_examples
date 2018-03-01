package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
	"fmt"
)


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		t, _ := template.ParseFiles("index.html")
		var p interface{}
		err := t.Execute(w, p)
		if err != nil{
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":8000", r)
}