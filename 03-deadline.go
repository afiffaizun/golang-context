package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Println("Deadline: ", deadline.Format("15:04:05"))
	fmt.Println("Memulai proses...")

	prosesData(ctx)
}

func prosesData(ctx context.Context) {
	i := 1
	for {
		select {
			case <-ctx.Done():
				fmt.Printf("Proses dihentikan. Total: %d item\n", i-1)
				return
			default:
				fmt.Printf("Memproses item ke-%d\n", i)
				i++
				time.Sleep(1000 * time.Millisecond)
		}
	}
}