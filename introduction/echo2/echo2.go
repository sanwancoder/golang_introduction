//必须定义为main 否则会报错
package main

import (
	"fmt"
	"os"
)

func main() {
	var s,sep string
	for i:=1; i<len(os.Args);i++  {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}