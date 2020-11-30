package model

import (
	"encoding/json"
)

type Post struct {
	Id   int
	Data string
}

func encode(p Post) string {
	res, err := json.Marshal(p)
	if err != nil {
		panic(err.Error())
	}
	return string(res)
}

func decode(s string) Post {
	var p Post
	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		panic(err.Error())
	}
	return p
}

/*
func main() {
	p1 := Post{123, "hello"}
	res := encode(p1)
	fmt.Println(res)

	p2 := decode(res)
	fmt.Println(p2)
}
*/
