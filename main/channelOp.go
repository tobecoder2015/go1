package main

import (
	f "fmt"
)

func sum2(a []int,c chan int)  {
	total:=0
	for _,v:=range a{
		total=total+v
	}
	c<-total
}

func fibonacci(n int,c chan int){
	x,y:=1,1
	for i:=0;i<n;i++ {
		c<-x
		x,y=y,x+y
	}
	close(c)
}
func main()  {
	c := make(chan int)
	arr:=[8]int{1,2,3,4,5,6,7,8}
	go sum2(arr[:len(arr)/2],c)
	go sum2(arr[len(arr)/2:],c)
	x,y:=<-c,<-c
	f.Println(x,y,x+y)

	fibonacciChan:=make(chan int,10)
	 go fibonacci(cap(fibonacciChan),fibonacciChan)
	for i:=range fibonacciChan{
		f.Println(i)
	}
}
