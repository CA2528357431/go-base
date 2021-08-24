// 动态切片

package main

import "fmt"

func main()  {

	li := [10]int{1,2,3,4,5,6,7,8,9,0}

	neosli := li[3:7]

	neosli1 := neosli[:2]
	fmt.Println(len(neosli1),cap(neosli1))
	// 再切片

	neosli2 := neosli[:2]
	fmt.Println(len(neosli2),cap(neosli2))
	// 前抛去

	neosli3 := neosli[:6]
	fmt.Println(len(neosli3),cap(neosli3))
	// 后延申

	// len是当前长度
	// cap是向后延伸到极限的长
	// sli的后延其实是从原数组多取l几个元素
	// sli的前抛其实是在前少取几个

	//对于string切片
	str := "hello,world"
	strsli := str[3:9]
	fmt.Println(strsli)
}