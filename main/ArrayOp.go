package main

import "fmt"

func main()  {
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	fmt.Println(array[3:])
	fmt.Println(array[:])
	fmt.Println(array[:5])
	fmt.Println(len(array))
	fmt.Println(cap(array))

	m:=make(map[string]string);
	m["name"]="wqs"
	m["age"]="28"
	fmt.Println(m)
	fmt.Println(m["name"])

	for k,v:=range m {
		fmt.Println("key ",k+" value ",v)
	}

	for _,v:=range m {
		fmt.Println("value ",v)
	}

	}
