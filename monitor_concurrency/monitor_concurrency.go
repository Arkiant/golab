// Concurrency Monitors
// https://medium.com/dm03514-tech-blog/golang-monitors-and-mutexes-a-light-survey-84f04f9b7c09
package main

import (
	"context"
	"fmt"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	in := NewCounterMonitor(ctx)
	in <- 10

}

func NewCounterMonitor(ctx context.Context) chan<- int {
	ch := make(chan int)

	go func() {
		counter := 0
		for {
			select {
			case i, ok := <-ch:
				if !ok {
					return
				}
				counter += i
			case <-ctx.Done():
				fmt.Printf("Processed: %d", counter)
				return
			}
		}

	}()

	return ch
}
