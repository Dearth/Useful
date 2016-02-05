package main

import (
	"fmt"
	"runtime"
	"os"
	"strconv"
)


func find_prime(num int, tf chan int) {

	for i := 2; i < num; i++ {

		if num%i == 0 {
			tf <- 0
			return
		}
	}
	tf <- num
}

func find_all_primes(num int, tf chan int) {

	for i := 2; i < num; i++ {
		go find_prime(i, tf)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	tf := make(chan int)

	args := os.Args[1:]

	limit, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error processing number.")
		panic(err)
	}

	go find_all_primes(limit, tf)

	for i := 2; i < limit; i++ {
		num := <-tf
		if num != 0 {
			fmt.Println(num, "is prime.")
		}
	}

	fmt.Printf("Jobs finished.\n")
}


