//每个功能必须有一个main包
package main


//导入包
import ("fmt")

/*
 这里面是注释代码
 */

//这个方法会初始化时执行
func init()  {
	fmt.Println("执行 init 方法")
}






//主函数，类似java
func main() {
	fmt.Println("Hello world!你好，世界！")

	var name string="wqs";
	fmt.Println("定义变量 name："+name);


	var age int=28;
	fmt.Println("定义变量 age："+string(age));

	var a, b, c = 5, 7, "abc"

	fmt.Println(a,b,c);

	fmt.Println(a>b)

	fmt.Println(a>>2)
	fmt.Println(b<<2)

	var point *int =&a;
	fmt.Println("定义指针变量 point：",point);
	fmt.Println("定义指针变量对应的值 *point：",*point);
	//循环语句
	for a:=0;a<10;a++{
		fmt.Println(a)
	}


}
