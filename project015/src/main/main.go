// 匿名函数

// 有些函数只用一次

package main

import "fmt"

// 全局匿名函数
var (
	itfun func(int,int)int = func(a int,b int) int{
		res := a+b*3
		return res
	}
)


func main()  {
	res := func(a int,b int) int{
		res := a+b*3
		return res
	}(3,6)
	fmt.Println(res)
	// 大括号内是函数
	// 后面小括号是调用



	// 用变量形式保存函数
	myfunc := func(a int,b int) int{
		res := a+b*3
		return res
	}
	neo := myfunc(13,6)
	fmt.Println(neo)


	// 全局匿名函数
	it := itfun(4,6)
	fmt.Println(it)



}