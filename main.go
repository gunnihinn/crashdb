package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := make(map[string]interface{})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	type Request struct {
		Key   string
		Value interface{}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			dec := json.NewDecoder(r.Body)
			req := Request{}
			err := dec.Decode(&req)
			if err != nil {
				panic("User error")
			}

			val, ok := db[req.Key]
			if !ok {
				panic("User error")
			}

			resp, err := json.Marshal(val)
			if err != nil {
				panic("User error")
			}

			_, err = w.Write(resp)
			if err != nil {
				panic("User error")
			}

		case "POST":
			dec := json.NewDecoder(r.Body)
			req := Request{}
			err := dec.Decode(&req)
			if err != nil {
				panic("User error")
			}

			db[req.Key] = req.Value

		default:
			panic("User error")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
