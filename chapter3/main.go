package main

import (
	"fmt"
)

func main()  {
	// Println
	fmt.Println("test")

	// Printf
	fmt.Printf("数値 = %d\n", 5)

	// %v, %T
	fmt.Printf("数値 = %v, 言語 = %v\n", 5, "文字列")
	fmt.Printf("数値 = %T, 言語 = %T\n", 5, "文字列")
}