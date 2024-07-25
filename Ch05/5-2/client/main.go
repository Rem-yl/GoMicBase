package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Msg string `json:"msg"`
}

func main() {
	r, err := http.Get("http://localhost:8080/hello?name=rem")
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var res Result
	err = json.Unmarshal(b, &res)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Msg)
}
