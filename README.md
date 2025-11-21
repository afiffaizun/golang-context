# Belajar Golang Context 📚

Repositori ini berisi contoh-contoh sederhana untuk memahami konsep **Context** di Go (Golang).

## 📋 Daftar Isi

- [Apa itu Context?](#apa-itu-context)
- [Konsep Dasar](#konsep-dasar)
- [Contoh-contoh](#contoh-contoh)
- [Cara Menjalankan](#cara-menjalankan)
- [Best Practices](#best-practices)

## 🤔 Apa itu Context?

Context adalah package standar Go yang digunakan untuk:

1. **Cancellation** - Membatalkan operasi yang sedang berjalan
2. **Timeout** - Mengatur batas waktu maksimal untuk operasi
3. **Deadline** - Mengatur waktu deadline tertentu
4. **Values** - Menyimpan data yang terkait dengan request/operasi

## 🎯 Konsep Dasar

### Context Background & TODO

// context.Background() - root context, biasanya untuk main()
ctx := context.Background()

// context.TODO() - placeholder saat belum jelas context mana yang digunakan
ctx := context.TODO()
### Context dengan Cancellation

Membatalkan operasi secara manual dengan memanggil `cancel()`.

### Context dengan Timeout

Membatalkan operasi secara otomatis setelah waktu tertentu.

### Context dengan Deadline

Membatalkan operasi pada waktu tertentu (deadline).

### Context dengan Values

Menyimpan dan mengambil nilai dari context.

## 📁 Contoh-contoh

### 1. `01-basic-cancel.go` - Dasar Cancellation

Contoh paling sederhana: cara membatalkan operasi dengan context.

**Konsep:**
- Membuat context yang bisa dibatalkan
- Menggunakan `ctx.Done()` untuk mendeteksi pembatalan
- Membatalkan dengan `cancel()`

**Jalankan:**
go run 01-basic-cancel.go### 2. `02-timeout.go` - Context dengan Timeout

Contoh penggunaan timeout: operasi otomatis dibatalkan setelah waktu tertentu.

**Konsep:**
- `context.WithTimeout()` untuk membuat context dengan timeout
- Operasi akan otomatis dibatalkan setelah timeout tercapai

**Jalankan:**
go run 02-timeout.go### 3. `03-deadline.go` - Context dengan Deadline

Contoh penggunaan deadline: operasi dibatalkan pada waktu tertentu.

**Konsep:**
- `context.WithDeadline()` untuk membuat context dengan deadline
- Deadline adalah waktu absolut, bukan durasi

**Jalankan:**
go run 03-deadline.go### 4. `04-with-value.go` - Context dengan Nilai

Contoh menyimpan dan mengambil nilai dari context.

**Konsep:**
- `context.WithValue()` untuk menyimpan nilai
- `ctx.Value(key)` untuk mengambil nilai
- Gunakan dengan hati-hati, hanya untuk request-scoped data

**Jalankan:**h
go run 04-with-value.go### 5. `05-http-example.go` - Contoh dengan HTTP Request

Contoh praktis menggunakan context dengan HTTP request.

**Konsep:**
- Menggunakan context untuk timeout HTTP request
- `http.NewRequestWithContext()` untuk membuat request dengan context

**Jalankan:**
go run 05-http-example.go### 6. `06-multiple-goroutines.go` - Multiple Goroutines

Contoh mengontrol multiple goroutine dengan satu context.

**Konsep:**
- Satu context bisa digunakan untuk membatalkan banyak goroutine
- Menggunakan `sync.WaitGroup` untuk menunggu semua goroutine selesai

**Jalankan:**
go run 06-multiple-goroutines.go## 🚀 Cara Menjalankan

### Prerequisites

- Go 1.16 atau lebih baru
- Terminal/Command Prompt

### Menjalankan Contoh

Jalankan setiap contoh dengan perintah:

# Contoh 1: Basic Cancel
go run 01-basic-cancel.go

# Contoh 2: Timeout
go run 02-timeout.go

# Contoh 3: Deadline
go run 03-deadline.go

# Contoh 4: With Value
go run 04-with-value.go

# Contoh 5: HTTP Example
go run 05-http-example.go

# Contoh 6: Multiple Goroutines
go run 06-multiple-goroutines.go### Menjalankan Semua Contoh

# Linux/Mac
for file in *.go; do echo "=== Running $file ==="; go run "$file"; echo; done

# Windows (PowerShell)
Get-ChildItem *.go | ForEach-Object { Write-Host "=== Running $($_.Name) ==="; go run $_.Name; Write-Host "" }## ✅ Best Practices

1. **Selalu terima context sebagai parameter pertama**
   func myFunction(ctx context.Context, otherParams ...) {
       // ...
   }
   2. **Jangan simpan context di struct**
   - Selalu pass context sebagai parameter
   - Jangan simpan sebagai field di struct

3. **Gunakan defer cancel()**
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel() // Pastikan selalu dipanggil
   4. **Periksa ctx.Done() dalam loop**
   
   for {
       select {
       case <-ctx.Done():
           return
       default:
           // kerja
       }
   }
   5. **Gunakan context.WithValue dengan hati-hati**
   - Hanya untuk request-scoped data
   - Jangan untuk parameter fungsi yang optional
   - Gunakan type-safe keys

6. **Jangan pass nil context**
   - Gunakan `context.Background()` jika tidak ada context yang tersedia

## 📖 Penjelasan Detail

### Context Methods

| Method | Deskripsi |
|--------|-----------|
| `ctx.Done()` | Channel yang akan ter-close ketika context dibatalkan |
| `ctx.Err()` | Error yang menjelaskan kenapa context dibatalkan |
| `ctx.Value(key)` | Mengambil nilai dari context |
| `ctx.Deadline()` | Mendapatkan deadline (jika ada) |

### Context Errors

- `context.Canceled`: Context dibatalkan secara manual
- `context.DeadlineExceeded`: Context dibatalkan karena timeout/deadline

### Kapan Menggunakan Context?

- ✅ HTTP requests dengan timeout
- ✅ Database queries dengan timeout
- ✅ Long-running operations yang perlu bisa dibatalkan
- ✅ Passing request-scoped values (user ID, request ID, dll)
- ✅ Cancelling multiple goroutines sekaligus

### Kapan TIDAK Menggunakan Context?

- ❌ Jangan untuk parameter fungsi yang optional
- ❌ Jangan untuk configuration
- ❌ Jangan untuk dependency injection

## 🔗 Referensi

- [Go Context Package Documentation](https://pkg.go.dev/context)
- [Go Blog: Context](https://go.dev/blog/context)
- [Effective Go - Context](https://go.dev/doc/effective_go#context)

## 📝 Catatan

- Contoh-contoh ini dibuat dengan sederhana untuk memudahkan pemahaman
- Dalam production, pastikan untuk handle error dengan baik
- Context adalah immutable - setiap `With*` method membuat context baru

## 🤝 Kontribusi

Jika ada pertanyaan atau saran, silakan buat issue atau pull request!

## 📄 License

MIT License - bebas digunakan untuk belajar dan referensi.

---

**Happy Learning! 🎉**
