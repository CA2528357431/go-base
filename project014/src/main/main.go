// map删除

package main

import "fmt"

func main()  {
	ver1 := make(map[string]per)
	ver1["good"] = per{"caoan",18}
	ver1["hey"] = per{"hk",199}
	fmt.Println(ver1)

	// 删除key-value
	delete(ver1,"hey")
	fmt.Println(ver1)

	// 删除key不存在则无事发生
	delete(ver1,"asdadadsasd")
	fmt.Println(ver1)
}

type per struct {
	name string
	age int

}
