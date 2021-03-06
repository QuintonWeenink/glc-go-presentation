package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"github.com/quintonweenink/glc-go-presentation/items"
)

// In mem DB
var vegetables = make(map[string]int)
var fruits = make(map[string]int)

//Payload used to contain requests
type Payload struct {
	Stuff interface{}
}

func restBase(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Fruit    interface{}
		Verggies interface{}
	}

	d := data{fruits, vegetables}
	p := Payload{d}

	response, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))
}

func postFruit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	type Request struct {
		Item map[string]interface{}
	}

	var p Request
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	name, ok := p.Item["name"].(string)
	if !ok {
		panic("Name should be string")
	}
	amount, ok := p.Item["amount"].(float64)
	if !ok {
		panic("Amount should be int")
	}
	fruits[name] = int(amount)

	fmt.Fprintf(w, string("Your fruit has been added"))
}

func getFruit(w http.ResponseWriter) {
	type Response struct {
		Item interface{}
	}

	d := Response{fruits}
	p := Payload{d}

	response, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))
}

func restFruit(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		postFruit(w, r)
	case "GET":
		getFruit(w)
	default:
		fmt.Fprintf(w, string("Method not mapped"))
	}
}

func main() {
	// Initialize in mem DB
	vegetables["Carrots"] = 21
	vegetables["Peppers"] = 0
	fruits["Apples"] = 25
	fruits["Peppers"] = 11

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	host := "localhost:" + port

	fmt.Println("Serving on " + host + "..")

	http.HandleFunc("/fruit", restFruit)
	http.HandleFunc("/", restBase)

	http.ListenAndServe(host, nil)
}
