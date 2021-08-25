// 不定参数函数

package main

import "fmt"

func main()  {
	x := do(3,1,2,3)
	fmt.Println(x)
}

func do(x int,item...int) int {
	sum := 0
	for _,num := range item{
		sum+=num
	}
	return sum*x
}
