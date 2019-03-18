// hi_fan.go

package main

func square(inCh <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		defer close(out)

		for n := range inCh {

			out <- n * n

		}

	}()

	return out

}

func main() {

	in := producer(10000000)

	// FAN-OUT

	c1 := square(in)

	c2 := square(in)

	c3 := square(in)

	// consumer

	for _ = range merge(c1, c2, c3) {

	}

}
