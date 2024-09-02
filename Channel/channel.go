package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)

	//продюсер, так как записывает
	go func() {
		defer close(out)
		for _, num := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- num:
			}

		}
	}()

	return out
}

func sq(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {

			select {
			case <-ctx.Done():
				return
			case out <- func() int {
				time.Sleep(100 * time.Millisecond)
				return num * num
			}():
			}
		}
	}()

	return out
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*400)
	for val := range sq(ctx, gen(ctx, 1, 2, 3, 4, 5)) {
		fmt.Println(val)
	}
}
