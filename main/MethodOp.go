package main

import "fmt"

func myfunc(arg ...int) {
	for _,n:=range arg{
		fmt.Println(n)
	}
}


//自定义方法
func selfMehtod()  {
	fmt.Println("执行 selfMehtod 方法")
}

//多个返回值
func swap(a,b string) (string,string) {
	return b,a;
}

//一个返回值
func sum(a ,b int ) int {
	return a+b;
}

//闭包函数   返回一个函数类型   func() int
//了解闭包的概念  https://zhuanlan.zhihu.com/p/22486908?refer=study-fe
/*
三行代码中，有一个局部变量 local，有一个函数 foo，foo 里面可以访问到 local 变量。

好了这就是一个闭包：

「函数」和「函数内部能访问到的变量」（也叫环境）的总和，就是一个闭包。

就这么简单。

有的同学就疑惑了，闭包这么简单么？

「我听说闭包是需要函数套函数，然后 return 一个函数的呀！」

比如这样：

function foo(){
  var local = 1
  function bar(){
    local++
    return local
  }
  return bar
}

var func = foo()
func()
这里面确实有闭包，local 变量和 bar 函数就组成了一个闭包（Closure）。

为什么要函数套函数呢？

是因为需要局部变量，所以才把 local 放在一个函数里，如果不把 local 放在一个函数里，local 就是一个全局变量了，达不到使用闭包的目的——隐藏变量（等会会讲）。

这也是为什么我上面要说「运行在一个立即执行函数中」。

有些人看到「闭包」这个名字，就一定觉得要用什么包起来才行。其实这是翻译问题，闭包的原文是 Closure，跟「包」没有任何关系。

所以函数套函数只是为了造出一个局部变量，跟闭包无关。

为什么要 return bar 呢？

因为如果不 return，你就无法使用这个闭包。把 return bar 改成 window.bar = bar 也是一样的，只要让外面可以访问到这个 bar 函数就行了。

所以 return bar 只是为了 bar 能被使用，也跟闭包无关。

闭包的作用

闭包常常用来「间接访问一个变量」。换句话说，「隐藏一个变量」。

假设我们在做一个游戏，在写其中关于「还剩几条命」的代码。
如果不用闭包，你可以直接用一个全局变量：

window.lives = 30 // 还有三十条命
这样看起来很不妥。万一不小心把这个值改成 -1 了怎么办。所以我们不能让别人「直接访问」这个变量。怎么办呢？

用局部变量。

但是用局部变量别人又访问不到，怎么办呢？

暴露一个访问器（函数），让别人可以「间接访问」。

代码如下：

!function(){

  var lives = 50

  window.奖励一条命 = function(){
    lives += 1
  }

  window.死一条命 = function(){
    lives -= 1
  }

}()
简明起见，我用了中文 :)

那么在其他的 JS 文件，就可以使用 window.奖励一条命() 来涨命，使用 window.死一条命() 来让角色掉一条命。

看到闭包在哪了吗？
 */
func getSequence() func() int  {
	i:=0;
	return func() int{
		i++;
		return i;
	}
}

func main()  {
	myfunc(1,2,3,4)
	selfMehtod();
	var a, b = 5, 7
	fmt.Println(sum(a,b));
	str_a,str_b:="str_a","str_b";
	str_a1,str_b1 := swap(str_a,str_b);
	fmt.Println("变量 str_a1, str_b1:",str_a1,str_b1);



	//变量是一个函数
	method:= func(info string )string {return "变量是一个函数："+info}

	//使用函数
	fmt.Println(method("调用"))


	// 闭包的使用
	nextSequence1 :=getSequence();
	fmt.Println(nextSequence1())
	fmt.Println(nextSequence1())
	fmt.Println(nextSequence1())

	nextSequence2 :=getSequence();
	fmt.Println(nextSequence2())
	fmt.Println(nextSequence2())
	fmt.Println(nextSequence2())



}
