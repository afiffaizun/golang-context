package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go kerja(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("Membatalkan context....")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Selesai")
}

func kerja(ctx context.Context) {
	for i := 1; ; i++ {
		select {
			case <-ctx.Done():

				fmt.Println("Kerja dihentikan: ", ctx.Err())
				return
			
			default:
				fmt.Printf("Kerja dihentikan ke-%d\n", i)
				time.Sleep(1 * time.Second)
		}
	}
}