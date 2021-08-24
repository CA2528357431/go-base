// 位运算

package main

import "fmt"

func main()  {

	a := 8
	res1 := a&(a-1)
	fmt.Println(res1)
	// 删除二进制末尾的1

	res2 := a&(-a)
	fmt.Println(res2)
	// 获取二进制末尾的1

	a = a<<2
	fmt.Println(a)
	// 左移动位
	a = a>>2
	fmt.Println(a)
	// 右移动位

	b := a^a
	fmt.Println(b)
	// 异或

}
