package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "nama", "apip")
	ctx = context.WithValue(ctx, "umur", 25)
	ctx = context.WithValue(ctx, "kota", "Jakarta")

	tampilkanData(ctx)
}

func tampilkanData(ctx context.Context) {
	nama := ctx.Value("nama")
	umur := ctx.Value("umur")
	kota := ctx.Value("kota")

	fmt.Println("Nama: ", nama)
	fmt.Println("Umur: ", umur)
	fmt.Println("Kota: ", kota)
}
