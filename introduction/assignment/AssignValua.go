package main

import "fmt"

func main() {
	fmt.Println("费布拉奇数列")
	fmt.Println(f(0))
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(10))
	fmt.Println("最大公约数")
	fmt.Println(gcd(36,48))
	fmt.Println(gcd(10,5))
	fmt.Println(gcd(10,7))


}

//费布拉奇数列
func f(n int) int {
	x,y := 0,1
	for i:=0; i<n;i++  {
		x,y = y,x+y
	}
	return x
}

//最大公约数
func gcd(x,y int) int  {
	for y!=0 {
		x,y = y,x%y
	}
	return x
}