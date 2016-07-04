package concurrency

import "fmt"

func ExampleFanIn() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		defer close(c1)
		defer close(c2)
		c1 <- 1
		c1 <- 2
		c1 <- 3
		c1 <- 4
		c2 <- 5
		c2 <- 6
		c2 <- 7
		c2 <- 8
	}()
	for c := range FanIn(c1, c2) {
		fmt.Println(c)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
}
