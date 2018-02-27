package main

import "fmt"

//传递不定数量的参数
func myfunc(args ...int){
        for _, arg := range args {
                fmt.Println(arg)
        }
}

func main() {
        myfunc(1, 2, 3)
}
//main函数要加()