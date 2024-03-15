package main

import (
	"crypto/rand"
	functions "finYP/utils/func"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func err_message() {
	fmt.Printf("\nHi...Welcome to Random File Generator program...\n")
	fmt.Printf("format: $ go build randFileGen.go\n")
	fmt.Printf("\t$ ./randFileGen PATH_TO_FILENAME\n\n")
	os.Exit(8)
}

func main() {
	if len(os.Args) < 2 {
		err_message()
	}

	//c := 10
	c, err := strconv.Atoi(os.Args[1])
	functions.CheckError(err)

	b := make([]byte, c)
	_, err = rand.Read(b)
	functions.CheckError(err)

	ioutil.WriteFile("../../input.test", b, 0666)

}
