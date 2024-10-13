package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"

	//"encoding/json"
	"strings"
	//"net/url"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"time"
)

func publenc(mes []byte) []byte {
	publicKeyPEM, err := ioutil.ReadFile("./public.pem") //ШИФРОВКА(Загразука ключа и его загрузка в переменную)
	if err != nil {
		panic(err)
	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEM)
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	//ШИФРОВКА
	//plaintext := []byte(mes)                                                                   //ПЕРВЫЙ ПОСЫЛАЕТ
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), mes) //ОПЯТЬ ЖЕ ПЕРВЫЙ ПАРАМЕТР МОЖНО СВОЙ РИДЕР ВЗЯТЬ, не безопасно
	if err != nil {
		panic(err)
	}

	//fmt.Printf("Encrypted: %x\n", ciphertext)
	return ciphertext
}
func privdec(ct []byte) []byte { //СЕРВЕР БЕРЁТ ЗАКРЫТЫЙ КЛЮЧ
	privateKeyPEM, err := ioutil.ReadFile("./private.pem")
	if err != nil {
		panic(err)
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	//ДЕШИФРОВКА
	//ciphertext := []byte{0x88, 0xaa, 0x63, 0x24, 0x2d, 0x48, 0xfd, 0xb1, 0x63, 0x71, 0x33, 0x17, 0x2a, 0x01, 0xce, 0x15, 0x1b, 0x25, 0xac, 0xcd, 0x35, 0xc1, 0x7c, 0x2a, 0x48, 0x58, 0x79, 0xae, 0x73, 0xf3, 0x5e, 0xc9, 0x89, 0xa7, 0x8a, 0x92, 0xa4, 0x3f, 0x3d, 0xb3, 0x43, 0x1d, 0x01, 0x74, 0xee, 0xd1, 0x1e, 0x95, 0x2b, 0x4f, 0x42, 0x46, 0x0b}
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ct) //СЕРВЕР ПРИНЯЛ ЗАШИФРОВАННЫЙ КЛЮЧОМ ШИФР И ИСПОЛЬЗУЕТ СВОЙ УЛЬТИМАТИВНЫЙ ЗАКР КЛЮЧ
	if err != nil {
		panic(err)
	}

	//fmt.Printf("Decrypted: %s\n", plaintext)
	return plaintext
}

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
	data = publenc(data)
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
	req, err := http.NewRequest(http.MethodDelete, path_user, bytes.NewBuffer(publenc(data)))

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
	req, err := http.NewRequest(http.MethodPut, path_user, bytes.NewBuffer(publenc(data)))

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
