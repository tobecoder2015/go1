package main

import (
	f "fmt"
)

const (
	WHITE=iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width,height,depth float64
	color Color
}

type  BoxList []Box

//正确写法
//func (b *Box) SetColor(c Color){
//    b.color=c
//}

//它的receiver是一个指向Box的指针，是的，你可以使用*Box。
//想想为啥要使用指针而不是Box本身呢？
//我们定义SetColor的真正目的是想改变这个Box的颜色，如果不传Box的指针，
//那么SetColor接受的其实是Box的一个copy，也就是说method内对于颜色值的修改，其实只作用于Box的copy，而不是真正的Box。所以我们需要传入指针。


//错误写法1
//func (b Box) SetColor(c Color){
//	b.color=c
//}



func (b *Box) Volume()float64  {
	return b.width*b.height*b.depth
}

func (bl BoxList)BiggestColor() Color  {
	v:=0.0
	//成员变量实例化，多个按顺序或者map
	k:=Color(WHITE)
	for _,b:=range bl{
		if bv:=b.Volume();bv>v{
			v=bv;
			k=b.color
		}
	}
	return k
}

func (bl BoxList) PainItBlack(){
	//错误写法
	//for _,b:=range bl{
	//	b.SetColor(BLACK)
	//}

	//在Go的for…range循环中，Go始终使用值拷贝的方式代替被遍历的元素本身，
	//简单来说，就是for…range中那个value，是一个值拷贝，而不是元素本身。
	//这样一来，当我们期望用&获取元素的地址时，
	//实际上只是取到了value这个临时变量的地址，
	//而非list中真正被遍历到的某个元素的地址。
	//而在整个for…range循环中，value这个临时变量会被重复使用

	//正确写法
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string{
	strings:=[]string{"WHITE","BLACK","BLUE","RED","YELLOW"}
        return strings[c]
}

func main()  {
   boxes:=BoxList{Box{4,4,4,RED},
	   Box{10,10,1,YELLOW},
	   Box{1, 1, 20, BLACK},
	   Box{10, 10, 1, BLUE},
	   Box{10, 30, 1, WHITE},
	   Box{20, 20, 20, YELLOW},
   }

	f.Printf("we have %d boxes in our set\n",len(boxes))
	f.Println("The volume of the first one is\n",boxes[0].Volume())
	f.Println("The color of the last one is",boxes[len(boxes)-1].color.String())


	f.Println("let me paint black ")
	boxes.PainItBlack()
	for _,b:=range boxes {
		f.Println(b.color.String())
	}
}
