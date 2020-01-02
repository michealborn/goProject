package main

import "fmt"

func main() {
	/* iota 特殊的常量，可以被编译器自动修改的常量
	每定义一个 const ， iota 初始值为0；
	每定义一个常量，iota+1;
	直到下一个const出现，iota清零*/

	const (
		a = iota //0
		b = iota //1
		c = iota //2
	)
	fmt.Println(a, b, c)
	const (
		d = iota
		e
	)
	fmt.Println(d, e)
	//枚举
	const (
		MALE = iota
		FEMAIL
		UNKNOW
	)
	fmt.Println(MALE, FEMAIL, UNKNOW)
}
