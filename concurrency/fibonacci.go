package concurrency

// Fibonacci is return Fibonacci numbers
func Fibonacci(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}
