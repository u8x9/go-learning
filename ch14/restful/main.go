package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var employeeDB map[string]*Employee

func init() {
	employeeDB = map[string]*Employee{
		"Mike": &Employee{"E1", "Mike", 35},
		"Rose": &Employee{"E2", "Rose", 45},
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome")
}
func GetEmployeeByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	e, ok := employeeDB[name]
	if !ok {
		fmt.Fprintf(w, "The employee %q is not exists.\n")
		return
	}
	buf, err := json.Marshal(e)
	if err != nil {
		fmt.Fprintln(w, "Marshal failed:", err.Error())
		return
	}
	w.Write(buf)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/employee/:name", GetEmployeeByName)
	log.Fatal(http.ListenAndServe(":8080", router))
}
