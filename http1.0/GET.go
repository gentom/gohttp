package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // read Body's content as a byte sequence
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}