package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"io/ioutil"

	"github.com/golang-jwt/jwt/v4"

	//"strconv"
	"strings"
	"time"

	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	_ "github.com/lib/pq"
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

var userMapA = map[string]string{
	"Admin":     "Admin",
	"testUser1": "testUser1",
	"testUser2": "testUser2",
}

type UserA struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func apiHandler(w http.ResponseWriter, req *http.Request) {
	middleware(w, req)
	if req.Method == "GET" {
		//http.ServeFile(w, req, "form.html")
		w.Write([]byte("Hello, world!"))
		fmt.Printf("hey")
		//fmt.Fprintf(w, "\nReceived time: %v\tType:GET\tRoute:/hello\n", time)
	}

}

func dataHandler(w http.ResponseWriter, req *http.Request) {
	middleware(w, req)
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if req.Method == "POST" {
		name := req.FormValue("name")
		occupation := req.FormValue("occupation")
		//fmt.Fprintf(w, req.PostFormValue("name"))
		fmt.Fprintf(w, "%s is a %s\n", name, occupation)
		fmt.Printf("%s is a %s\n", name, occupation)

		//fmt.Fprintf(w, "\nReceived time: %v\tType: POST\tRoute:/data\n", time)

	} else {
		//fmt.Fprintf(w, "\nReceived time: %v\tType:NOT POST\tRoute:/data\n", time)
	}

	//fmt.Fprintf(w, "ok\n")
}
func muxPathLastCut(s string) (last string) {
	mass := strings.SplitN(s, "/", 10)
	last = mass[len(mass)-1]
	return last
}

type users struct {
	name string
	id   int
}
type user struct {
	//Struct fields must start with upper case letter (exported) for the JSON package to see their value
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func checkStrForPostgre(s string) (res bool) {
	res = true
	bad := []rune("^[ -~]*$\"/*@")
	if strings.Contains(s, "/*") || strings.Contains(s, "\"") {
		//fmt.Println("SHOULD F")
		res = false
	}
	for i := 0; i < len(bad); i++ {
		//fmt.Printf("%s", string(bad[i]))
		if strings.Contains(s, string(bad[i])) {
			//fmt.Println("SHOULD F")
			res = false
		}
	}
	return res
}
func loginHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var user UserA

		// decode the request body into the struct, If error, respond to the client with the error message and a 400 status code.
		// fmt.Println(request.Body)
		// err := json.NewDecoder(request.Body).Decode(&user)
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			fmt.Fprintf(w, "invalid body")
			return
		}

		if userMapA[user.Username] == "" || userMapA[user.Username] != user.Password {
			fmt.Fprintf(w, "can not authenticate this user")
			return
		}

		token, err := generateJWT(user.Username)
		if err != nil {
			fmt.Fprintf(w, "error in generating token")
		}

		fmt.Fprintf(w, token)

	case "GET":
		fmt.Fprintf(w, "only POST methods is allowed.")
		return
	}

}

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
func validateToken(w http.ResponseWriter, r *http.Request) (authority string, err error) {

	if r.Header["Token"] == nil {
		fmt.Fprintf(w, "can not find token in header")
		return
	}

	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return sampleSecretKey, nil
	})

	if token == nil {
		fmt.Fprintf(w, "invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Fprintf(w, "couldn't parse claims")
		return "no", fmt.Errorf("There was an error in parsing")
	}
	//fmt.Println("Claims:", claims["username"].(string))//ВЫВЕСТИ ОПРЕДЕЛЁННУЮ ЧАСТЬ ТОКЕНА
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		fmt.Fprintf(w, "token expired")
		return "no", fmt.Errorf("Expired token")
	}

	return claims["username"].(string), nil
}

func usersHandler(w http.ResponseWriter, req *http.Request) {
	_, err := validateToken(w, req)
	if err != nil {
		return
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if req.Method == "POST" {
		header_type := req.Header.Get("Content-Type")
		if header_type != "application/json" {
			panic("NOT JSON")
		}
		u := user{}
		//var unmarshalErr *json.UnmarshalTypeError
		//decoder := json.NewDecoder(req.Body)
		//decoder.DisallowUnknownFields()
		//fmt.Println(decoder)
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(privdec(body), &u)
		if err != nil {
			panic(err)
		}

		//name := req.FormValue("name")
		//name := req.URL.Query().Get("name")
		//fmt.Println(req.GetBody())
		//id := req.FormValue("id")
		//id := req.URL.Query().Get("id")

		//id := req.FormValue("id")
		if checkStrForPostgre(u.Name) {
			result, err := db.Exec("insert into users (name) values ($1)",
				u.Name)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.RowsAffected()) // количество добавленных строк
		} else {
			fmt.Fprintf(w, "\nHello, name was cursed, try again\n")
		}

		//======================================
	} else if req.Method == "GET" {
		result, err := db.Query("SELECT * from users;")
		if err != nil {
			panic(err)
		}
		users_q := []users{}
		for result.Next() {
			u := users{}
			err := result.Scan(&u.name, &u.id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			users_q = append(users_q, u)

		}
		for _, u := range users_q {
			//fmt.Println(u.id, u.name)
			str := fmt.Sprintf("%d\t%s", u.id, u.name)
			//w.Write([]byte(str))
			fmt.Fprintf(w, "\n%s\n", str)
		}

		//fmt.Fprintf(w, "%s\n", result)
		//======================================
	}
	//fmt.Fprintf(w, "ok\n")
}
func usersFilteredByIdHolder(w http.ResponseWriter, req *http.Request) {
	_, err := validateToken(w, req)
	if err != nil {
		return
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if req.Method == "GET" {

		header_type := req.Header.Get("Content-Type")
		if header_type != "application/json" {
			panic("NOT JSON")
		}
		u := user{}
		//var unmarshalErr *json.UnmarshalTypeError
		//decoder := json.NewDecoder(req.Body)
		//decoder.DisallowUnknownFields()
		//fmt.Println(decoder)
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &u)
		if err != nil {
			panic(err)
		}
		//name := u.Name
		//id := u.Id
		//fmt.Println(name, id)
		result, err := db.Query("SELECT * from users where id= $1;", u.Id)
		if err != nil {
			panic(err)
		}
		users_q := []users{}
		for result.Next() {
			u := users{}
			err := result.Scan(&u.name, &u.id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			users_q = append(users_q, u)

		}
		for _, u := range users_q {
			//fmt.Println(u.id, u.name)
			str := fmt.Sprintf("%d\t%s", u.id, u.name)
			//w.Write([]byte(str))
			fmt.Fprintf(w, "\n%s\n", str)
		}

		//fmt.Fprintf(w, "ok\n")
		//======================================
	}
}
func usersFilteredByNameHolder(w http.ResponseWriter, req *http.Request) {
	_, err := validateToken(w, req)
	if err != nil {
		return
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if req.Method == "GET" {

		header_type := req.Header.Get("Content-Type")
		if header_type != "application/json" {
			panic("NOT JSON")
		}
		u := user{}
		//var unmarshalErr *json.UnmarshalTypeError
		//decoder := json.NewDecoder(req.Body)
		//decoder.DisallowUnknownFields()
		//fmt.Println(decoder)
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &u)
		if err != nil {
			panic(err)
		}

		if checkStrForPostgre(u.Name) {
			result, err := db.Query("SELECT * from users where name= $1;", u.Name)
			if err != nil {
				panic(err)
			}
			users_q := []users{}
			for result.Next() {
				u := users{}
				err := result.Scan(&u.name, &u.id)
				if err != nil {
					fmt.Println(err)
					continue
				}
				users_q = append(users_q, u)

			}
			for _, u := range users_q {
				//fmt.Println(u.id, u.name)
				str := fmt.Sprintf("%d\t%s", u.id, u.name)
				//w.Write([]byte(str))
				fmt.Fprintf(w, "\n %s\n", str)
			}

		} else {
			fmt.Fprintf(w, "\nHello, name was cursed, try again\n")
		}

		//fmt.Fprintf(w, "ok\n")
		//======================================
	}
}

func userHandler(w http.ResponseWriter, req *http.Request) {
	role, err := validateToken(w, req)
	if err != nil {
		return
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	/*
		id_user := muxPathLastCut(req.URL.Path)
		id_user_dig, _ := strconv.Atoi(id_user)
		fmt.Fprintf(w, "working on %d\n", id_user_dig)*/
	if req.Method == "PUT" {
		header_type := req.Header.Get("Content-Type")
		if header_type != "application/json" {
			panic("NOT JSON")
		}
		u := user{}
		//var unmarshalErr *json.UnmarshalTypeError
		//decoder := json.NewDecoder(req.Body)
		//decoder.DisallowUnknownFields()
		//fmt.Println(decoder)
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(privdec(body), &u)
		if err != nil {
			panic(err)
		}

		//name := req.FormValue("name")
		//name := req.URL.Query().Get("name")
		//fmt.Println(req.GetBody())
		//id := req.FormValue("id")
		//id := req.URL.Query().Get("id")
		name := u.Name
		id := u.Id
		fmt.Println(name, id)
		//idd, _ := strconv.Atoi(id)

		if checkStrForPostgre(u.Name) {
			result, err := db.Exec("update users set name = $1 where id = $2",
				name, id)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.RowsAffected()) // количество добавленных строк

		} else {
			fmt.Fprintf(w, "\nHello, name was cursed, try again\n")
		}

		//======================================
	} else if req.Method == "GET" {

		header_type := req.Header.Get("Content-Type")
		if header_type != "application/json" {
			panic("NOT JSON")
		}
		u := user{}
		//var unmarshalErr *json.UnmarshalTypeError
		//decoder := json.NewDecoder(req.Body)
		//decoder.DisallowUnknownFields()
		//fmt.Println(decoder)
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &u)
		if err != nil {
			panic(err)
		}

		result, err := db.Query("SELECT * from users where id= $1;", u.Id)
		if err != nil {
			panic(err)
		}
		users_q := []users{}
		for result.Next() {
			u := users{}
			err := result.Scan(&u.name, &u.id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			users_q = append(users_q, u)

		}
		for _, u := range users_q {
			fmt.Println(u.id, u.name)
			str := fmt.Sprintf("%d\t%s", u.id, u.name)
			//w.Write([]byte(str))
			fmt.Fprintf(w, "\n %s\n", str)
		}

		//fmt.Fprintf(w, "ok\n")
		//======================================
	} else if req.Method == "DELETE" {
		//validateToken
		if role != "Admin" {
			return
		}
		header_type := req.Header.Get("Content-Type")
		if header_type != "application/json" {
			panic("NOT JSON")
		}
		u := user{}
		//var unmarshalErr *json.UnmarshalTypeError
		//decoder := json.NewDecoder(req.Body)
		//decoder.DisallowUnknownFields()
		//fmt.Println(decoder)
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(privdec(body), &u)
		if err != nil {
			panic(err)
		}
		result, err := db.Exec("delete from users where id = $1",
			u.Id)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.RowsAffected()) // количество добавленных строк
		//======================================
	}

}
func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", middleware)
	mux.HandleFunc("/data", dataHandler)
	mux.HandleFunc("/hello", apiHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/id", usersFilteredByIdHolder)
	mux.HandleFunc("/users/name", usersFilteredByNameHolder)
	mux.HandleFunc("/user", userHandler)
	mux.HandleFunc("/login", loginHandler)
	//http.HandleFunc("/login", loginHandler)
}
func middleware(w http.ResponseWriter, req *http.Request) {
	time := time.Now().Format(time.ANSIC)
	fmt.Printf("\nExecuting middleware at: %v Method:%s URL:%s\n", time, req.Method, req.URL.Path)
	fmt.Print("\nExecuting middleware\n")
}

var connStr string = "user=postgres password=1 dbname=pp sslmode=disable"

func main() {

	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	//router := mux.NewRouter()
	mux := http.NewServeMux()
	setupHandlers(mux)

	log.Fatal(http.ListenAndServe(listenAddr, mux))

}
