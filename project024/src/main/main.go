// 自定义错误

package main

import (
	"errors"
	"fmt"
)

func main()  {

	fmt.Println("start")
	x := test()
	fmt.Println(x)
	fmt.Println("end")

}

func test()int  {
	defer check()
	a := 0
	b := 10
	c := a/b
	if a==0{
		err := errors.New("a is zero")
		panic(err)
	}
	fmt.Println(c)
	return c
}

func check() {
	err := recover()
	if err!=nil{
		fmt.Println(err)
	}
}
// recover 让所在函数不报错
// 即 main完全执行，test执行一半
// 如果 check在main ，则main执行万test结束