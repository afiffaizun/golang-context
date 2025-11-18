package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 6 * time.Second)
	defer cancel()

	fmt.Println("Memulai operasi yang lama....")

	hasil := operasiLama(ctx)

	if hasil {
		fmt.Println("Operasi berhasil: ")
	} else {
		fmt.Println("Operasi gagal")
	}
}

func operasiLama(ctx context.Context) bool {
	for i := 1; i <= 10; i++ {
		select {
			case <-ctx.Done():
				fmt.Println("Operasi dibatalkan: ", i-1)
				return false
			default:
				fmt.Println("Iterasi: ", i)
				time.Sleep(500 * time.Millisecond)
		}
	}
	return true
}