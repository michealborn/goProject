package main

import "fmt"

func main() {
	//bool,数值,字符串

	//bool略
	var bol bool
	fmt.Println(bol)

	//int共8位，7为表示数值，一位表示正负
	var a int8 = 100
	//uint共8位，全8位表示数值，无负数
	var a2 uint8 = 200
	fmt.Println(a, a2)

	/* 浮点 */
	var f1 float32 = 3.15
	var f2 float64 = 4.15
	//小数位精确：在%f里加入.x，点几就几位
	fmt.Printf("%T,%.1f\n", f1, f1)
	fmt.Printf("%T,%.1f\n", f2, f2)
	fmt.Println(f1)

	/* 字符串 */
	var str1 string = "阿里"
	str2 := `ali2`
	fmt.Println(str1, str2)

	str3 := 'A'
	str4 := "B"
	fmt.Println(str3, str4)

	str5 := "\"ssss\""
	fmt.Println(str5)

	/* 基本数据类型的转换 */
	var vv int8
	vv = 10
	var bb int16
	bb = int16(vv)
	fmt.Println(vv, bb)

	ff1 := 4.83
	var c int
	c = int(f1)
	fmt.Println(ff1, c)

}
