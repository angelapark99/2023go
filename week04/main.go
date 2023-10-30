package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	var a int = 9
	var b float32
	var c, d string
	var e float32 = 3.14
	b = 2.7
	//a = 9
	c = "inha"
	d = "go..."
	h := 'Z'
	fmt.Println(float64(a) * float64(b))
	fmt.Println(reflect.TypeOf(h))
	fmt.Println(reflect.TypeOf(b)) //float64
	fmt.Println(a, b, c, d, e, h)
	fmt.Println(reflect.TypeOf('B'))
	//fmt.Println(math.Floor("inha"))
	fmt.Println('k', 'a', 'n')
	fmt.Println(math.Floor(2.15))
	fmt.Println(strings.Title("open source~\n\"go\""))
}
