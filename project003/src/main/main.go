// if for switch

package main

import (
	"fmt"
	"time"
)

func main()  {
	for i:=0;i<2;i++{
		fmt.Println(i)
	}
	fmt.Println()

	r := 0
	for{
		fmt.Println(r)
		r++
		if r==2{
			break
		}
	}
	fmt.Println()

	// if创建临时变量
	s := 0
	for{
		fmt.Println(s)
		s++
		if ss :=s*2;ss==4{
			break
		}

	}
	fmt.Println()

	t := 9
	if t>10{
		fmt.Println("a")
	}else if t>5 {
		fmt.Println("b")
	}else {
		fmt.Println("c")
	}
	fmt.Println()


	// switch 的 case 无需为常量，可以是任何类型包括函数
	// 如果是函数则执行后取返回值
	// case的数据类型和switch一致
	u := 9.0
	ex1 := 3.0
	ex2 := 5.0
	switch u {
	case ex1:
		fmt.Println("a")
	case ex2:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
	fmt.Println()

	// switch创建临时变量
	switch r := u/2;r {
	case ex1:
		fmt.Println("a")
	case ex2:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
	fmt.Println()

	// 无条件switch
	// 更清晰的ifelse
	timing := time.Now()
	switch {
	case timing.Hour() < 12:
		fmt.Println("Good morning!")
	case timing.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	fmt.Println()


}
