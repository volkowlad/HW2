package five

import (
	"fmt"
	"strings"
)

func Five() {
	// first
	str := "Привет"
	sRune := []rune(str)
	sRune[2] = 'е'
	fmt.Println(string(sRune))

	// second
	s := strings.Replace(str, "и", "е", 1)
	fmt.Println(s)
}
