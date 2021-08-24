// time

package main

import (
	"fmt"
	"time"
)

func main()  {

	// 当前时间
	now := time.Now()
	fmt.Println(now)
	// time.Time类型

	// 通过参数获取时间数据
	fmt.Println(now.Year(),
		now.Month(),
		now.Day(),
		now.Weekday())

	// 标准时间
	fmt.Println(now.UTC())

	// 地方时间
	fmt.Println(now.Local())

	// 时区
	fmt.Println(now.Zone())

	// 判断相等,before,after
	fmt.Println(now.UTC().Equal(now.Local()),
		now.UTC().Before(now.Local()),
		now.UTC().After(now.Local()),)

	// 格式化
	mytime := now.Format("2006-01-02, 15H 04M 05S")
	// 其中2006、15等数字不可改，但顺序格式自由设置
	fmt.Println(mytime)

	// 设定时间戳
	unix := time.Unix(2000,19)
	fmt.Println(unix)

	// 对象时间戳
	unixnow := now.Unix()
	fmt.Println(unixnow)





}
