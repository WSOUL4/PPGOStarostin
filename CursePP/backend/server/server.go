package main

import (
	"fmt"
	//"io"
	"log"
	"net/http"
	"os"
	"time"

	//"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v4"
	//"bufio"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"

	//"strconv"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"strings"

	_ "github.com/lib/pq"
)

// ФУНКЦИИ-САППОРТЫ
func muxPathLastCut(s string) (last string) {
	mass := strings.SplitN(s, "/", 10)
	last = mass[len(mass)-1]
	return last
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
func publenc(mes []byte) []byte {
	publicKeyPEM, err := ioutil.ReadFile("./client.key") //ШИФРОВКА(Загразука ключа и его загрузка в переменную)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEM)
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	//ШИФРОВКА
	//plaintext := []byte(mes)                                                                   //ПЕРВЫЙ ПОСЫЛАЕТ
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), mes) //ОПЯТЬ ЖЕ ПЕРВЫЙ ПАРАМЕТР МОЖНО СВОЙ РИДЕР ВЗЯТЬ, не безопасно
	if err != nil {
		log.Println(err)
		panic(err)
	}
	//fmt.Printf("Encrypted: %x\n", ciphertext)
	return ciphertext
}
func privdec(ct []byte) []byte { //СЕРВЕР БЕРЁТ ЗАКРЫТЫЙ КЛЮЧ
	/*
		file, err := os.Open("./server.key")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var line string
		for scanner.Scan() {
			line += scanner.Text() + "\n"
		}
		//fmt.Println(line)
		//privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	*/
	privateKey := `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC8HYSVe0MvlzOJ
J1xgMY9K71rWyMRfqrRsagergkdoVmL4vdf1kE1vkOZfJlukLq9Nw+dst3Kir9Dr
R1TloyphzOCePt3Zte/LHBxb5I6PhOAt2jT/LjVFH13Y7nJfhcEMzngcdGohnIqM
XGbgC/r9xzw6x6JUO7MbacrD4Ecdil7QtnMq7TtvLp+ejgzTdVPHTKXrJ+8I9F17
cU/wJvrbbG/V/RZlasqYzs4I5OhCRuYoSsjdqm8TKdpEVB7PWT8LanD7+8fsyISD
7SDYvVmo9BhAcPXgZMOLKaw7UU7+0boHIIWu3R7wgiKEjv64z5PnA0uGqY6H4lqN
YPaIKFUdAgMBAAECggEACmriuhnRZq6g6VTcsMG4hcCA71QUS+1QKzt+oT4zENF3
5FPfGczxUjRcDaOPf6V2NSrxg1EaxWz8d4sZEH6GDP8jkUqPTqHcswrOJz3zsc78
peK8/JoP09F2M6/rdY7FWh2KWUkOd5NhDiov0vMsOybwFvE8nsg6QFJjgDrItWyp
SceVqS2AQi5XNHhPs6FnPpweIHDADdRONKA1JwBirGdsl2jz9QHOQ7aWo4U6qgmH
noYA0tOg3DsOamSovi2zScS2seSSs7tqxzuKdivFC3kCS4zM+0xO4Srg6VUmGlxi
vgZatxe/M1eMxwVVgVO1HtG9k9RVJTB4HMyr4mu7oQKBgQDfDUj5qjK7biy2jM5a
YIEuSoqo62vug11BQKcbSvARPc14PeP+izZS4sSzxhVtB3XvU2LIVLeceUY7gr10
DZTQSUnz3Ugs9NNtD8At0QqoF21TBTh+Cz8/wEGKr7EIDiG17rqLC+PxENXSSOqW
FGd4eYQ98ZYok00tcdMBLJG1SwKBgQDX5xrAXJQ3C3OcgPLA17IcFOEQ/e7A225i
gyGLmho0I/ZcrzgUbCNO4YYUqExRVfuLnNU9Cyx/GJQOekVgQijF4nq+HJcREmxz
PxANDgesuzuwaD1SNMwD4Gn1nCCZXHrNAD6dVR2aMBQicCkS9uVNGF5j5vOKhucA
KDIebUvmNwKBgQDJVaCNW4e3j1dk3+xTv9BbDIXku7wM6x5+C/HKvPW9Wl/hLVxj
Ix3B61dKPn6Qj69we4Aq++1QnFc82GJSIwA0kjLioNbJXaSKSTFbKdnIqHzR92Bq
xZQt40hF+xh9AOSE6BwR7oWtz1hyG8dD+N787BLmJu83aN69KoUgBi7vyQKBgDOl
T8vmGXpVXfF5Exi4QB3hjLkg1UUC+JPOJG8djNkeJSekrniMKaIL5qP4YlEujT6n
ZIb2rk001u3jp8bP7Krxc0UY17Y4vwKOekt1KLbUDwIy3UBV6tueiho7n7yv15xE
S7YdDzi7+YUHaXvk3ZMkmiexrl5byNRLyTloEbjfAoGBAJOYVCJFSYL25cgDxgab
jX+hgciKDqDpeZREY83fNtPuY6d0YomZPgY4d2cogXBmgrZpZAt9F56FZORz/CHX
i4fIonoFxVYZKrkm6llJnKnhebiXV+cnUNy8lEZYqQGT4p+PKfpJA5I6l93v1+Yz
Imps4ZUheDjYnQGZeJnKnESz
-----END PRIVATE KEY-----
`
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		log.Println("Failed to decode")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		//fmt.Println("Error decrypting data:", err)
		log.Println(err)
	}
	decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key, ct, nil)
	if err != nil {
		//fmt.Println("Error decrypting data:", err)
		log.Println(err)
	}
	return decryptedData
}

// СТРУКТУРЫ ДЛЯ ЗАПРОСОВ И ПАРСИНГА
var userMapA = map[string]string{
	"Admin":     "Admin",
	"testUser1": "testUser1",
	"testUser2": "testUser2",
}

type UserA struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type users struct { //Сделать подобные структуры для всех таблиц  или запросов!!!
	name string
	id   int
}
type user struct {
	//Struct fields must start with upper case letter (exported) for the JSON package to see their value
	Name string `json:"name"`
	Id   int    `json:"id"`
}
type TaskStruct struct {
	Id       int    `json:"id"`
	DateBeg  string `json:"dateBeg"`
	DateEnd  string `json:"dateEnd"`
	ParentId int    `json:"parentId"`
	Task     string `json:"task"`
	Mark     string `json:"mark"`
	Status   string `json:"status"`
}

type DoneTaskStruct struct {
	Id          int    `json:"id"`
	DateBeg     string `json:"dateBeg"`
	DateEnd     string `json:"dateEnd"`
	ParentId    int    `json:"parentId"`
	Task        string `json:"task"`
	Mark        string `json:"mark"`
	Status      string `json:"status"`
	DateEndReal string `json:"dateEndReal"`
}
type DistributionStruct struct {
	IdTask   int `json:"idTask"`
	IdWorker int `json:"idWorker"`
}
type DistributionStruct2 struct {
	IdTask int    `json:"idTask"`
	Login  string `json:"login"`
}
type Worker struct {
	Login string `json:"login"`
}
type LoginStruct struct {
	Login string `json:"login"`
	//Email string `json:"email"`
	Password string `json:"password"`
}
type LoginStruct2 struct {
	Login []byte `json:"login"`
	//Email string `json:"email"`
	Password []byte `json:"password"`
}
type RegStruct struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	FIO      string `json:"fio"`
	Password string `json:"password"`
}
type RegStruct2 struct {
	Login    []byte `json:"login"`
	Email    []byte `json:"email"`
	FIO      []byte `json:"fio"`
	Password []byte `json:"password"`
}
type ParentIdForm struct {
	PI int `json:"pi"`
}
type RespFormTask struct {
	TaskForm    TaskStruct `json:"TaskF"`
	WorkersForm []Worker   `json:"WorkersF"`
}
type hierarchy struct {
	ParentId int    `json:"parentId"`
	Id       int    `json:"id"`
	Depth    int    `json:"depth"`
	Status   string `json:"status"`
}
type Token struct {
	TokenStr string `json:"tokenStr"`
}
type AnswerStr struct {
	Text string `json:"text"`
}
type LeftStruct struct {
	Id       int    `json:"id"`
	Task     string `json:"task"`
	Mark     string `json:"mark"`
	Status   string `json:"status"`
	DaysHad  int    `json:"daysHad"`
	DaysLeft int    `json:"daysLeft"`
}
type DoneForStruct struct {
	Id         int    `json:"id"`
	Task       string `json:"task"`
	Mark       string `json:"mark"`
	Status     string `json:"status"`
	DaysPassed string `json:"daysPassed"`
}

// ТОКЕНЫ СЕССИИ
func DecodeToken(tokenStr string) (map[string]interface{}, error) {
	// Разделяем токен на части
	parts := strings.Split(tokenStr, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token")
	}

	// Декодируем полезную нагрузку
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	// Парсим полезную нагрузку в map
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, err
	}

	return claims, nil
}
func generateJWT(userID string) (string, error) {
	secretKey := sampleSecretKey
	// Создаем новый токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(), // срок действия токена
	})

	// Подписываем токен с использованием ключа
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func validateToken(tokenStr string) (jwt.Claims, error) {
	secretKey := sampleSecretKey
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

/*
func generateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		log.Printf("Something Went Wrong while generating token: %s\n", err.Error())
		return "", err
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)
	return tokenString, nil
}

func validateToken(w http.ResponseWriter, r *http.Request) (authority string, err error) {

	if r.Header["Authorization"] == nil {
		fmt.Fprintf(w, "can not find token in header")
		return
	}

	token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
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
*/
/*
func validateToken(tokenString string) error {
	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем тип подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неверный метод подписи %v", token.Header["alg"])
		}
		// Возвращаем секретный ключ, который мы использовали для подписи токена
		return sampleSecretKey, nil
	})

	if err != nil {
		return fmt.Errorf("ошибка при разборе токена: %v", err)
	}

	// Проверяем, что токен действителен и не истек
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		authorized := claims["authorized"].(bool)
		if !authorized {
			return fmt.Errorf("не авторизован")
		}
		return nil
	}
	return fmt.Errorf("неверный токен")
}*/

// ФУНКЦИИ ДЛЯ МАРШРУТОВ

// Настройка маршрутов
func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", middleware)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/registration", registration)
	mux.HandleFunc("/selectAll", selectAll)
	mux.HandleFunc("/data", data1)
	mux.HandleFunc("/selectByParent", selectByParent)
	mux.HandleFunc("/children", children)
	mux.HandleFunc("/parents", parents)
	mux.HandleFunc("/createTask", createTask)
	mux.HandleFunc("/changeTask", changeTask)
	mux.HandleFunc("/deleteTask", deleteTask)
	mux.HandleFunc("/task/findByTask", findByTask)
	mux.HandleFunc("/task/addWorker", addWorker)
	mux.HandleFunc("/report/left", RepLeft)
	mux.HandleFunc("/report/forLast", RepforLast)
	fmt.Println("Сервер запущен")
}
func RepLeft(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}
	// Получаем токен из заголовков запроса
	tokenString := req.Header.Get("Authorization")

	// Проверяем токен
	cl, err := validateToken(tokenString) //cl
	if err != nil {
		http.Error(w, " BAD TOKEN", http.StatusUnauthorized)
		return
	}
	//Узнаём кто сделал запрос
	claims, _ := cl.(jwt.MapClaims)
	sub := claims["sub"]
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, " Database died<:(|", http.StatusServiceUnavailable)
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "POST" {

		var data ParentIdForm
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {

			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}
		result, err := db.Query(`WITH LeftD as (SELECT id,  task, mark, status, "dateEnd"-"dateBeg" as daysHad, "dateEnd"-CURRENT_DATE as daysLeft
	FROM public."Tasks"),usr as(select id from public."Workers" where "login"=$1),
	his as(select "idTask" from public."Distribution" where "idWorker"= (select * from usr))

	select * from LeftD where daysLeft<=$2 and status not like 'done' and id in (select * from his) order by daysLeft asc
`, sub, data.PI)
		if err != nil {
			fmt.Println("Ошибка выполнения запроса:", err)
			return
		}

		datamass := []LeftStruct{}
		// Проходим по всем строкам результата
		for result.Next() {
			t := LeftStruct{}
			err := result.Scan(&t.Id, &t.Task, &t.Mark, &t.Status, &t.DaysHad, &t.DaysLeft)
			if err != nil {
				fmt.Println(err)
				continue
			}
			datamass = append(datamass, t)
		}

		jsonResponse, err := json.Marshal(datamass)
		if err != nil {
			fmt.Println("Ошибка создания ответа:", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	}
}
func RepforLast(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}
	// Получаем токен из заголовков запроса
	tokenString := req.Header.Get("Authorization")

	// Проверяем токен
	cl, err := validateToken(tokenString) //cl
	if err != nil {
		http.Error(w, " BAD TOKEN", http.StatusUnauthorized)
		return
	}
	//Узнаём кто сделал запрос
	claims, _ := cl.(jwt.MapClaims)
	sub := claims["sub"]
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, " Database died<:(|", http.StatusServiceUnavailable)
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "POST" {

		var data ParentIdForm
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {

			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}
		result, err := db.Query(`WITH LeftD as (SELECT id,  task, mark, status, CURRENT_DATE-"dateEnd" as daysPassed
	FROM public."Tasks"),usr as(select id from public."Workers" where "login"=$1),
	his as(select "idTask" from public."Distribution" where "idWorker"= (select * from usr))

	select distinct * from LeftD where daysPassed between 0 and $2 and status like 'done' and id in (select * from his) order by daysPassed asc
	`, sub, data.PI)
		if err != nil {
			fmt.Println("Ошибка выполнения запроса:", err)
			return
		}

		datamass := []DoneForStruct{}
		// Проходим по всем строкам результата
		for result.Next() {
			t := DoneForStruct{}
			err := result.Scan(&t.Id, &t.Task, &t.Mark, &t.Status, &t.DaysPassed)
			if err != nil {
				fmt.Println(err)
				continue
			}
			datamass = append(datamass, t)
		}

		jsonResponse, err := json.Marshal(datamass)
		if err != nil {
			fmt.Println("Ошибка создания ответа:", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	}
}
func addWorker(w http.ResponseWriter, req *http.Request) { //DistributionStruct
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}
	// Получаем токен из заголовков запроса
	tokenString := req.Header.Get("Authorization")

	// Проверяем токен
	_, err := validateToken(tokenString) //cl
	if err != nil {
		http.Error(w, " BAD TOKEN", http.StatusUnauthorized)
		return
	}
	//Узнаём кто сделал запрос
	//claims, _ := cl.(jwt.MapClaims)
	//sub := claims["sub"]
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, " Database died<:(|", http.StatusServiceUnavailable)
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "POST" {
		var data DistributionStruct2
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {

			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}

		//fmt.Println(data)
		_, err := db.Query(`
WITH usr as(select id from public."Workers" where "login"=$1)
INSERT INTO public."Distribution"(
	"idTask", "idWorker")
	VALUES ($2, (select * from usr));`, data.Login, data.IdTask)
		if err != nil {
			fmt.Println("Ошибка выполнения запроса:", err)
			a := AnswerStr{Text: "Error Inserting Worker"}
			jsonResponse, err := json.Marshal(a)
			if err != nil {
				fmt.Println("Ошибка создания ответа:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		} else {
			a := AnswerStr{Text: "OK"}
			jsonResponse, err := json.Marshal(a)
			if err != nil {
				fmt.Println("Ошибка создания ответа:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		}
	}
}
func findByTask(w http.ResponseWriter, req *http.Request) {

}
func deleteTask(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET,DELETE, OPTIONS")                // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}
	// Получаем токен из заголовков запроса
	tokenString := req.Header.Get("Authorization")

	// Проверяем токен
	_, err := validateToken(tokenString) //cl
	if err != nil {
		http.Error(w, " BAD TOKEN", http.StatusUnauthorized)
		return
	}
	//Узнаём кто сделал запрос
	//claims, _ := cl.(jwt.MapClaims)
	//sub := claims["sub"]
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, " Database died<:(|", http.StatusServiceUnavailable)
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "DELETE" {
		var data ParentIdForm
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {

			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}

		//fmt.Println(data)
		_, err := db.Query(`DELETE FROM public."Tasks"
	WHERE id=$1;`, data.PI)
		if err != nil {
			fmt.Println("Ошибка выполнения запроса:", err)
			a := AnswerStr{Text: "Error Update"}
			jsonResponse, err := json.Marshal(a)
			if err != nil {
				fmt.Println("Ошибка создания ответа:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		} else {
			a := AnswerStr{Text: "OK"}
			jsonResponse, err := json.Marshal(a)
			if err != nil {
				fmt.Println("Ошибка создания ответа:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		}
	}
}
func changeTask(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}
	// Получаем токен из заголовков запроса
	tokenString := req.Header.Get("Authorization")

	// Проверяем токен
	_, err := validateToken(tokenString) //cl
	if err != nil {
		http.Error(w, " BAD TOKEN", http.StatusUnauthorized)
		return
	}
	//Узнаём кто сделал запрос
	//claims, _ := cl.(jwt.MapClaims)
	//sub := claims["sub"]
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, " Database died<:(|", http.StatusServiceUnavailable)
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "POST" {
		var data TaskStruct
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {

			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}

		//fmt.Println(data)
		_, err := db.Query(`UPDATE public."Tasks"
	SET "dateBeg"=$1, "dateEnd"=$2, "parentId"=$3, task=$4, mark=$5, status=$6
	WHERE id=$7;`, data.DateBeg, data.DateEnd, data.ParentId, data.Task, data.Mark, data.Status, data.Id)
		if err != nil {
			fmt.Println("Ошибка выполнения запроса:", err)
			a := AnswerStr{Text: "Error Update"}
			jsonResponse, err := json.Marshal(a)
			if err != nil {
				fmt.Println("Ошибка создания ответа:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		} else {
			a := AnswerStr{Text: "OK"}
			jsonResponse, err := json.Marshal(a)
			if err != nil {
				fmt.Println("Ошибка создания ответа:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		}
	}
}
func createTask(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}
	// Получаем токен из заголовков запроса
	tokenString := req.Header.Get("Authorization")

	// Проверяем токен
	cl, err := validateToken(tokenString)
	if err != nil {
		http.Error(w, " BAD TOKEN", http.StatusUnauthorized)
		return
	}
	//Узнаём кто сделал запрос
	claims, _ := cl.(jwt.MapClaims)
	sub := claims["sub"]
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, " Database died<:(|", http.StatusServiceUnavailable)
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "POST" {

		var data TaskStruct
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {

			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}

		//fmt.Println(data)
		result, err := db.Query(`INSERT INTO public."Tasks"(
	 "dateBeg", "dateEnd", "parentId", task, mark, status)
	VALUES ( $1, $2, $3, $4, $5, $6)
	RETURNING "id";`, data.DateBeg, data.DateEnd, data.ParentId, data.Task, data.Mark, data.Status)
		if err != nil {
			fmt.Println("Ошибка выполнения запроса:", err)
			a := AnswerStr{Text: "Error Insert"}
			jsonResponse, err := json.Marshal(a)
			if err != nil {
				fmt.Println("Ошибка создания ответа:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		} else {
			var t TaskStruct
			for result.Next() {
				err = result.Scan(&t.Id)

			}
			//fmt.Println(t.Id)
			_, err = db.Query(`WITH usr as(select id from public."Workers" where "login"=$2)
INSERT INTO public."Distribution"(
	"idTask", "idWorker")
	VALUES ($1, (select * from usr limit 1));`, t.Id, sub)

			a := AnswerStr{Text: "OK"}
			jsonResponse, err := json.Marshal(a)
			if err != nil {
				fmt.Println("Ошибка создания ответа:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		}

	}
}
func login(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}

	if req.Method == "POST" {

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		defer db.Close()

		var data LoginStruct
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {

			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}
		//fmt.Println(data)
		result, err := db.Query(`SELECT p.login,p.password
	FROM public."Workers" as p
	where (p.email=$1 and p.password=$2) or (p.login=$1 and p.password=$2);`, data.Login, data.Password)
		t := LoginStruct{}
		//fmt.Println(result)
		for result.Next() {

			err := result.Scan(&t.Login, &t.Password)
			if err != nil {
				//fmt.Println(err)
				http.Error(w, err.Error(), http.StatusNotFound)
				//jsonResponse, _ := json.Marshal("Login Error!")
				//w.Header().Set("Content-Type", "application/json")
				//w.Write(jsonResponse)
				return
			}
			//fmt.Println(t.Login)

		}
		//jsonResponse, err := json.Marshal(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if t.Login != "" {
			token, err := generateJWT(t.Login)
			if err != nil {
				fmt.Fprintf(w, "error in generating token")
				return
			}
			tokenS := Token{}
			tokenS.TokenStr = token
			jsonResponse, err := json.Marshal(tokenS)
			//fmt.Fprintf(w, token)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		}

	}
}
func registration(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")

	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "POST" {
		var data RegStruct
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}

		result, err := db.Query(`SELECT id
	FROM public."Workers"
	WHERE login=$1;`, data.Login)
		if err != nil {
			fmt.Println("Ошибка выполнения запроса:", err)
			return
		}
		count := 0

		// Проходим по всем строкам результата
		for result.Next() {
			count++ // Увеличиваем счетчик на 1 для каждой строки
		}
		if count != 0 {
			a := AnswerStr{Text: "HAS"}
			jsonResponse, err := json.Marshal(a)
			if err != nil {
				fmt.Println("Ошибка создания ответа:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		} else {
			_, err := db.Query(`INSERT INTO public."Workers"(
	fio, email, login, password)
	VALUES ($1, $2, $3, $4);`, data.FIO, data.Email, data.Login, data.Password)
			if err != nil {
				fmt.Println("Ошибка выполнения запроса:", err)
				return
			}

			a := AnswerStr{Text: "OK"}
			jsonResponse, err := json.Marshal(a)
			if err != nil {
				fmt.Println("Ошибка создания ответа:", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		}

	}

}
func children(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}

	// Получаем токен из заголовков запроса
	tokenString := req.Header.Get("Authorization")

	// Проверяем токен
	cl, err := validateToken(tokenString)
	if err != nil {
		http.Error(w, "TOKEN", http.StatusUnauthorized)
		return
	}
	//Узнаём кто сделал запрос
	claims, _ := cl.(jwt.MapClaims)
	sub := claims["sub"]

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "GET" {
		result, err := db.Query(`
with recursive whosYourChildren("parentId", id,depth) as (

  --start with the "anchor" row
  select
    "parentId", id,0
  from public."Tasks" as t
  where
    id = 0 --parameterize me

  union all

  select
    t."parentId", t.id, whosYourChildren.depth + 1
  from public."Tasks" as t
  join whosYourChildren on whosYourChildren.id = t."parentId"
),

 allt("parentId", id,depth, "status") AS ( 
 select whosYourChildren."parentId",whosYourChildren.id,whosYourChildren.depth,t.status
from whosYourChildren, public."Tasks" as t
where whosYourChildren."parentId" is not null and depth>0 and t.id=whosYourChildren.id
order by
  whosYourChildren.depth
  )
  ,BYAUTH AS ( select distinct t.id
from public."Tasks" as t, public."Distribution" as d,public."Workers" as w
where (t.id= d."idTask" and d."idWorker"=w.id and w.login=$1) or t.id=0)


select distinct * from allt where allt.id in (select * from BYAUTH)`, sub)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		tree := []hierarchy{}
		for result.Next() {
			t := hierarchy{}
			err := result.Scan(&t.ParentId, &t.Id, &t.Depth, &t.Status)
			if err != nil {
				fmt.Println(err)
				continue
			}
			tree = append(tree, t)

		}
		//fmt.Println(tree)
		jsonResponse, err := json.Marshal(tree)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	} else if req.Method == "POST" {
		var data TaskStruct
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}

		result, err := db.Query(`
with recursive whosYourChildren("parentId", id,depth) as (

  --start with the "anchor" row
  select
    "parentId", id,0
  from public."Tasks" as t
  where
    id = $1 --parameterize me

  union all

  select
    t."parentId", t.id, whosYourChildren.depth + 1
  from public."Tasks" as t
  join whosYourChildren on whosYourChildren.id = t."parentId"
)

select 
  *
from whosYourChildren
where "parentId" is not null and depth>0
order by
  whosYourChildren.depth;`, data.Id)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		tree := []hierarchy{}
		for result.Next() {
			t := hierarchy{}
			err := result.Scan(&t.ParentId, &t.Id, &t.Depth)
			if err != nil {
				fmt.Println(err)
				continue
			}
			tree = append(tree, t)

		}
		jsonResponse, err := json.Marshal(tree)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	}
}
func parents(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}

	// Получаем токен из заголовков запроса
	tokenString := req.Header.Get("Authorization")

	// Проверяем токен
	_, err := validateToken(tokenString)
	if err != nil {
		http.Error(w, "TOKEN", http.StatusUnauthorized)
		return
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "POST" {
		var data TaskStruct
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}

		result, err := db.Query(`
with recursive whosYourDaddy("parentId", id,depth) as (

  --start with the "anchor" row
  select
    "parentId", id,1
  from public."Tasks" as t
  where
    id = $1 --parameterize me

  union all

  select
    t."parentId", t.id, whosYourDaddy.depth + 1
  from public."Tasks" as t
  join whosYourDaddy on whosYourDaddy."parentId" = t.id
)

select
  *
from whosYourDaddy
where "parentId" is not null
order by
  whosYourDaddy.depth;`, data.Id)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		tree := []hierarchy{}
		for result.Next() {
			t := hierarchy{}
			err := result.Scan(&t.ParentId, &t.Id, &t.Depth)
			if err != nil {
				fmt.Println(err)
				continue
			}
			tree = append(tree, t)

		}
		jsonResponse, err := json.Marshal(tree)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	}
}
func selectByParent(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}
	// Получаем токен из заголовков запроса
	tokenString := req.Header.Get("Authorization")

	// Проверяем токен
	_, err := validateToken(tokenString)
	if err != nil {
		http.Error(w, "TOKEN", http.StatusUnauthorized)
		return
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "POST" {
		//application/json; charset=utf-8
		contentType := req.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Данные не в формате JSON", http.StatusUnsupportedMediaType)
			return
		}
		var data ParentIdForm
		err = json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Ошибка парсинга JSON данных: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Использование полученных данных
		parentId := data.PI
		//fmt.Printf("Received parentId: %d\n", parentId)
		//return
		result, err := db.Query("SELECT * from public.\"Tasks\" where \"parentId\" IS NOT NULL AND \"parentId\"=$1", parentId)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		task_q := []TaskStruct{}
		task_r := []Worker{}
		task_f := []RespFormTask{}
		for result.Next() {
			t := TaskStruct{}
			err := result.Scan(&t.Id, &t.DateBeg, &t.DateEnd, &t.ParentId, &t.Task, &t.Mark, &t.Status)
			if err != nil {
				fmt.Println(err)
				continue
			}
			task_q = append(task_q, t)
			resultResponsibles, err := db.Query(`SELECT distinct w."login"
FROM public."Distribution" as d,  public."Workers" as w
WHERE d."idWorker" = w.id and d."idWorker"=1 and d."idTask"=$1`, t.Id)
			if err != nil {
				log.Println(err)
				panic(err)
			}
			for resultResponsibles.Next() {
				r := Worker{}
				err := resultResponsibles.Scan(&r.Login)
				//fmt.Println(r)
				if err != nil {
					fmt.Println(err)
					continue
				}
				task_r = append(task_r, r)
			}
			f := RespFormTask{}
			f.TaskForm = t
			f.WorkersForm = task_r
			task_f = append(task_f, f)
			//fmt.Println(task_r)
			task_r = []Worker{}
		}
		jsonResponse, err := json.Marshal(task_f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}

}

func data1(w http.ResponseWriter, req *http.Request) { //Проверочная функция
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}
	json.NewEncoder(w).Encode("OKOK")
}

func selectAll(w http.ResponseWriter, req *http.Request) {
	// Устанавливаем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                                   // Разрешаем все источники
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                 // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Authorization") // Разрешаем заголовки
	//w.Header().Set("Access-Control-Allow-Headers", "Token")
	// Обработка запросов OPTIONS (предварительные запросы CORS)
	if req.Method == http.MethodOptions {
		return
	}
	// Получаем токен из заголовков запроса
	tokenString := req.Header.Get("Authorization")

	// Проверяем токен
	_, err := validateToken(tokenString)
	if err != nil {
		http.Error(w, "TOKEN", http.StatusUnauthorized)
		return
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	if req.Method == "GET" {
		result, err := db.Query("SELECT * from public.\"Tasks\" where \"parentId\" IS NOT NULL")
		if err != nil {
			log.Println(err)
			panic(err)
		}
		task_q := []TaskStruct{}
		task_r := []Worker{}
		task_f := []RespFormTask{}
		for result.Next() {
			t := TaskStruct{}
			err := result.Scan(&t.Id, &t.DateBeg, &t.DateEnd, &t.ParentId, &t.Task, &t.Mark, &t.Status)
			if err != nil {
				fmt.Println(err)
				continue
			}
			task_q = append(task_q, t)
			resultResponsibles, err := db.Query("SELECT distinct w.\"login\" FROM public.\"Distribution\" as d,  public.\"Workers\" as w WHERE d.\"idWorker\" = w.id and d.\"idWorker\"=1 and d.\"idTask\"=$1", t.Id)
			if err != nil {
				log.Println(err)
				panic(err)
			}
			for resultResponsibles.Next() {
				r := Worker{}
				err := resultResponsibles.Scan(&r.Login)
				if err != nil {
					fmt.Println(err)
					continue
				}
				task_r = append(task_r, r)
			}
			f := RespFormTask{}
			f.TaskForm = t
			f.WorkersForm = task_r
			task_f = append(task_f, f)
		}
		jsonResponse, err := json.Marshal(task_f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}
func middleware(w http.ResponseWriter, req *http.Request) {
	time := time.Now().Format(time.ANSIC)
	log.Printf("\nExecuting middleware at: %v Method:%s URL:%s\n", time, req.Method, req.URL.Path)
	log.SetOutput(logFile)
}

/*
ГЛОБАЛЬНЫЕ ПЕРЕМЕННЫЕ GLOBAL VARIABLES
*/
var connStr string = "user=postgres password=1 dbname=pp sslmode=disable"

// var db *sql.DB
var logFile, _ = os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
var sampleSecretKey = []byte("Abrakadabra") //os.ReadFile("server.pem")
// ГЛАВНАЯ
func main() {
	//Попытка подключчиться к порту
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	//Попытка подключчиться к БД
	/*
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
			log.Fatal(err)

		}
		defer db.Close()
	*/
	//Создание маршрутизатора для запросов
	mux := http.NewServeMux()
	setupHandlers(mux)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
	log.SetOutput(logFile)

}
