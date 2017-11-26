package main

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Fruits map[string] int
type Vegetables map[string] int

type Data struct {
	Fruit Fruits
	Verggies Vegetables
}

type Payload struct {
	Stuff Data
}



func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))
}

func serveRestGopher(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var p Payload

	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)

	fmt.Fprintf(w, string(response))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := "localhost:" + port

	fmt.Println("Serving on " + host + "..")

	http.HandleFunc("/gopher", serveRestGopher)
	http.HandleFunc("/", serveRest)

	http.ListenAndServe(host, nil)
}


func getJsonResponse() ([]byte, error){
	fruits := make(map[string] int)
	fruits["Apples"] = 25
	fruits["Peppers"] = 11

	vegetables := make(map[string] int)
	vegetables["Carrots"] = 21
	vegetables["Peppers"] = 0

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
