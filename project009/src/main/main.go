// 数组与切片

package main

import "fmt"

func main()  {

	// 数组

	li1 := [3]int {3,4,1}
	fmt.Println(li1)
	// 标准构造

	li2 := [3]int {3,4}
	fmt.Println(li2)
	// 缺省构造

	var li3 [3]int
	li3[0],li3[2]=9,6
	fmt.Println(li3)
	// 声明构造


	// 切片

	li := [10]int {1,2,3,4,5,6,7,8,9,0}

	sli1 := li[1:5]
	fmt.Println(sli1)
	var sli2 = li[3:8] // var sli2 []int = li[3:8] 可省略类型
	fmt.Println(sli2)
	// 两种有源构造

	sli3 := li[:3]
	sli4 := li[6:]
	// 如果切片 从开头开始/在结尾结束 则可省略
	fmt.Println(sli3)
	fmt.Println(sli4)

	chasl1 := li[:6]
	fmt.Println(len(chasl1),cap(chasl1))
	chasl2 := li[3:6]
	fmt.Println(len(chasl2),cap(chasl2))
	// len长度指的是含有元素数
	// cap容量指的是可容元素数量，是原数组除去头部之前的长

	sli1[3] = 999999
	fmt.Println(li)
	// 切片只是观测方式、一个指向原数组的指针集
	// 修改slice中的元素实质上会修改原数组的数据

	// 无源构造
	// 利用切片创造不定长数组
	sli := []int{1,2,3,4,5}
	fmt.Println(sli)
	// 本质上，创建了[1,2,3,4]数组，又构造[:]切片

	dou := [3][9]int{}
	fmt.Println(dou)
	// 二维数组

	douli := [2][3]int{{1,2,3},{3,4,5}}
	fmt.Println(douli)

}