// 错误处理

package main

import "fmt"

func main()  {
	fmt.Println("start")
	x := test() // 出错则为默认值
	fmt.Println(x)
	fmt.Println("end")
}

func test() int {

	defer check()


	a := 0
	b := 10
	c := b/a
	fmt.Println(c)
	return c
}

// 处理
func check() {
	err := recover()
	if err!=nil{
		fmt.Println(err)
	}
}