package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>Hello Kubernetes!!!</h1>"))
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s.", name, age)
}
func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/myTasks/task.txt")
	if err != nil {
		log.Fatalf("Error reading file", err)
	}
	fmt.Fprintf(w, "Todo: %s.", string(data))
}
