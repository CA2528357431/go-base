// 调整/自定time

package main

import (
	"fmt"
	"time"
)

func main()  {

	dur := time.Second*10+time.Hour*3+time.Minute
	// 时间段

	now := time.Now()
	fmt.Println(now)
	neo := now.Add(dur)
	fmt.Println(neo)
	// 改变时间


	// sleep
	temp := time.Millisecond*200
	for _,value := range []int{1,2,3,4}{
		fmt.Println(value)
		time.Sleep(temp)
	}

	// 定义time
	time1, _ := time.Parse("2006-01-02-15-04-05","2020-09-01-10-00-55")
	fmt.Println(time1)

	time2 := time.Unix(1251145125,123967)
	fmt.Println(time2)

	time3 := time.Date(2020,9,3,1,44,32,1000098,time.Local)
	fmt.Println(time3)
}
