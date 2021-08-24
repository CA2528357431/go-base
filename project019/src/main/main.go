package main

import (
	"fmt"
	"strings"
)

func main()  {

	emp := "  aa  bb  cc  "

	// 除去空格
	both := strings.TrimSpace(emp)
	fmt.Println(both)


	fmt.Println()


	// 删除特定首尾巴
	neo := ".  .r. adc .r. ."
	mid:=strings.Trim(neo," .")
	l:= strings.TrimLeft(neo,". ")
	r:= strings.TrimRight(neo,". ")
	fmt.Println(mid,l,r)
	// 若首尾若干字符出现在把后字符串中，则全删除


	fmt.Println()

	// 判断前后缀
	ori := "hello"
	app1 := "ha"
	app2 := "lo"
	hju := strings.HasPrefix(ori,app1)
	tju := strings.HasSuffix(ori,app2)
	fmt.Println(hju,tju)


	fmt.Println()


	// 删除前后缀
	head := strings.TrimPrefix(ori,app1) // 删除前缀
	tail := strings.TrimSuffix(ori,app2) // 删除后缀
	fmt.Println(head,tail)
	// 必须完全符合才删除，否则不动


	fmt.Println()


	// split
	// 同pysplit
	// 即使连着也切两次
	splitli := strings.Split(emp," ")
	fmt.Println(splitli,len(splitli))


	fmt.Println()


	// splitafter
	// 相当于partition
	partli := strings.SplitAfter(emp," ")
	fmt.Println(partli,len(partli))


	fmt.Println()


	// join
	// 用后者连接前者
	neostr := strings.Join(splitli,",")
	fmt.Println(neostr)


	fmt.Println()


	// repeat
	re := strings.Repeat(ori,3)
	fmt.Println(re)

	// replace
	// 返回将s中前n个不重叠old子串都替换为new的新字符串，如果 n<0或者n>子串数 会替换所有old子串
	replace1 := strings.Replace(ori,"l","fuck",1)
	fmt.Println(replace1)
	replace2 := strings.Replace(ori,"l","fuck",-1)
	fmt.Println(replace2)



}
