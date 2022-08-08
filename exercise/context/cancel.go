package main

import (
	"context"
	"fmt"
)

func child(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		return
	}
	fmt.Println("キャンセルされていない")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx) // キャンセルするのでprintされない
	cancel()
	child(ctx) // printされる
}
