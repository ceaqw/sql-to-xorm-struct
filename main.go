package main

import (
	"fmt"
	"sql2struct/unmarshal"
)

func main() {
	err := unmarshal.UnmarshalFile("./test.sql", "falcon_portal")
	if err != nil {
		fmt.Println(err)
	}
}
