package main

import "fmt"

//参数不定数量，不定类型，...interface{}
func myfunc(args ...interface{} {
	for _, arg := range args{
		switch arg.(type) {
			case int:
				fmt.Println(arg,"is int")
			case string:
				fmt.Println(arg,"is a string value")
			case int64:
				fmt.Println(arg,"is a int64 value")
			default:
				fmt.Println(arg, "is an unknow type")
		}
	}
}

func main() {
	v1 := 1
	v2 := 23454
	v3 := "hah"
	v4 := 1.23
	myfunc(v1, v2 ,v3 ,v4)
}