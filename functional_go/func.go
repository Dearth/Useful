package main

import (
	"fmt"
	"flag"
	"strconv"
)

type math func(int, int) int

func add(x,y int) int {
	return x+y
}

func mult(x,y int) int {
	return x*y
}

func addOrMult(x,y int) math {

	if x > y {
		return add
	} else {
		return mult
	}
}

func main() {

	flag.Parse()
	args := flag.Args();

	x,_ := strconv.Atoi(args[0])
	y,_ := strconv.Atoi(args[1])

	z := addOrMult(x,y)(x,y)

	fmt.Println(z)
}
