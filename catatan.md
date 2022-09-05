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

### `Fail()` dan `FailNow()`
- Behavior: tetap melanjutkan unit test
  - `Fail()` | Blok program selanjutnya tetap dijalankan. Result: FAIL
  - `FailNow()` | Blok program selanjutnya tidak dijalankan. Result: FAIL

### `Error()` dan `Fatal()`
- Behavior: Logging error dan memanggil function `Fail()` atau `FailNow()`
- `Error()` : Logging error lalu memanggil `Fail()`, blok program selanjutnya tetap dijalankan.
- `Fatal()` : Logging error lalu memanggil `FailNow()`, blok program selanjutnya tetap dijalankan.

## Assertion
- Logging error menggunakan library Testify github.com/stretchr/testify/assert
- Lengkap, mudah, mantap lah

## Testify: assert vs require
- assert: jika test gagal, akan memanggil `Fail()`
- require: jika test gagal, akan memanggil `FailNow()`

## Skip Test atau Membatalkan Test (bukan membuat test jadi FAIL)
- Dibutuhkan saat adanya test untuk os tertentu saja
- Menggunakan function `Skip()`

## Before dan After Testing
- Bisa menggunakan testing.M | Methodnya `Run()`

## Table Testing
- Membuat list test-case dan jalankan di `t.Run()` dalam loop test-case tersebut

## Mock: Testify Mock
- Unit Testing yang baik tidak harus menjalankan third-party service apapun.
- Membuat testing untuk objek yang sulit untuk di-test.
  - Cth: Testing third-party API
- Jika desain kode program tidak baik, akan sulit untuk mengimplementasikan mock.

## Benchmark
- Mekanisme menghitung performa kode aplikasi
- Dilakukan berkali-kali secara otomatis
- Menghasilkan laporan total run code dalam nanosecond, microsecond, milisecond, second

## Benchmark Menggunakan `testing.B()`
- Mirip dengan testing.T
- Memiliki function `Fail()`, `FailNow()`, `Error()`, `Fatal()` dll
- Dijalankan bersamaan dengan Unit Test.
- Terdapat beberapa atribut tambahan untuk melakukan benchmarking
- Attribute `N` digunakan untuk menghitung total iterasi benchmark
- RULES
  - Nama function re"Benchmark.+"
  - Memiliki parameter `*testing.B`
  - Tidak ada return value
  - Nama file benchmark diakhiri dengan `_test`