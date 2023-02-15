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
	flag.Parse()

	//check if the user set the flag
	if *caps {
		*msg = strings.ToUpper(*msg)
	}
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}

}
