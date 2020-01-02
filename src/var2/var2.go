package main

import "fmt"

var a = 10000 //全局变量不支持：=的方式
func main() {
	//num的内存地址
	var num int = 100
	fmt.Printf("数值：%d，地址：%p\n", num, &num)
	num = 200
	fmt.Printf("数值：%d，地址：%p\n", num, &num)
	//:=这种方式只要前面有一个声明是新的就可以，如果有旧的就会update
	num, name := 300, "go"
	fmt.Println(num, name)

	//var sum int = 100  //声明未使用的变量在go中编译不通过

	fmt.Println("------------默认值--------------")
	var x int     //0
	var y float64 //0.0
	var z string  //""
	var b bool    //false
	var c []int   //[]
	fmt.Println(x, y, z, b, c)
}
