package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func httpget() {
	fmt.Println("httpget")

	resp, err := http.Get("https://fgtn9f7u1j.execute-api.us-east-2.amazonaws.com/dev/compare-yourself/all")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	fmt.Printf("body %v", body)
	var rs []Record
	err = json.Unmarshal(body, &rs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)

}
