package one

import "fmt"

func PrintNumber(ptrToNumber interface{}) {
	if ptrToNumber == nil {
		fmt.Println("nil")
		return
	}

	// Проверка типа через type assertion
	if p, ok := ptrToNumber.(*int); ok {
		if p != nil {
			fmt.Println(*p)
		} else {
			fmt.Println("nil")
		}
	} else {
		fmt.Println("invalid type")
	}
}
