package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Record struct {
	Userid string
	Age    int `json:"age"`
	Height int `json:"height"`
	Income int `json:"income"`
}

func main() {

	data, err := ioutil.ReadFile("input.json")
	if err != nil {
		log.Fatal(err)
	}
	nread := len(data)
	fmt.Println("Bytes Read: ", nread)
	fmt.Printf("%v", data)
	fmt.Println()
	fmt.Printf("%s", data)

	var sr []Record

	err = json.Unmarshal(data, &sr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sr")
	fmt.Print(sr)
	fmt.Println()

	out, err := json.Marshal(sr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", out)

}
