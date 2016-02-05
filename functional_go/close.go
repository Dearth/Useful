package main

import (
	"fmt"
	"strconv"
	"flag"
)

type increment func(int) int

func startAt(x int) increment {
	return func(y int) int {
		return x + y
	}
}


func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) <= 1 {
		panic("Not enough input")
	}

	x,_ := strconv.Atoi(args[0])
	y,_ := strconv.Atoi(args[1])

	z := startAt(x)

	fmt.Println(z(y))

}
