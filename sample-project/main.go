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
		ctx = context.WithValue(ctx, "baz", "qux")
		_ = ctx.Value("foo")
	}

	fmt.Println(time.Since(start))
}
