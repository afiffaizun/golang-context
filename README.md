# Belajar Golang Context

Contoh-contoh sederhana untuk memahami konsep Context di Go.

## Apa itu Context?

Context adalah package standar Go yang digunakan untuk:
- Membatalkan operasi yang sedang berjalan (cancellation)
- Mengatur batas waktu maksimal (timeout)
- Mengatur waktu deadline tertentu
- Menyimpan data yang terkait dengan request/operasi (values)

## Contoh-contoh

### 01-basic-cancel.go
Dasar cancellation - cara membatalkan operasi dengan context.

**Jalankan:**
go run 01-basic-cancel.go### 02-timeout.go
Context dengan timeout - operasi otomatis dibatalkan setelah waktu tertentu.

**Jalankan:**
go run 02-timeout.go### 03-deadline.go
Context dengan deadline - operasi dibatalkan pada waktu tertentu.

**Jalankan:**sh
go run 03-deadline.go### 04-with-value.go
Menyimpan dan mengambil nilai dari context.

**Jalankan:**
go run 04-with-value.go### 05-http-example.go
Contoh praktis menggunakan context dengan HTTP request.

**Jalankan:**sh
go run 05-http-example.go### 06-multiple-goroutines.go
Mengontrol multiple goroutine dengan satu context.

**Jalankan:**
go run 06-multiple-goroutines.go## Cara Menjalankan

Jalankan setiap contoh dengan perintah:
go run <nama-file>.go## Best Practices

1. **Selalu terima context sebagai parameter pertama**
   func myFunction(ctx context.Context, otherParams ...) {
       // ...
   }
   2. **Jangan simpan context di struct** - selalu pass sebagai parameter

3. **Gunakan defer cancel()** setelah membuat context dengan cancel
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   4. **Periksa ctx.Done()** dalam loop atau operasi panjang
   select {
   case <-ctx.Done():
       return
   default:
       // kerja
   }
   5. **Gunakan context.WithValue dengan hati-hati** - hanya untuk request-scoped data

## Context Methods

- `ctx.Done()` - Channel yang akan ter-close ketika context dibatalkan
- `ctx.Err()` - Error yang menjelaskan kenapa context dibatalkan
- `ctx.Value(key)` - Mengambil nilai dari context
- `ctx.Deadline()` - Mendapatkan deadline (jika ada)

## Referensi

- [Go Context Package](https://pkg.go.dev/context)
- [Go Blog: Context](https://go.dev/blog/context)
