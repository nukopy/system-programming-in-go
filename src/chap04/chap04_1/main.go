package main

import (
	"fmt"
	"time"
)

func main() {
	testGoroutine1()
	testGoroutine2()
	// testGoroutine3()
}

func sub() {
	fmt.Println("start sub()")
	fmt.Println("sub(): Wait for 1 second...")
	time.Sleep(1 * time.Second)
	fmt.Println("end sub()")
}

func testGoroutine1() {
	fmt.Println("start main()")
	sub()
	fmt.Println("main(): Wait for 2 seconds...")
	time.Sleep(2 *time.Second)
	fmt.Println("end main()")
	fmt.Println()
}

func testGoroutine2() {
	fmt.Println("start main()")
	go sub()
	fmt.Println("main(): Wait for 2 seconds...")
	time.Sleep(2 *time.Second)
	fmt.Println("end main()")
	fmt.Println()
}

func testGoroutine3() {
	fmt.Println("start main()")
	go func() {
		fmt.Println("start sub()")
		fmt.Println("sub(): Wait for 1 second...")
		time.Sleep(1 * time.Second)
		fmt.Println("end sub()")
	}()
	fmt.Println("main(): Wait for 2 seconds...")
	time.Sleep(2 *time.Second)
	fmt.Println("end main()")
	fmt.Println()
}