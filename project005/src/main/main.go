// import

// 命名时，开头大写的变量是全局
// 小写的是包内

// 同包不同文件内也不能有同名量/函数

// 必须要有main包，包内有main函数

// 每一个源文件都可以有init
// 先于main，后于定义全局变量
// 优先执行pkg中的init
// 用于初始化


package main

import (
	neo "aa/bb/cc/dd" // 包别名
	"fmt"
	"test"
)

func init()  {
	fmt.Println("init main")
}

func main()  {
	fmt.Println(test.Animal)
	fmt.Println(test.Plant)
	// fmt.Println(test.dont)
	fmt.Println(neo.Xperia)


	// defer 依次压入栈中，返回函数时弹栈

	fmt.Println("defer st")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("defer end")


	// defer中的量输出值保持defer所在语句的结果，一旦入栈则不会变
	x := 193
	defer fmt.Println(x)
	x += 100
}
