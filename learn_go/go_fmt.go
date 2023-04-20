package main

import (
	"fmt"
	"os"
)

//https://pkg.go.dev/fmt#hdr-Printing

func main() {

	// Print 函数直接输入内容到终端
	fmt.Print("hello")
	fmt.Print("我叫Print\n")

	//Println 输出都是独占一行的,自动帮我们添加了 \n
	fmt.Println("我是Println")
	fmt.Println("人称fmt")

	//Printf 函数支持格式化输出字符串
	fmt.Printf("我是Printf，人称: %s", "fmt")
	fmt.Printf("我是Printf，人称: %s \n", "fmt")
	// Using %d to print an int value
	fmt.Printf("%d", 36) // 36
	// Using %f to print a float value with default width and precision
	fmt.Printf("%f", 20.66) // 20.660000
	// Using %f with the precision length 1
	fmt.Printf("%.1f", 20.33) // 20.3
	// The %c formatter is used to format a single character
	letter := 't'
	fmt.Printf("%c", letter) // t
	secret := '🤫'
	fmt.Printf("%c", secret) // 🤫
	// The %s formatter is great to format a string
	fmt.Printf("%s", "This is a random string") // This is a random string
	// The %t formatter is suitable for Boolean values
	fmt.Printf("%t", false) // false
	// Formatting the width of a string
	fmt.Printf("|%8s|", "pikachu") // | pikachu| (this action added one additional space where the string begins)
	fmt.Printf("%%")               // %
	// %[]s 可指定顺序，重复使用
	a, b := "First variable", "Second variable"
	fmt.Printf("%[1]s | %[1]s \n", a) // First variable | First variable
	fmt.Printf("%[2]s | %[1]s", a, b) // Second variable | First variable

	// Sprintf() create and return strings
	first := fmt.Sprintf("%d", 500)          // 'first' variable now has the value of 500
	binaryVariable := fmt.Sprintf("%b", 500) // %b 'binaryVariable' variable now has the value of 111110100
	fmt.Println(first, binaryVariable)
	// 向标准输出写入内容
	// os.Stdout 只能输出到 控制台
	_, err := fmt.Fprintln(os.Stdout, "向标准输出写入内容") // io结束后输出内容
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}

	// os.OpenFile(文件路径，文件打开模式，文件权限)
	fileObj, err2 := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// 异常捕获
	if err2 != nil {
		fmt.Println("打开文件出错，err2:", err2)
		return
	}
	text := "写入文件的信息"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s\n", text)

	fmt.Printf("---------%%t----------\n")
	//fmt.Printf("%t\n", 0)//不可以
	fmt.Printf("%t\n", false)
	// %s
	fmt.Printf("----------%%s---------\n")
	fmt.Printf("%8s\n", "12345")
	fmt.Printf("%8.2s\n", "12345")
	fmt.Printf("%-8.4s\n", "12345")
	fmt.Printf("%.4s\n", "12345")
	fmt.Printf("%.6s\n", "12345")

	//fmt.Scan 输入
	var foo int    // foo is 0
	var str string // string is ""

	fmt.Scan(&foo, &str) // If you enter a string character, the scan function
	// will leave the variable foo unchanged
	fmt.Println(foo, str)

	var name string
	var age int
	fmt.Print("Enter your name: ") // Writing a request message to the stdout
	fmt.Scan(&name)                // Reading from the stdin into the name variable
	fmt.Println("")                // Going to the next line by writing /n to the stdout

	fmt.Print("Enter your age: ") // Writing a request message to the stdout
	fmt.Scan(&age)                // Reading from the stdin into the age variable
	fmt.Println("")               // Going to the next line by writing /n to the stdout

	fmt.Print(name, age) // Writing to the stdout the values of name and
	// age variables that you have entered

}
