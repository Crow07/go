package main

import "fmt"


//三种变量的用法
func main(){
	var v1 int = 1
	var v2 = 1
	v3 := 1   //出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误
	fmt.Println(v1, v2, v3)
}
