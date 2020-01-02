package main

import "fmt"

func main() {
	/* 变量:variable;声明;访问 */
	var name int
	name = 10
	fmt.Println(name)

	var name2 int = 15
	fmt.Println(name2)

	/*go的类型推断*/
	var name3 = "奥利给"
	fmt.Printf("类型是：%T，数值是：%s\n", name3, name3)

	/*省略声明*/
	name4 := 20
	fmt.Println(name4)

	/*多变量声明*/
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	var l, m, n int = 11, 22, 33
	fmt.Println(l, m, n)

	var x, y, z = 100, 2.567, "string"
	fmt.Println(x, y, z)

	var (
		student = "小明"
		age     = 18
	)
	fmt.Println(student, age)
}
