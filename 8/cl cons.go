package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"

	//"encoding/json"
	"strings"
	//"net/url"
	"time"
)

var path_users = "http://localhost:8080/users"
var path_user = "http://localhost:8080/user"
var path_user_id = "http://localhost:8080/users/id"
var path_user_name = "http://localhost:8080/users/name"
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
func ReadUser() (str string) {
	/*
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		str = string(stdin)
		str = strings.TrimSuffix(str, "\n")

	*/
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	fmt.Printf("\n--\t%s\t--\n", in.Text())
	str = fmt.Sprintf("%s", in.Text())
	return str
	//line = strings.TrimSuffix(line, "\n")
	//fmt.Println(line)
	//fmt.Scanln(&str)
	//fmt.Println("ReadUser:", str)
	//return line
}
func V1(Token string) {

	//id_int,_:= strconv.Atoi(id)
	data := []byte(`{"name":"0", "id":0}`)
	req, err := http.NewRequest(http.MethodGet, path_users, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else if res.StatusCode == 200 {
		fmt.Println("Удалось(с нашей стороны)!")
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
	defer res.Body.Close()
}
func V2(Token string) {

	fmt.Println("Введите Имя для добавления:")
	name := ReadUser()

	data_str := `{"name":"` + name + `"}`
	//fmt.Println(data_str)
	data := []byte(data_str)
	req, err := http.NewRequest(http.MethodPost, path_users, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	defer req.Body.Close()
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	} else if res.StatusCode == 200 {
		fmt.Println("Удалось(с нашей стороны)!")
	}
	defer res.Body.Close()
}
func V3(Token string) {
	fmt.Println("Введите id для удаления:")
	id := ReadUser()
	//id_int,_:= strconv.Atoi(id)
	data := []byte(`{ "id":` + id + `}`)
	req, err := http.NewRequest(http.MethodDelete, path_user, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	} else if res.StatusCode == 200 {
		fmt.Println("Удалось(с нашей стороны)!")
	}
	defer res.Body.Close()
}
func V4(Token string) {
	fmt.Println("Введите id изменения:")
	id := ReadUser()
	fmt.Println("Введите Имя которое будет новое:")
	name := ReadUser()
	//id_int,_:= strconv.Atoi(id)
	//`{"name":"barisrt", "id": 4}`
	data := []byte(`{"name":"` + name + `", "id":` + id + `}`)
	req, err := http.NewRequest(http.MethodPut, path_user, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	} else if res.StatusCode == 200 {
		fmt.Println("Удалось(с нашей стороны)!")
	}
	defer res.Body.Close()
}
func V5_id(Token string) {
	fmt.Println("Введите id для поиска:")
	id := ReadUser()
	//id_int,_:= strconv.Atoi(id)
	data := []byte(`{"name":" ", "id":` + id + `}`)
	req, err := http.NewRequest(http.MethodGet, path_user_id, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else if res.StatusCode == 200 {
		fmt.Println("Удалось(с нашей стороны)!")
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
	defer res.Body.Close()
}
func V5_name(Token string) {
	fmt.Println("Введите Имя для поиска:")
	name := ReadUser()
	//id_int,_:= strconv.Atoi(id)
	data := []byte(`{ "name":"` + name + `", "id":0}`)
	req, err := http.NewRequest(http.MethodGet, path_user_name, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Token", Token)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else if res.StatusCode == 200 {
		fmt.Println("Удалось(с нашей стороны)!")
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
	defer res.Body.Close()
}
func V5(Token string) {
	fmt.Println("1.Найти по id")
	fmt.Println("2.Найти по имени")
	str := ReadUser()
	switch str {
	case "1":
		V5_id(Token)
	case "2":
		V5_name(Token)
	}

}

func ReactToAnswer(s string, Token string) {
	switch s {
	case "1":
		V1(Token)
	case "2":
		V2(Token)
	case "3":
		V3(Token)
	case "4":
		V4(Token)
	case "5":
		V5(Token)
	case "exit":
		return
	}

	PrintMenu(Token)
}
func PrintMenu(Token string) {
	fmt.Println("1.Найти всех")
	fmt.Println("2.Добавить")
	fmt.Println("3.Удалить")
	fmt.Println("4.Изменить")
	fmt.Println("5.Найти с фильтром")
	str := ReadUser()

	ReactToAnswer(str, Token)
	//fmt.Println(strings.TrimSuffix(str, "\n"))
}

/*
func CreateClient() (client *http.Client) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client = &http.Client{Transport: tr}
	return client
}*/

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}
var client = http.Client{Transport: tr}

func main() {

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

	//client:=CreateClient()

	PrintMenu(Token)

}
