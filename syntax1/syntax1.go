package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var c rune = '가'
	//var a int16 = 7
	//var a = 7
	//a := 7
	//a := 7.3
	//var b float64 = 5.3
	//a = int(b) //Type conversion, casting , 5.3 -> a = 5
	d := false

	fmt.Println(d)
	fmt.Printf("%T\n", d)

	fmt.Println(c)        // //**//unicode 출력
	fmt.Printf("%c\n", c) //one character 출력
	fmt.Printf("%T\n", c) // rune == int32 별명

	fmt.Println(math.Ceil(2.71)) //3
	fmt.Println(math.Floor(2.71)) //2
	fmt.Println(math.Round(2.71)) //3
	fmt.Println("hello go")
	fmt.Println(strings.Title("go git github"))
}
