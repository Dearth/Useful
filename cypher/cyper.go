package main

import (
	"os"
	"flag"
	"io"
)

func checkErr(e error){
	if e != nil {
		panic(e)
	}
}

func XORdata(m, k, c *os.File) (n int) {
	msgt := make([]byte, 1)
	keyt := make([]byte, 1)
	cypt := make([]byte, 1)

	_, err := m.Read(msgt)
	if err != nil {
		if err == io.EOF {
			return 1
		}else {
			checkErr(err)
		}
	}

	_, err = k.Read(keyt)
	if err != nil {
		if err == io.EOF {
			_, err := k.Seek(0,0)
			checkErr(err)
			_, err = k.Read(keyt)
			checkErr(err)
		} else {
			checkErr(err)
		}
	}

	cypt[0] = msgt[0] ^ keyt[0]

	_, err = c.Write(cypt)
	checkErr(err)

	return 0
}


func main() {

	msgPtr := flag.String("msg", "msg.txt", "message to be encrypted")
	keyPtr := flag.String("key", "key.txt", "key text to be used for encryption")
	cypPtr := flag.String("cyp", "out.txt", "output name of the encypted text")
	decrypt := flag.Bool("d", false, "Decrypt out.txt with key.txt")

	flag.Parse()

	if *decrypt {
		msg, err := os.Create(*msgPtr)
		checkErr(err)

		key, err := os.Open(*keyPtr)
		checkErr(err)

		cyp, err := os.Open(*cypPtr)
		checkErr(err)

		for i := 0; i != 1; i = XORdata(cyp, key, msg) {}
	} else {
		msg, err := os.Open(*msgPtr)
		checkErr(err)

		key, err := os.Open(*keyPtr)
		checkErr(err)

		cyp, err := os.Create(*cypPtr)
		checkErr(err)

		for i := 0; i != 1; i = XORdata(msg, key, cyp) {}
	}
}
