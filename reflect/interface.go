package main

import (
	"fmt"
	"reflect"
)

type User1 struct {
	Username string
	age      int
}

func main() {
	u := User1{"Jack", 23}
	v := reflect.ValueOf(u)
	fmt.Println(v.FieldByName("Username").Interface())
	fmt.Println(v.FieldByName("Username"))
	//fmt.Println(v.FieldByName("age").Interface())
	fmt.Println(v.FieldByName("age"))
}
