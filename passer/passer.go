package main

import (
	"github.com/dchest/uniuri"
	"os"
)

func main () {

	fo, err := os.Create("passwords.txt")
	if err != nil {
		panic(err)
	}

	defer fo.Close()

	for i:= 0; i < 20; i++ {
		str := uniuri.NewLen(20)
		
		if _, err := fo.WriteString(str); err != nil {
			panic(err)
		}

		if _, err := fo.WriteString("\n"); err != nil {
			panic(err)
		}
	}
}
