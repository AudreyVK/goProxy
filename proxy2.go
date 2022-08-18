/*package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	demoURL, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(demoURL)
	http.ListenAndServe(":8081", proxy)
}

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
