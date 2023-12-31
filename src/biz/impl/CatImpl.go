package impl

import (
	"fmt"
	"github.com/Mr-ShiWen/MyGo1/src/biz"
)

func init() {
	catImpl := &CatImpl{
		name: "Milly",
	}
	biz.InjectCat(catImpl)
}

type CatImpl struct {
	name string
}

func (c *CatImpl) SayHello() {
	fmt.Println("Hello World, I am Cat ")
}

func (c *CatImpl) GetName() string {
	return c.name
}
