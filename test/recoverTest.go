package main

import "fmt"

func main() {

	//defer f()
	defer f2()()

	a()

}

func f() {
	fmt.Printf("f runing\n")
	if err := recover(); err != nil {
		fmt.Printf("ocurr error: %v !\n", err)
	}
	fmt.Printf("f end\n")
}

func f2() func() {
	fmt.Printf("f2 runing\n")
	fmt.Printf("f2 end\n")
	return func() {
		if err := recover(); err != nil {
			fmt.Printf("ocurr error: %v !\n", err)
		}
		fmt.Printf("f2 inner function runing\n")
	}
}

func a() {
	//panic("111")
	b()
}

func b() {
	fmt.Printf("b runing\n")
	panic("111")

}
