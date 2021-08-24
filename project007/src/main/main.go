// 输入

package main

import "fmt"

func main()  {
	var(
		name string
		age int
		judge bool
	)
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Scanln(&judge)
	// 遇到回车终止
	fmt.Println(name+"帅哥",age+2,judge)


	var(
		name1 string
		age1 int
		judge1 bool
	)
	fmt.Scanf("%s %d %t", &name1,&age1,&judge1)
	fmt.Println(name1+"帅哥",age1+2,judge1)




}
