package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func apiHandler(w http.ResponseWriter, req *http.Request) {
	middleware(w, req)
	if req.Method == "GET" {
		//http.ServeFile(w, req, "form.html")
		w.Write([]byte("Hello, world!"))
		fmt.Printf("hey")
		//fmt.Fprintf(w, "\nReceived time: %v\tType:GET\tRoute:/hello\n", time)
	}

	//fmt.Fprintf(w, "Hello, world!\n")
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

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", middleware)
	mux.HandleFunc("/data", dataHandler)
	mux.HandleFunc("/hello", apiHandler)
	//mux.HandleFunc("/hello", apiHandler)
	//middleware mux.Handle("when",func)

}
func middleware(w http.ResponseWriter, req *http.Request) {
	time := time.Now().Format(time.ANSIC)
	fmt.Printf("\nExecuting middleware at: %v Method:%s URL:%s\n", time, req.Method, req.URL.Path)
	fmt.Print("\nExecuting middleware\n")
}

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
