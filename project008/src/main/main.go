// strcut

package main

import "fmt"

func main()  {
	// 空创建
	p1 := per{}
	p1.name = "ca"
	p1.age = 19
	p1.cool = true
	fmt.Println(p1)

	// 给定部分值
	var p2 per = per{cool:true}
	fmt.Println(p2)

	// 给全值
	var p3 per = per{19,"ca",true}
	fmt.Println(p3)

	pp4 := per1{16,"xxy",true}
	p4 := per(pp4)
	fmt.Println(p4)



}

// 定义struct
type per struct {
	age int
	name string
	cool bool
}
// 只有属性名、类型、属性顺序完全一致，两个struct才能转换
type per1 struct {
	age int
	name string
	cool bool
}