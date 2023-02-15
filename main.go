package main

import (
	"flag"
	"fmt"
	"strings"
)

//demonstrate flags

func main() {
	//set the flags
	msg := flag.String("msg", "HOWDY, STRANGER", "the message to display")
	num := flag.Int("num", 1, "How many times to print the message")
	caps := flag.Bool("U", false, "Should the string be all caps")
	reverse := flag.Bool("r", false, "reverse the string (true) ")
	flag.Parse()

	//check if the user set the flag
	if *caps {
		*msg = strings.ToUpper(*msg)
	}
///check if we should reverse the string
if *reverse {
	//reverse string
	var result string
	for _, value := range *msg {
		result = string(value) + result
	}
	*msg = result
}

	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}

}
