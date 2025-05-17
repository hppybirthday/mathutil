package main

import (
	"fmt"
	"github.com/hppybirthday/jinshan/mathutil"
)

func main() {
	a := 2
	b :=3

	addResult := mathutil.Add(a,b)
	fmt.Printf("和：%d\n",addResult)

	multiplyResult := mathutil.Multiply(a,b)
	fmt.Printf("积：%d",multiplyResult)

}