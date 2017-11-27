package main

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Item interface {

}

type Fruit struct {
	name string
	amount int
}

type Fruits map[string] int
type Vegetables map[string] int

type Data struct {
	Fruit Fruits
	Verggies Vegetables
}

type Payload struct {
	Stuff Data
}

type Request struct {
	Item map[string] interface{}
}

var vegetables map[string] int = make(map[string] int)
var fruits map[string] int = make(map[string] int)


func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))
}

func postFruit(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

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

func main() {
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

	http.HandleFunc("/fruit", postFruit)
	http.HandleFunc("/", serveRest)

	http.ListenAndServe(host, nil)
}

func getJsonResponse() ([]byte, error){
	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
