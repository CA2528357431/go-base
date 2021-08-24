// 指针

package main

import "fmt"

func main()  {

	// 地址
	var x int = 3;
	fmt.Println(&x)

	// 获取指针
	var id *int = &x
	par := &x
	fmt.Println(id)
	fmt.Println(par)

	// 指针背后的值
	back := *par
	fmt.Println(back)

	// 间接修改
	*par+=1
	fmt.Println(x)


	// int float string bool struct 数组 是值类型
	// 指针 slice map channel interface 是引用类型，都相当于指针
}
