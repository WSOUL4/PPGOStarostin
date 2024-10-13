package main

import (
	//"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	//"os"

	"github.com/golang-jwt/jwt/v4"

	//"encoding/json"
	"strings"
	//"net/url"
	"time"
)

var sampleSecretKey = []byte("Abrakadabra")

func generateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
func main() {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	url := "http://localhost:8080/login"
	method := "POST"

	payload := strings.NewReader(`{
    "username": "Admin",
    "password": "Admin"
}`)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	Token := string(body)

	d1, err := client.Get("http://localhost:8080/users")
	if err != nil {
		log.Fatal(err)
	}
	defer d1.Body.Close()
	fmt.Println()
	if d1.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(d1.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		//log.Fatal().Info(bodyString)
	}
	//d1, _ = client.PostForm("http://localhost:8080/users", url.Values{"name": {"posted"}})
	//d1, _ = client.PostForm("http://localhost:8080/user", url.Values{"name": {"chaned"}, "id": {"1"}})
	//defer d1.Body.Close()

	ur := "http://localhost:8080/user"
	data := []byte(`{"name":"barisrt", "id": 4}`)
	req, err = http.NewRequest(http.MethodPut, ur, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data = []byte(`{ "id": 1}`)
	req, err = http.NewRequest(http.MethodDelete, ur, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data = []byte(`{ "id": 2}`)
	req, err = http.NewRequest(http.MethodGet, ur, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	ur = "http://localhost:8080/users/id"
	data = []byte(`{"name":"barisrt", "id": 4}`)
	req, err = http.NewRequest(http.MethodGet, ur, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	//log.Fatal().Info(bodyString)

	ur = "http://localhost:8080/users/name"
	data = []byte(`{"name":"posted", "id": 4}`)
	req, err = http.NewRequest(http.MethodGet, ur, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString = string(bodyBytes)
	fmt.Println(bodyString)
	//log.Fatal().Info(bodyString)

	ur = "http://localhost:8080/users/name"
	data = []byte(`{"name":"poste$d", "id": 4}`)
	req, err = http.NewRequest(http.MethodGet, ur, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString = string(bodyBytes)
	fmt.Println(bodyString)
	//log.Fatal().Info(bodyString)

	defer resp.Body.Close()

	//d1, _=client.Do("http://localhost:8080/users", url.Values{"name": {"New Insert234234", "id": {3}}})
	//r := bytes.NewReader(data)
	//d1, _ = client.PostForm("http://localhost:8080/data", url.Values{"name": {"Value"}, "occupation": {"Value"}})
	//fmt.Println(d1)
	//http.PostForm("http://example.com/form",
	//url.Values{"key": {"Value"}, "id": {"123"}})
}
