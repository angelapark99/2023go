package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//fatal function

func main() {
	fmt.Print("typing:")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')

	if err != nil { //에러발생하면, err=nil 이면 에러 없음
		log.Fatal(err)
	}
	in = strings.TrimSpace(in)
	dan, err := strconv.Atoi(in)
	//	dan, err := strconv.ParseInt(in, 8, 64) // 8진수 - 8*2 = 16
	if err != nil { //에러발생하면
		log.Fatal(err)
	}
	i := 1
	// for i<10{
	// 	fmt.Println(dan, "*", i, "=", (dan * i))
	// 	i++
	// }
	// for i := 1; i < 10; i++ {
	// 	fmt.Println(dan, "*", i, "=", (dan * i))
	// 	fmt.Printf("%d*%d=%d\n", dan, i, (dan * i))
	// }
	// //fmt.Println(dan * 2)

}
