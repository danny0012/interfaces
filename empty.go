package main

import _ "fmt"

func AcceptAnything(thing interface{}) {

}

func main() {
	AcceptAnything(3.14159)
	AcceptAnything("A string")
	AcceptAnything(true)
}