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

		//r.Header.Add("test", "123")
		w.Header().Add("test", "123")
		//w.Header().Values("test")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "header: %v\n", r.Header)
		/*for key, element := range r.Header {
			element := strings.Replace(element[0], "[", "", -1)
			fmt.Fprintf(w, "%v: %v\n", key, element)
			if err != nil {
				log.Fatal(err)
			}
		}*/

		w.Write([]byte(reqDump))

	})
	log.Printf("Starting HTTP server at port: %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
