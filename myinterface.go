package main

import (
	"fmt"
	"interfaces/mypkg"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(15.43)
	fmt.Println(value.MethodWithReturnValue())
}