package main

import "fmt"
type Circle struct {
	radius float32;
}

func (c Circle)getArea() float32{
	return 3.14*c.radius*c.radius;
}


func (c Circle) getLength() float32{
	return 3.14*c.radius*2;
}


func (c Circle) getInfo() {
	fmt.Println("这是一个圆");
}


func main()  {
	//类型的使用
	var c1 Circle;
	c1.radius=10;
	fmt.Println(c1.getArea())
	fmt.Println(c1.getLength())
	c1.getInfo()
}
