package main

import "fmt"

func modify(array [10]int) {
	array[0] = 10 // 试图修改数组的第一个元素
	fmt.Println("In modify(), array values:", array)
}

func main() {
	array := [5]int{1,2,3,4,5} // 定义并初始化一个数组
	modify(array) // 传递给一个函数，并试图在函数体内修改这个数组内容
	fmt.Println("In main(), array values:", array)
}

//# command-line-arguments
//./array.go:12:8: cannot use array (type [5]int) as type [10]int in argument to modify

//修改modify函数的array长度为5
