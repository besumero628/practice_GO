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

	// 変数
	var x, y, z int

	var (
		n int
		name string
	)

	i := 1
	b := true
	f := 3.14
	s := "abc"

	n2 := one()

	fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n",x,y,z,n,name,i,b,f,s,n2)

}

func one() int {
	return 1
}