// builtin

package main

import "fmt"

func main()  {

	// new
	// 获取的是指针
	// 可对于各种类型
	nor := new(int)
	fmt.Printf("type: %T\n",nor)
	*nor +=1
	fmt.Printf("type: %T\n",*nor)

	// make
	// 获取对象本身
	// 对于引用类型
	li := make([]int,10)
	fmt.Printf("type: %T\n",li)

	//

}
