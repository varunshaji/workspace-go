package main

import "fmt"
import "strconv"

var test int
var x int
var y int
var z string
var w string

func main() {
	w = "33"
	x = 33
	y = 21
	z = "Go is fun"
	fmt.Printf("%v + %v = %v \n", x, y, x+y)
	fmt.Printf("%d + %d = %s \n", x, y, x+y)
	fmt.Println(string(x) + z)
	fmt.Printf("This '%s' is actually a string \n", strconv.Itoa(x))
	val, _ := strconv.Atoi(w)
	fmt.Printf("This %d is actually a decimal \n", val)
	fmt.Println(test)
}
