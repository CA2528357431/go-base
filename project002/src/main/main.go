// 各类<-->string



package main

import (
	"fmt"
	"strconv"
)

func main()  {

	// 不可直接强制转型
	mark := 65
	strmark := string(mark)
	fmt.Println(mark,strmark)


	a := 1
	b := 9.9
	c := true
	d := complex(2,3)


	as := fmt.Sprintf("%d",a)
	bs := fmt.Sprintf("%g",b)   // %f也可，但%g在一定条件下自动科学计数法 %e只用科学计数法
	cs := fmt.Sprintf("%t",c)
	fmt.Println(as,bs,cs)
	// 此法也用于格式输出

	// 格式输出
	aori := fmt.Sprintf("ori: %v",a)
	atype := fmt.Sprintf("type: %T",a)
	fmt.Println(aori,atype)


	// strconv

	// xx-->string
	ass := strconv.FormatInt(int64(a),10) // 后面数字表示进制
	bss := strconv.FormatFloat(b,'g',9,64) // 表示模式(f,g,e)(同格式输出)、精确度、float类型(64/32)
	css := strconv.FormatBool(c)
	dss := strconv.FormatComplex(d,'g',10,64)  // 同float，因为complex的参数是float
	fmt.Println(ass,bss,css,dss)
	// int-->string
	assbeta := strconv.Itoa(a)
	fmt.Println(assbeta)

	// string-->xx
	aba, _ := strconv.ParseInt(ass,10,64)  // 进制，int类型(64/32)
	bba, _ := strconv.ParseFloat(bss,64)  // float类型(64/32)
	cba, _ := strconv.ParseBool(css)
	dba, _ := strconv.ParseComplex(dss,64) // 同float
	fmt.Println(aba,bba,cba,dba)

	// string-->int
	ababeta, _ := strconv.Atoi(ass)
	fmt.Println(ababeta)




}
