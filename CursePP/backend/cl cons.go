package main

import (
	"bufio"
	"fmt"

	//"io/ioutil"

	"net/http"
	"os"

	//"github.com/golang-jwt/jwt/v4"

	"encoding/json"
	"strings"

	//"net/url"
	//"crypto/rand"
	//"crypto/rsa"
	//"crypto/x509"
	//"encoding/pem"
	"time"
)

type Tok struct {
	TokenStr string `json:"tokenStr"`
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

var path_1 = "http://localhost:8080/login"
var path_2 = "http://localhost:8080/registration"
var path_3 = "http://localhost:8080/selectByParent"
var path_4 = "http://localhost:8080/children"
var path_5 = "http://localhost:8080/createTask"
var path_6 = "http://localhost:8080/changeTask"
var path_7 = "http://localhost:8080/deleteTask"
var path_8 = "http://localhost:8080/task/findByTask"
var path_9 = "http://localhost:8080/task/addWorker"
var path_10 = "http://localhost:8080/report/left"
var path_11 = "http://localhost:8080/report/forLast"

func ReadUser() (str string) {

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	fmt.Printf("\n--\t%s\t--\n", in.Text())
	str = fmt.Sprintf("%s", in.Text())
	return str

}

func main() {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	fmt.Println("УДАЧНЫЙ ЗАПРОС НА ЛОГИРОВАНИЕ, ВОЗВРАЩАЕТ СТРОКУ ТОКЕНА")

	url := path_1
	method := "POST"
	payload := strings.NewReader(`{
		"login": "Admin",
		"password": "Admin"
	}`)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		//return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return
	}
	defer res.Body.Close()

	var token Tok
	err = json.NewDecoder(res.Body).Decode(&token)
	fmt.Println(token.TokenStr)
	var token_good = token
	fmt.Println(token_good)
	/*
						token.TokenStr = " "
						fmt.Println("НЕ УДАЧНЫЙ ЗАПРОС НА ЛОГИРОВАНИЕ, НЕ ВЕРНЁТ НИЧЕГО")
						url = path_1
						method = "POST"
						payload = strings.NewReader(`{
							"login": "Admin",
							"password": "Admiааn"
						}`)
						req, err = http.NewRequest(method, url, payload)

						if err != nil {
							fmt.Println(err)
							//return
						}
						req.Header.Add("Content-Type", "application/json")

						res, err = client.Do(req)
						if err != nil {
							fmt.Println(err)
							//return
						}
						defer res.Body.Close()

						err = json.NewDecoder(res.Body).Decode(&token)
						fmt.Println(token.TokenStr)


						fmt.Println("ОТВЕТ О НЕ СОЗДАНИИ, ТК ПОВТОРЕНИЕ")
						url = path_2
						method = "POST"
						payload = strings.NewReader(`{
								"login": "Admin",
								"email": "Admin@gmail.com",
								"fio": "Admin",
								"password": "Admin"
							}`)
						req, err = http.NewRequest(method, url, payload)

						if err != nil {
							fmt.Println(err)
							//return
						}
						req.Header.Set("Content-Type", "application/json")
						//req.Header.Add("Authorization", token_good.TokenStr)
						res, err = client.Do(req)
						if err != nil {
							fmt.Println(err)
							//return
						}
						defer res.Body.Close()
						var ans AnswerStr

						err = json.NewDecoder(res.Body).Decode(&ans)
						fmt.Println(ans)
						fmt.Println("ОТВЕТ О СОЗДАНИИ")
						url = path_2
						method = "POST"
						payload = strings.NewReader(`{
								"login": "TEST1",
								"email": "TEST@gmail.com",
								"fio": "TEST TEST",
								"password": "TEST1"
							}`)
						req, err = http.NewRequest(method, url, payload)

						if err != nil {
							fmt.Println(err)
							//return
						}
						req.Header.Set("Content-Type", "application/json")
						//req.Header.Add("Authorization", token_good.TokenStr)
						res, err = client.Do(req)
						if err != nil {
							fmt.Println(err)
							//return
						}
						defer res.Body.Close()
						//var ans AnswerStr
						err = json.NewDecoder(res.Body).Decode(&ans)
						fmt.Println(ans)

						fmt.Println("Поиск информации по детям 0 задачи")
					url = path_3
					method = "POST"
					payload = strings.NewReader(`{
							"pi": 0
						}`)
					req, err = http.NewRequest(method, url, payload)

					if err != nil {
						fmt.Println(err)
						//return
					}
					req.Header.Add("Content-Type", "application/json")
					req.Header.Add("Authorization", token_good.TokenStr)
					res, err = client.Do(req)
					if err != nil {
						fmt.Println(err)
						//return
					}
					defer res.Body.Close()

					task_f := []RespFormTask{}
					err = json.NewDecoder(res.Body).Decode(&task_f)
					if err != nil {
						fmt.Println(err)
						//return
					}
					fmt.Println(task_f)


						fmt.Println("Иерархия по задачам")
				url = path_4
				method = "POST"
				payload = strings.NewReader(`{
							"id": 1
						}`)
				req, err = http.NewRequest(method, url, payload)

				if err != nil {
					fmt.Println(err)
					//return
				}
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", token_good.TokenStr)
				res, err = client.Do(req)
				if err != nil {
					fmt.Println(err)
					//return
				}
				defer res.Body.Close()
				tree := []hierarchy{}

				err = json.NewDecoder(res.Body).Decode(&tree)
				if err != nil {
					fmt.Println(err)
					//return
				}
				fmt.Println(tree)

				fmt.Println("Иерархия по задачам GET")
				url = path_4
				method = "GET"
				req, err = http.NewRequest(method, url, payload)

				if err != nil {
					fmt.Println(err)
					//return
				}
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", token_good.TokenStr)
				res, err = client.Do(req)
				if err != nil {
					fmt.Println(err)
					//return
				}
				defer res.Body.Close()
				tree = []hierarchy{}

				err = json.NewDecoder(res.Body).Decode(&tree)
				if err != nil {
					fmt.Println(err)
					//return
				}
				fmt.Println(tree)

					fmt.Println("ДОБАВЛЕНИЕ ЗАДАЧИ")
				url = path_5
				method = "POST"
				payload = strings.NewReader(`{
						"dateBeg": "2012-12-12",
						"dateEnd": "2012-12-12",
						"parentId": 0,
						"task": "test",
						"mark": "test",
						"status": "test"
					}`)
				req, err = http.NewRequest(method, url, payload)

				if err != nil {
					fmt.Println(err)
					//return
				}
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", token_good.TokenStr)
				res, err = client.Do(req)
				if err != nil {
					fmt.Println(err)
					//return
				}
				defer res.Body.Close()
			var a AnswerStr
				err = json.NewDecoder(res.Body).Decode(&a)
				if err != nil {
					fmt.Println(err)
					//return
				}


					fmt.Println("Изменение ЗАДАЧИ")
				url = path_6
				method = "POST"
				payload = strings.NewReader(`{
						"id": 1,
						"dateBeg": "2012-12-12",
						"dateEnd": "2012-12-12",
						"parentId": 0,
						"task": "test",
						"mark": "test",
						"status": "test"
					}`)
				req, err = http.NewRequest(method, url, payload)

				if err != nil {
					fmt.Println(err)
					//return
				}
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", token_good.TokenStr)
				res, err = client.Do(req)
				if err != nil {
					fmt.Println(err)
					//return
				}
				defer res.Body.Close()
			var a AnswerStr
				err = json.NewDecoder(res.Body).Decode(&a)
				if err != nil {
					fmt.Println(err)
					//return
				}


					fmt.Println("Удвление ЗАДАЧИ")
				url = path_7
				method = "DELETE"
				payload = strings.NewReader(`{
						"pi": 1
					}`)
				req, err = http.NewRequest(method, url, payload)

				if err != nil {
					fmt.Println(err)
					//return
				}
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", token_good.TokenStr)
				res, err = client.Do(req)
				if err != nil {
					fmt.Println(err)
					//return
				}
				defer res.Body.Close()
			var a AnswerStr
				err = json.NewDecoder(res.Body).Decode(&a)
				if err != nil {
					fmt.Println(err)
					//return
				}

					fmt.Println("добавление людей к задаче")
				url = path_9
				method = "POST"
				payload = strings.NewReader(`{
						"idTask": 1,
						"login": "User1"
					}`)
				req, err = http.NewRequest(method, url, payload)

				if err != nil {
					fmt.Println(err)
					//return
				}
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", token_good.TokenStr)
				res, err = client.Do(req)
				if err != nil {
					fmt.Println(err)
					//return
				}
				defer res.Body.Close()
			var a AnswerStr
				err = json.NewDecoder(res.Body).Decode(&a)
				if err != nil {
					fmt.Println(err)
					//return
				}

					fmt.Println("поиск ЗАДАЧ с остатком дней")
		url = path_10
		method = "POST"
		payload = strings.NewReader(`{
				"pi": 7
			}`)
		req, err = http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			//return
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", token_good.TokenStr)
		res, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			//return
		}
		defer res.Body.Close()
		datamass := []LeftStruct{}
		err = json.NewDecoder(res.Body).Decode(&datamass)
		if err != nil {
			fmt.Println(err)
			//return
		}
		fmt.Println(datamass)
	*/

	//var data DistributionStruct2[]DoneForStruct{}
	fmt.Println("поиск ЗАДАЧ выполненнх за последний период")
	url = path_11
	method = "POST"
	payload = strings.NewReader(`{
			"pi": -100
		}`)
	req, err = http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		//return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token_good.TokenStr)
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return
	}
	defer res.Body.Close()
	datamass := []DoneForStruct{}
	err = json.NewDecoder(res.Body).Decode(&datamass)
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println(datamass)
}
