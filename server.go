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

		//fmt.Printf("REQUEST:\n%s", string(reqDump))
		w.Write([]byte("Request: \n"))
		w.Write([]byte(reqDump))
	})
	log.Printf("Starting HTTP server at port: %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

/*func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("http://www.google.com")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			body, _ := ioutil.ReadAll(response.Body)
			fmt.Println("Statuscode: ", response.StatusCode)
			fmt.Println("Statustext: ", http.StatusText(response.StatusCode))
			fmt.Println(string(body))
		}
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}*/

/*const port = 8082

func indexHandler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s", string(reqDump))
	w.Write([]byte("Hello World"))
}*/
