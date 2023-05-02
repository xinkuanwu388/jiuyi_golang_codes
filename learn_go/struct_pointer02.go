package main

import "fmt"

// Rect
// struct 类似于 java 中的类，可以在 struct 中定义成员变量。
// 要访问成员变量，可以有两种方式：
// 1.通过 struct 变量.成员 变量来访问。
// 2.通过 struct 指针.成员 变量来访问。
// 不需要通过 getter, setter 来设置访问权限。/*
type rect struct { //定义矩形类
	x, y          float64 //类型只包含属性，并没有方法
	width, height float64
}

func (r *rect) area() float64 { //为Rect类型绑定Area的方法，*Rect为指针引用可以修改传入参数的值
	return r.width * r.height //方法归属于类型，不归属于具体的对象，声明该类型的对象即可调用该类型的方法
}

func main() {
	var rect1 rect
	rect1.width = 10
	rect1.height = 20

	fmt.Println(rect1, rect1.area())
}
