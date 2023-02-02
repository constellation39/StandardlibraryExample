package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type R struct {
	a string `json:"a"`
}

type ResultTestCards struct {
	All     map[string]map[string]int32
	Cards   []int32   `json:"cards"`
	Players [][]int32 `json:"players"`
	Num     int32     `json:"num"`
}

func main() {
	r := new(ResultTestCards)
	err := json.Unmarshal([]byte("{\"all\":{\"1\":{\"11\":4,\"12\":0,\"13\":3,\"14\":0,\"15\":0,\"16\":0,\"17\":4,\"18\":3,\"19\":4},\"2\":{\"21\":4,\"22\":4,\"23\":2,\"24\":0,\"25\":0,\"26\":0,\"27\":3,\"28\":0,\"29\":3},\"3\":{\"31\":4,\"32\":4,\"33\":4,\"34\":4,\"35\":4,\"36\":4,\"37\":4,\"38\":4,\"39\":4}},\"cards\":[\"28\",\"29\",\"25\",\"25\",\"15\",\"14\",\"14\",\"13\",\"12\",\"12\",\"12\",\"12\"],\"players\":[[15,15,15,16,16,16,16,25,25,26,26,26,26],[14,14,18,23,23,24,24,24,24,27,28,28,28]],\"num\":13}"), r)
	if err != nil {
		log.Println("1111 ", err)
	}
	fmt.Println("ok")
}
