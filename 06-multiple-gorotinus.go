package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Jalankan 3 goroutine sekaligus
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// Tunggu 4 detik, lalu batalkan semua
	time.Sleep(4 * time.Second)
	fmt.Println("\nMembatalkan semua worker...")
	cancel()

	wg.Wait()
	fmt.Println("Semua worker selesai")
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d berhenti (total kerja: %d)\n", id, i-1)
			return
		default:
			fmt.Printf("Worker %d: bekerja ke-%d\n", id, i)
			time.Sleep(1 * time.Second)
		}
	}
}