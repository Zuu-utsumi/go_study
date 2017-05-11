package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}

	i = "hoge"
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Printf("i = '%s' <type %s>\n", v, t)
}
