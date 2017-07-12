package main

import (
	"fmt"
	"strings"
	"regexp"
)

func main()  {
	s := "hello"
	s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)

	s = s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)

	s = "hellohello"
	fmt.Println("字符串长度",len(s))
	fmt.Println("子串出现次数",strings.Count(s,"he"))
	fmt.Println("是否存在某个子串",strings.Contains(s,"he"))
	fmt.Println("是否存在某个子串",strings.ContainsAny(s,"he"))

	fmt.Println("是否存在某个子串前缀",strings.HasPrefix(s,"he"))
	fmt.Println("是否存在某个子串后缀",strings.HasSuffix(s,"o"))


	fmt.Println("替换子串",strings.Replace(s,"o","oo",1))

	//rune        alias for int32
	str:="abcdef"
	dst:=[]rune(str)
	result:=make([]rune,0)
	for i:=len(str)-1;i>=0;i-- {
		result=append(result,dst[i])
	}
	fmt.Println(string(result))

	m,_:=regexp.MatchString("^h",s)
	fmt.Println(m)



}
