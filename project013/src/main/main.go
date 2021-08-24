// 映射
// （另有make分析）

// 类似于py中的dic

package main

import "fmt"

func main()  {
	// 定义与赋值
	ver1 := make(map[string]per)
	ver1["good"] = per{"caoan",18}
	ver1["hey"] = per{"hk",199}
	fmt.Println(ver1)

	fmt.Println()

	// 另一定义
	var ver2 map[int]per
	ver2 = make(map[int]per)
	ver2[33] = per{"caoan",18}
	fmt.Println(ver2)

	fmt.Println()

	// 静态定义
	ver3 := map[string]per{
		"yep": per{
			"ca",19,
		},
		// 类型名可省略
		"nope": {
			"jk",16,
		},

	}
	fmt.Println(ver3)
	// 对于make的分析
	// make是引用型数据的初始化方法
	// 采用静态定义，则已经手动初始化，不用make
	// 常规定义则要make初始化


	// 遍历
	for key,value := range ver1{
		fmt.Println(key,value)
	}

	// 切片


}

type per struct {
	name string
	age int
}


