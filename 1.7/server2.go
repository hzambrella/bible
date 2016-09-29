package main

import (
	"bible/1.7/datastore"
	"fmt"
	"log"
	"net/http"
)

 var data datastore.Data

func main() {
	var d datastore.Data
	d, err := datastore.ParseDataFromFile("datastore/test")
	if err != nil {
		log.Fatal(err)
	}
	data=d
	fmt.Println("listen and serve: 8080")
	http.HandleFunc("/view", view)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func view(w http.ResponseWriter, r *http.Request) {
	var result []byte
	result = data.ToJson()
	what, err := w.Write(result)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(what)
}
