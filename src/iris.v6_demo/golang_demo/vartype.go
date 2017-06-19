package main

import (
	"fmt"
)

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool // 默认为false
)

var c, d int = 1, 2
var e, f = 123, "hello"

const (
	a1 = 100
	b2
	c3
)

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {
	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
	fmt.Println(a1, b2, c3)
}
