package concurrency

import "golang.org/x/net/context"

// PlusOne returns a channel of num + 1 for nums received form in.
func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

// IntPipe is return channel int
type IntPipe func(<-chan int) <-chan int
type IntPipeWithCtx func(context.Context, <-chan int) <-chan int

// Chain is return IntPipe
func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

// Distribute ...
func Distribute(p IntPipe, n int) IntPipe {
	return func(in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(in)
		}
		return FanIn(cs...)
	}
}
