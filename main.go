package main

import "fmt"

func main() {
	numbers := generator(1, 2, 3, 4, 5)
	squared := square(numbers)
	doubled := double(squared)

	print(doubled)
}

func generator(input ...int) <-chan int {

	out := make(chan int)
	go func() {
		for _, n := range input {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func print(in <-chan int) {
	for n := range in {
		fmt.Println("Result: ", n)
	}
}
