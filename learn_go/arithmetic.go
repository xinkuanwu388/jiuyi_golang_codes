package main

func main() {
	//优先级 ： 括号 > 一元正负 > 乘法、除法和模 > 加减法

	/*  +
	result := 1 + 5 // result = 6
	result2 := 1234 + 4321 // result = 5555

	Because integer variables can be both positive and negative numbers (and 0 too), we can add a negative number to a positive number:
	result := -1000 + 2000 // result = 1000

	The addition will behave similarly with a float64 variable:
	result := 10.1 + 1.234 // result = 11.334
	result2 := -5.6789 + 19.23 // result = 13.551100000000002

	Go will only allow us to use arithmetic operators on the same data types. For example, we can't add an int and a float64.
	var x = 8
	var y = 13.0
	result := x + y // invalid operation: x + y (mismatched types int and float64)
					// IDE will warn you, and if you try to compile,
					// the Go compiler will print a compile-time error
	*/
	/*   * / %

	x := 42.3
	y := 13.1
	result := x * y // result = 554.13

	Division results in Go will vary according to the numeric type of the values.
	When dividing int values, the result will always be rounded down to the nearest integer;
	to get the result as a decimal number,
	we can wrap our values within float32() or float64(). For example:
	// Dividing 'int' values
	x := 13
	y := 2
	result := x / y // result = 6

	a := -100
	b := 3
	result2 := a / b // result2 = -33

	// Casting 'int' values as 'float64'
	x = 20
	y = 3
	result3 := float64(x) / float64(y) // result3 = 6.666...7

	// % 取余 can only be used with the int type values.
	x := 100
	y := 15
	result := x % y // result = 10

	*/

	/* =

	x := 5

	// Compound addition
	x += 1 // equivalent to x = x + 1 -> x = 6

	// Compound subtraction
	x -= 10 // equivalent to x = x - 10 -> x = -4

	// Compound multiplication
	x *= 4 // equivalent to x = x * 4 -> x = -16

	// Compound division
	x /= 2 // equivalent to x = x / 2 -> x = -8

	// Compound modulo
	x %= 3 // equivalent to x = x % 3 -> x = -2 (final result.)

	注意，当在 = 符号的右边使用带有算术运算的复合赋值运算符时，算术运算优先于复合赋值运算符，例如:
	y := 6
	y *= 3 + 4 // equivalent to y = y * (3 + 4) -> y = 42
	*/

	/* ++ --
	num := 0
	num++ // num = 1
	num++ // num = 2
	num-- // num = 1
	num-- // num = 0
	它们是语句，而不是表达式; 在计算机科学中，语句永远不会返回值。
	因为它不能作为值计算，所以我们
	不能将它作为参数传递给另一个函数，不能将它直接赋给另一个变量，也不能将这些操作符合并到另一个表达式或语句中
	num1 := 1
	num1++ // Works!
	num2 := num1++ // Doesn't work!
	num3 := (num1++)++ // Doesn't work!
	num4 := 1 + num1++ // Doesn't work!

	*/
}
