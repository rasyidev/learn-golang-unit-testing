# Golang Testing Package

## testing.T
- struct bernama testing.T
- Untuk unit test

## testing.M
- Struct untuk mengatur testing lifecycle

## testing.B
- Struct untuk benchmarking
- Untuk mengukur kecepatan program

## Aturan Function Unit Test
- Nama function: re"Test.+"
- Best Practice: Test<Nama Function>
  - Cth: TestHelloWorld()
- Harus memiliki parameter *testing.T, tidak ada return

## Menjalankan Test
- Pindah working directory ke file yang ada testnya
- Jalankan `go test`
- `go test -v` | Mengaktifkan verbose (Untuk melihat function mana saja yang ditest) saat menjalankan test
- `go test ./...` | Menjalankan test dari module working directory dan nested-nya.


## Menggagalkan Test
- Menggagalkan test menggunakan `panic` tidak disarankan, karena sekali gagal maka test yang lain tidak akan dijalankan
- `Fail()`
- `FailNow()`
- `Error()`
- `Fatal()`
- 