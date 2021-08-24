// range遍历

package main

import "fmt"

func main()  {
	li := []int{9,8,0,7,6,5,4,1,2,3}
	for index,num := range li{
		fmt.Printf("index: %d	num: %d\n",index,num)
	}
	// 获取两个值，分别是index、value

	dou := [3][2]int{{1,2},{9,8},{5,7}}
	for x,do := range dou{
		for y,value := range do{
			fmt.Printf("(%d, %d)  %d\n",x,y,value)
		}
	}

}
