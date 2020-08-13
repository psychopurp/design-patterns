package main

import "fmt"

/*
生成器模式  使用channel 阻塞原理
类似python的 generator
*/

func fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		//i代表前一项，j代表当前项
		for i, j := 1, 0; i < n; i, j = j, i+j {
			// fmt.Println(i, j)
			out <- j
		}
	}()

	return out
}

func main() {
	for v := range fib(100) {
		fmt.Println(v)
	}

}
