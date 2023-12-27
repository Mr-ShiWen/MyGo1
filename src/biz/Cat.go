package biz

type ICat interface {
	SayHello()
	GetName() string
}

var catImpl ICat

func RefCat() ICat {
	return catImpl
}

func InjectCat(impl ICat) {
	catImpl = impl
}
