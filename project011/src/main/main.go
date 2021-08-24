// slice实现动态数组

package main

import "fmt"

func main()  {

	// 创建
	a := make([]int, 6)  // 指定len=cap=5
	a[0] = 9999
	fmt.Printf("%s len=%d cap=%d %v\n", "a", len(a), cap(a),a)
	b := make([]int, 1, 6)  // 指定len=1 cap=5
	b[0] = 9999
	fmt.Printf("%s len=%d cap=%d %v\n", "b", len(b), cap(b),b)

	// 追加
	bapp := append(b,19) // 追加单个
	fmt.Printf("%s len=%d cap=%d %v\n", "bapp", len(bapp), cap(bapp),bapp)
	bapp = append(b,19,6,-9) // 追加多个
	fmt.Printf("%s len=%d cap=%d %v\n", "bapp", len(bapp), cap(bapp),bapp)

	// 如果长度大于原cap，cap*2
	aapp := append(a,19)
	fmt.Printf("%s len=%d cap=%d %v\n", "bapp", len(aapp), cap(aapp),aapp)

	// 拷贝
	// 将后者拷贝给前者

	// 如果c1小，则只把c1填满，后面不要
	c1 := make([]int,5)
	copy(c1,aapp )
	fmt.Println(c1)
	// 如果c2大，则把原版全复制之后，后面不管
	c2 := make([]int,20)
	copy(c2, aapp)
	fmt.Println(c2)



}
