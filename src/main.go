package main

import (
	"fmt"
	"github.com/Mr-ShiWen/MyGo1/src/biz"
	_ "github.com/Mr-ShiWen/MyGo1/src/biz/impl"
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
