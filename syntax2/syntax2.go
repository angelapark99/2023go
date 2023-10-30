package main

import (
	"fmt"
	"strings"
)

func main() {
	texts := "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)

}
