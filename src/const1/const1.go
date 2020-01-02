package main

import "fmt"

func main() {
	/* 常量定义，不能改 */
	const A string = "a"
	const B = "b"
	fmt.Println(A, B)
	//常量不使用也可以编译
	const C = 3.141562653
	//一堆常量
	const A1, A2, A3 = 1, 2.3, "aaa"
	fmt.Println(A1, A2, A3)
	/* 如果定义的常量没有赋值，默认和上一个一样，如果一开始就没赋值，则编译不通过 */
	const (
		B1 = 11
		B2 = 22.2
		B3 = "bbb"
		B4
	)
	fmt.Println(B1, B2, B3, B4)

	/* 枚举类型 */
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
