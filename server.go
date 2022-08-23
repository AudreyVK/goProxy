package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	const port = 8082
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte("Request: \n"))
		w.Write([]byte("The Dummy-Statuscode is: 200\n"))
		w.Write([]byte(reqDump))
	})
	log.Printf("Starting HTTP server at port: %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
