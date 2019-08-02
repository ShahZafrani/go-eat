package main

import "fmt"

var greeting = "howdy"
const hello string = "hello"
const unused = "I'm declared but not used"

func main() {
  greeting += " y'all"
  hello += " world"
  fmt.Println(greeting)
  fmt.Println(hello)
}
