package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func setheaders(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		
		w.Header().Set("Access-Control-Allow-Origin","*")
    	w.Header().Set("Content-Type","application/json")
    	w.Header().Set("Access-Control-Allow-Methods","POST,GET")
    	w.Header().Set("Access-Control-Allow-Headers","Accept, Content-Type, Content-Length, Accept-Encoding")	
	
		h(w,r)
	}	
}
func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method:",r.Method)
		next(w,r)
	}
}

func details(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	fmt.Println("Email:",email)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Methods","POST,GET")
	w.Header().Set("Access-Control-Allow-Headers","Accept, Content-Type, Content-Length, Accept-Encoding")

	response, err := json.Marshal(&struct{
		FirstName string
		LastName string
		Email string
		DateOfBirth string
	}{
		FirstName: "Adriel",
		LastName: "Artiza",
		Email: "adriel.artiza@gmail.com",
		DateOfBirth:"2019-08-10",
	})

	if err != nil {
		panic(err)
	}

	fmt.Fprint(w,string(response))
}

func main() {
	http.HandleFunc("/user",logger(home))
	http.HandleFunc("/details",setheaders(details))
	if err := http.ListenAndServe(":3000",nil); err != nil {
		panic(err)
	}
}
