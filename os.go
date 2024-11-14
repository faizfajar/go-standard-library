package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	test, err := os.Getwd()

	if err == nil {
		fmt.Println(hostname)
		fmt.Println(test)
	} else {
		fmt.Println("Error", err.Error())
	}
}
