package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	age	 int
}

type Admin struct {
	User
	title string
}

func main() {
	var ad Admin
	t := reflect.TypeOf(ad)
	printFiled(t,0)


	var adPtr *Admin
	t = reflect.TypeOf(adPtr)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	printFiled(t,0)

}

func printFiled(t reflect.Type, level int) {
	for i,n := 0,t.NumField();i<n;i++ {
		f := t.Field(i)

		if fmt.Sprintf("%s",f.Type) ==  "main.User" {
			printFiled(f.Type,level+1)
		}
		fmt.Println(level,f.Index,f.Name,f.Type,f.Offset)
	}
}
