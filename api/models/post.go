package main

import (
	"encoding/json"
	"fmt"
)

type Post struct {
	Id   int
	Data string
}

func main() {
	p1 := Post{123, "hello"}
	res, err := json.Marshal(p1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(res))

	var p2 Post
	err = json.Unmarshal(res, &p2)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(p2)
}
