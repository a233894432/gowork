package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"goji.io"
	"goji.io/pat"
)

var db *bolt.DB

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {

	db, err := bolt.Open("/tmp/data.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	http.ListenAndServe("localhost:8000", mux)
}
