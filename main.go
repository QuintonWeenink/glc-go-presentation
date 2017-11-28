package main

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"

	// "github.com/quintonweenink/glc-go-presentation/items"
)

// In mem DB
type Fruits map[string] int
type Vegetables map[string] int
var vegetables map[string] int = make(map[string] int)
var fruits map[string] int = make(map[string] int)

type Data struct {
	Fruit Fruits
	Verggies Vegetables
}

type Payload struct {
	Stuff interface{}
}

type Request struct {
	Item map[string] interface{}
}

type Response struct {
	Item interface{}
}

func restBase(w http.ResponseWriter, r *http.Request) {
	d := Data{fruits, vegetables}
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
