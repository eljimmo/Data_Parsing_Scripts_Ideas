package main
// fmt is a Go package that is used to format basic strings, values, inputs, and outputs
//  implements some I/O functions.
// package for formatting output
import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	content, err := ioutil.ReadFile("factbookcopy.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
