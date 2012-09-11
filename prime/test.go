package main

import (
	"fmt"
)


func find_prime(num int) {
	
	for i := 2; i < num; i++ {
		
		if num%i == 0 {
			return
		}
	}

	fmt.Printf("%d is prime\n", num)
}

func find_all_primes(num int){
	
	for i := 2; i < num; i++ {
		go find_prime(i)
	}
}

func main() {
 	
	test := 100000

	find_all_primes(test)

	fmt.Printf("Jobs finished.\n")
}


