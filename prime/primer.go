package main

import (
	"fmt"
	"math/big"
	"os"
	"runtime"
)


//NEXT
func NEXT(n, i *big.Int) *big.Int {
	temp_n := new (big.Int)
	temp_n.Set(n)

	return temp_n.Rsh(temp_n.Add(n, temp_n.Div(i, n)),1)
}

func bigIntSqrt(number *big.Int) *big.Int {
	temp := new (big.Int)
	one := new (big.Int)
	one.SetInt64(1)

	n := new (big.Int)
	n.SetInt64(1)

	n1 := new (big.Int)
	n1 = NEXT(n, number)

	for temp.Abs(temp.Sub(n1, n)).Cmp(one) == 1 {
		n.Set(n1)
		n1 = NEXT(n, number)
	}

	for temp.Mul(n1, n1).Cmp(number) == 1 {
		n1.Sub(n1, one)
	}

	return n1
}

func isPrime(p, i *big.Int, tf chan bool) {
	temp := new (big.Int)
	zero := new (big.Int)
	zero.SetInt64(0)

	if i.ProbablyPrime(10) {
		if temp.Mod(p, i).Cmp(zero) == 0 {
			tf <- false
		} else {
			tf <- true
		}
	} else {
		tf <- true;
	}

	runtime.Goexit()
}

func spawner(pos_prime, sqrt *big.Int, tf chan bool) {
	base := new (big.Int)
	increment := new (big.Int)
	increment.SetInt64(2)

	for base.SetInt64(3); base.Cmp(sqrt) < 0; base.Add(base,increment) {
		temp := new (big.Int)
		temp.Set(base)

		go isPrime(pos_prime, temp, tf)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	args := os.Args[1:]

	pos_prime := new (big.Int)
	temp := new (big.Int)
	base := new (big.Int)
	square := new (big.Int)
	zero := new (big.Int)
	increment := new (big.Int)
	tf := make(chan bool)

	is_prime := true

	pos_prime.SetString(args[0], 10)
	increment.SetInt64(2)
	base.SetInt64(2)
	zero.SetInt64(0)
	square = bigIntSqrt(pos_prime)

	if pos_prime.Cmp(base) == 0 {
		is_prime = true
	} else if temp.Mod(pos_prime, base).Cmp(zero) == 0 {
		is_prime = false
	} else {

		go spawner(pos_prime, square, tf)

		for base.SetInt64(3); base.Cmp(square) < 0; base.Add(base,increment) {
			is_prime = is_prime && <-tf
		}
	}

	if is_prime {
		fmt.Println(pos_prime, "is prime.")
	} else {
		fmt.Println(pos_prime, "is not prime.")
	}
}
