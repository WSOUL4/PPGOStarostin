package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	d1, _ := client.Get("http://localhost:8080/hello")
	fmt.Println(d1)
	//data := []byte(`{"name":"bar", "occupation":"bar"}`)
	//r := bytes.NewReader(data)
	d1, _ = client.PostForm("http://localhost:8080/data", url.Values{"name": {"Value"}, "occupation": {"Value"}})
	fmt.Println(d1)
	//http.PostForm("http://example.com/form",
	//url.Values{"key": {"Value"}, "id": {"123"}})
}
