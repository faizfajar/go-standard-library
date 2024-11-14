package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line\nanjayy faiz")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		// lines, err := reader.Read()

		if err == io.EOF {
			fmt.Println(err)
			break
		}

		fmt.Println(string(line))
		// fmt.Println(string(lines))
	}
}
