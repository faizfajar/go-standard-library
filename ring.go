package main

import (
	"container/ring"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	for i := 1; i <= data.Len(); i++ {
		if i == 5 {
			data.Value = i
		} else {
			data.Value = "Value " + strconv.Itoa(i)
		}

		data = data.Next()
	}

	//data.Value = "Value 1"
	//
	//data = data.Next()
	//data.Value = "Value 2"
	//
	//data = data.Next()
	//data.Value = "Value 3"
	//
	//data = data.Next()
	//data.Value = "Value 4"
	//
	//data = data.Next()
	//data.Value = "Value 5"

	data.Do(func(value any) {
		fmt.Println(reflect.TypeOf(value))
		if value != nil {
			fmt.Println(value)
		}
	})
}
