package main

// 联合import
import (
	"fmt"
	"math"
)

var na int = 996
// 函数外只用此定义


func main(){

	fmt.Println("hello,world")

	fmt.Println(math.Pi)

	fmt.Println("abcdefghijklmnopqrstuvwxyz",
		"abcdefghijklmnopqrstuvwxyz",
		"abcdefghijklmnopqrstuvwxyz")

	var a int = 19
	var b = 78
	// 类型往往可以省略
	c := a*b+na
	// 函数内两种定义

	fmt.Println(c)



	d,e := 13,76
	fmt.Println(d,e)
	d,e = e,d
	fmt.Println(d,e)
	// 多变量定义\赋值

	var (
		x1 int = 9
		x2 = 10
		// 类型往往可以省略
	)
	fmt.Println(x1,x2)
	// 多变量定义

	const op int = 19
	fmt.Println(op)
	// 单常量

	const (
		op1 int = 6
		op2 int = 3
	)
	fmt.Println(op1,op2)
	// 多常量

	// 基本类型有
	// bool string int uint float complex

	i := complex(1,2)
	fmt.Println(i)
	// 虚数


	// 默认值
	// bool false
	// string ""
	// 数字 0
	var (
		aa int
		bb bool
		cc string
		dd complex64
		ee float64
	)
	fmt.Println(aa,bb,cc,dd,ee)



	vapor1 := 8.6
	vapor2 := 3
	line := vapor1 + float64(vapor2)
	fmt.Println(line)
	// 必须要显示类型转换

	// 加减乘除摩和c一样
	fmt.Println(10/4)
	fmt.Println(10.0/4)

	// 支持字符串加减
	s1 := "aa"
	s2 := "bb"
	s3 := s1+s2
	fmt.Println(s3)

	// 逻辑、比较同c
	//  || 与 | 区别 (&&与&同理)
	//  如a||b
	//  a对，a||b不执行b，而a|b还要执行b
	judge := (a>b) || (c>d)
	fmt.Println(judge)


}
