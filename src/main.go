package main

import (
	"awesomeProject/src/biz"
	_ "awesomeProject/src/biz/impl"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	biz.RefCat().SayHello()
	catName := biz.RefCat().GetName()
	fmt.Println(catName)
}
