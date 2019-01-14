package main

import "fmt"

func main() {
	var a int = 10
	var ap *int
	fmt.Println("a\t=\t", a)
	fmt.Println("&a\t=\t", &a)
	fmt.Println("ap\t=\t", ap)
	fmt.Println("&ap\t=\t", &ap)
	fmt.Println()

	ap = &a

	fmt.Println("ap\t=\t", ap)
	fmt.Println("&ap\t=\t", &ap)
	fmt.Println("*ap\t=\t", *ap)

}
