// 闭包

// hey,别把函数看的太高
// 它也可以是量
// 将fuc当作变量

package main

import "fmt"

// 返回func
// 本例中每个myfunc都与x绑定
func createfunc(x int) func(int,int)int {
	myfunc := func(a int,b int) int{
		add := a+b
		res := add*x
		return res
	}
	return myfunc
}
// func(int,int)int 是类型


// 将func做参数
func usefunc(a int,b int, myfunc func(int,int)int)int{
	return myfunc(a,b)
}
// func(int,int)int 是类型


func main()  {
	myfunc := createfunc(3)
	res := usefunc(3,5,myfunc)
	fmt.Println(res)
}