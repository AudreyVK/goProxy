package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    proxyUrl, err := url.Parse("http://167.235.57.197:8080")
    http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
    response, err := http.Get("http://www.google.com")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        body, _ := ioutil.ReadAll(response.Body)
        fmt.Println("Statuscode: ", response.StatusCode)
        fmt.Println("Statustext: ", http.StatusText(response.StatusCode))
        fmt.Println("Body: ", body)
    }
}	