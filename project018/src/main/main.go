// 字符串方法

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {

	example := "hello,world,曹安"

	// 长度
	length := len(example)
	fmt.Println(length)
	// 字母数字一个字节
	// 汉字三个字节


	fmt.Println()


	// 遍历
	// 法一
	for index,value := range example{
		// value获取是字符对应编码
		// 用string()转换
		fmt.Println(index,string(value))
	}
	// 法二\
	// 按字符处理而不是字节
	r := []rune(example)	// 有中文不可直接example[i]，否则中文乱码
	for i:=0;i<len(r);i++{
		fmt.Println(r[i])
	}


	fmt.Println()


	// atoi
	numstr1 := "123"
	num1,err := strconv.Atoi(numstr1)
	if err==nil{
		fmt.Println(num1)
	}else {
		fmt.Println(err)
	}
	// 如果无问题

	fmt.Println()


	// itoa
	numint2 := 321
	num2 := strconv.Itoa(numint2)
	fmt.Println(num2)


	fmt.Println()

	// 改变字符串？
	// 把他改成数组，修改，在改回字符串

	// string 转 字节码
	li := []byte(example)
	fmt.Println(li)
	// 每个汉字三个字节码


	// 字节码 转 string
	neoli := []byte{230,155,185,229,174,137}
	neoexp := string(neoli)
	fmt.Println(neoexp)


	fmt.Println()


	// 含有子串
	check := "hello"
	ju := strings.Contains(example,check)
	fmt.Println(ju)


	fmt.Println()


	// 是否有共有任何一个子串
	cro := strings.ContainsAny(check,example)
	fmt.Println(cro)


	fmt.Println()


	// 子串数量
	sonnum := strings.Count(example,check)
	fmt.Println(sonnum)


	fmt.Println()


	// index首次
	index := strings.Index(example,check)
	fmt.Println(index)


	fmt.Println()


	// index末次
	lastindex := strings.LastIndex(example,check)
	fmt.Println(lastindex)


	fmt.Println()


	// 忽略大小写比较
	cmpexample := "HEllO,worLD,曹安"
	ju1 := strings.EqualFold(cmpexample,example)
	fmt.Println(ju1)



	fmt.Println()


	// 转换大小写
	title := strings.Title(example)
	lower := strings.ToLower(example)
	upper := strings.ToUpper(example)
	fmt.Println(title,lower,upper)


}
