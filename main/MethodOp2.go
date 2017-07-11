package main

import (
	"fmt"
)

type testInt func (int) bool

func isOdd(a int)bool {
	if a % 2 == 0 {
		return true
	}
	return false
}

func filter(slice []int ,f testInt)[]int{
	var result  []int;
	for _,value:=range slice{
		if f(value) {
			result=append(result,value)
		}
	}
	return result
}

/**
函数当做值和类型在我们写一些通用接口的时候非常有用，通过上面例子我们看到testInt这个类型是一个函数类型，
然后两个filter函数的参数和返回值与testInt类型是一样的，但是我们可以实现很多种的逻辑，这样使得我们的程序变得非常的灵活。
 */
func main()  {
	slice:=[]int{1,2,3,4,5,6,7}
	fmt.Print(slice)
	odd:=filter(slice,isOdd)
	fmt.Print(odd)



}
