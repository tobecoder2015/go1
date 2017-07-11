package main

import (
	f "fmt"
)

//interface{} 可以存储任意类型
func  allPara(any interface{}) {
	f.Println(any)
}


func  allPara2(any ...interface{}) {
	f.Println(any)
}

func main()  {
	allPara(1)
	allPara("aa")
	allPara(3.23)
	arr :=[4]int{1,2,3,4}
	allPara(arr)
	allPara2(arr)


}
