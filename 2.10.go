package main

import (
	"fmt"
	"log"
)

func main() {

	//showBytes()
	//swapTwoInts()
	ShowSwapStep()

}

func showBytes() {
	log.Println(fmt.Sprintf("%02x", 3490593))
	log.Println(fmt.Sprintf("%x", 3490593.0))
	log.Println(fmt.Sprintf("%x", "12345"))
	log.Println(fmt.Sprintf("%x", "A"))
	log.Println(fmt.Sprintf("%x", "Z"))
}

func swapTwoInts() {
	// 通过二进制环属性交换变量
	var a = 3
	var b = 7

	log.Println("a:", a, "b:", b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	log.Println("a:", a, "b:", b)
}

// 二进制亦或操作可实现环
func ShowSwapStep() {
	var a = 18
	var b = 25

	log.Println("a:", fmt.Sprintf("%b", a))
	log.Println("b:", fmt.Sprintf("%b", b))

	log.Println("^: ", fmt.Sprintf("%b", a^b))
	log.Println("^:", fmt.Sprintf("%b", a^b^a))
	log.Println("^:", fmt.Sprintf("%b", a^b^b)) // 亦或同一个值两次则变回原值

}
