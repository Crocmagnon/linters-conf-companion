package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "foo", "bar")

	for range 10_000 {
		ctx = context.WithValue(ctx, "baz", "qux") // want "nested context in loop"
		_ = ctx.Value("foo")
	}

	for i := 0; i < 10_000; i++ {
		ctx = context.WithValue(ctx, "baz", "qux") // want "nested context in loop"
		_ = ctx.Value("foo")
	}

	fmt.Println(time.Since(start))
}
