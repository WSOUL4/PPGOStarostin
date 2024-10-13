package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func main() {
	call("http://localhost:8080/data", "POST")
}
func call(urlPath, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormField("name")
	if err != nil {
	}
	_, err = io.Copy(fw, strings.NewReader("John"))
	if err != nil {
		return err
	}
	fw, err = writer.CreateFormField("occupation")
	if err != nil {
	}
	_, err = io.Copy(fw, strings.NewReader("23"))
	if err != nil {
		return err
	}

	// Close multipart writer.
	writer.Close()
	req, err := http.NewRequest("POST", "http://localhost:8080/data", bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	return nil
}
