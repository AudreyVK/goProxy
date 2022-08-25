package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Flusher interface {
	// Flush sends any buffered data to the client.
	Flush()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(r.Method, "http:/"+r.RequestURI, nil)

		if err != nil {
			log.Fatalf("Error Occurred. %+v", err)
		}

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalf("Error sending request to API endpoint. %+v", err)
		}

		defer response.Body.Close()

		/*body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Couldn't parse response body. %+v", err)
		}*/
		//fmt.Fprintf(w, "%s %q", r.Method, html.EscapeString(r.URL.Path)) debugging

		//fmt.Println(r.Body.Read([]byte(response.Status)))

		//fmt.Fprintf(w, "%v\n", headerRes)
		fmt.Fprintf(w, "header: %v\n", response.Header)
		fmt.Fprintf(w, "%v\n", response.StatusCode)
		fmt.Fprintf(w, "%v\n", http.StatusText(response.StatusCode))
		fmt.Fprintf(w, "%v / %v\n", r.Method, response.Request.Proto)
		fmt.Fprintf(w, "Host: %v\n", r.Host)

		//w.Header().Add("statuscode:", r.RemoteAddr)
		//fmt.Fprintf(w, "%v\n", body)
		/*for key, element := range r.Header {
			element := strings.Replace(element[0], "[", "", -1)
			fmt.Fprintf(w, "%v: %v\n", key, element)
			if err != nil {
				log.Fatal(err)
			}
		}*/
		headerRes := response.Header.Clone()
		for key, element := range headerRes {
			element := strings.Replace(element[0], "[", "", -1)
			fmt.Fprintf(w, "%v: %v\n", key, element)
			w.Header().Add(key, element)
			if err != nil {
				log.Fatal(err)
			}
		}

	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

//fmt.Println(http.StatusText(response.StatusCode))
//fmt.Println(string(body))

/*func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}*/
/*var website string
fmt.Println("Type in the URL you want to reach.")
fmt.Scan(&website)
fmt.Println()*/
//response := sendRequest(http.MethodPost)
//log.Println("Response Body:", string(response))
//response, err := sendRequest(http.MethodPost)
//if err != nil {
//	fmt.Println(err.Error())
//} else {
//defer response.Body.Close()
//body, _ := ioutil.ReadAll(response.Body)
//fmt.Println("Statuscode: ", response.StatusCode)
//fmt.Println("Statustext: ", http.StatusText(response.StatusCode))
//fmt.Println(string(response))
//}

/*func sendRequest(method string) ([]byte, error) {
	server := "http://:8082"
	values := map[string]string{"foo": "bar"}
	jsonData, err := json.Marshal(values)

	req, err := http.NewRequest(method, server, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	return body, err
}*/

/*type Pxy struct{}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport
	// step 1
	outReq := new(http.Request)
	*outReq = *req // this only does shallow copies of maps
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}
	// step 2
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	// step 3
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	fmt.Println("Statuscode: ", res.StatusCode)
	fmt.Println("Statustext: ", http.StatusText(res.StatusCode))
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)

	res.Body.Close()
}

func main() {
	fmt.Println("Serve on :8080")
	http.Handle("/", &Pxy{})
	http.ListenAndServe(":8080", nil)
}*/

/*func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Print(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}

		go handleConn(conn)
	}

}

func handleConn(src net.Conn) {
	dst, err := net.Dial("tcp", "www.th-luebeck.de:80")
	if err != nil {
		log.Print(err)
	}

	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Print(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Print(err)
	}
}*/

/*func main() {
	/*conn, err := net.Dial("tcp", "http://172.17.0.2:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	http.ListenAndServe(":8080", nil)
	conn, err := net.Dial("tcp", "http://172.17.0.2:80")

	proxyUrl, err := url.Parse("http://167.235.57.197:8080")
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	response, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Statuscode: ", response.StatusCode)
		fmt.Println("Statustext: ", http.StatusText(response.StatusCode))
		fmt.Println(body)
	}
}*/
//response, err := http.Get("http://www.google.com")
/*if err != nil {
	fmt.Println(err.Error())
} else {
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Statuscode: ", response.StatusCode)
	fmt.Println("Statustext: ", http.StatusText(response.StatusCode))
	fmt.Println(body)
}*/
