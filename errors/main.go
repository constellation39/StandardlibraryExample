package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type R struct {
	a string `json:"a"`
}

func main() {
	r := new(R)
	err := json.Unmarshal([]byte("{\"12312321312pu1409eu\":11}"), r)
	if err != nil {
		log.Println("1111 ", err)
	}
	fmt.Println("ok")
}
