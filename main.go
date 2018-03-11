package main

import (
	"encoding/json"
	"fmt"
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
				panic(err)
			}

			val, ok := db[req.Key]
			if !ok {
				panic(fmt.Sprintf("Key '%s' not in DB", req.Key))
			}

			resp, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}

			_, err = w.Write(resp)
			if err != nil {
				panic(err)
			}

		case "POST":
			dec := json.NewDecoder(r.Body)
			req := Request{}
			err := dec.Decode(&req)
			if err != nil {
				panic(err)
			}

			db[req.Key] = req.Value

		default:
			panic(fmt.Sprintf("HTTP method %s not supported", r.Method))
		}
	})

	http.ListenAndServe(":8080", nil)
}
