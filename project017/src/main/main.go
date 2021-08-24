package main

import "fmt"

// int float string bool struct 数组 是值类型
// 指针 slice map channel interface 是引用类型，都相当于指针




func changelist(x [10]int)  {
	x[0]+=10
	fmt.Println("func",x)
}
func changeslice(x []int)  {
	x[0]+=10
	fmt.Println("func",x)
}
func changelistori(x *[10]int)  {
	x[0]+=10
	fmt.Println("func",*x)
}

func main()  {

	sli := make([]int,10)
	changeslice(sli)
	fmt.Println("main",sli)

	li := [10]int{}
	changelist(li)
	fmt.Println("main",li)

	ori := [10]int{}
	changelistori(&ori)
	fmt.Println("main",ori)


// 引值型是修改拷贝来的值
// 引用型是直接修改原值
// 引值型是修改？ 传入地址
}