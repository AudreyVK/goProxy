package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		req, err := http.NewRequest(r.Method, "http:/"+r.URL.Path, r.Body)
		if err != nil {
			log.Fatalf("Error Occurred. %+v", err)
		}

		reqHeaderClone := r.Header.Clone()
		for key, element := range reqHeaderClone {
			for i := 0; i < len(element); i++ {
				req.Header.Add(key, element[i])
			}
		}
		/*defer r.Body.Close()
		reqbody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("Couldn't parse request body. %+v", err)
		}*/

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalf("Error sending request. %+v", err)
		}

		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Couldn't parse response body. %+v", err)
		}

		headerRes := response.Header.Clone()

		for key, element := range headerRes {
			for i := 0; i < len(element); i++ {
				w.Header().Add(key, element[i])
			}
		}

		w.WriteHeader(response.StatusCode)

		fmt.Fprintf(w, "%v", string(body))

		//fmt.Printf("%v", string(reqbody))
		//fmt.Print("http:/" + r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
