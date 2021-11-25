package main

import (
	"fmt"
	"math"
)

func one() int {
	return 1
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

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

	// math
	fmt.Printf("unit32 max value = %d\n", math.MaxUint32)

	// 配列
	a := [3]int{1,20,300}
	fmt.Printf("%v\n", a)

	// interface{},nil
	var xx interface{}
	fmt.Printf("%#v\n", xx)

	// 関数
	q,r := div(19,7)
	fmt.Printf("商= %d, 余剰= %d\n", q, r)

	// 無名関数 
	fmt.Printf("f= %v\n", func (x, y int) int {return x+y}(2,3))
}