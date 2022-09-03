package helper

import (
	"fmt"
	"testing"
)

func TestPembagian(t *testing.T) {
	res, err := Pembagian(24.3, 0)

	// jika error, gagalkan pakai t.Fail()
	if err != nil {
		t.FailNow()
	}

	fmt.Println("Hasil pembagian adalah", res)
}

func TestPembagianLagi(t *testing.T) {
	res, err := Pembagian(53, 0)

	// jika error, gagalkan pakai t.FailNow()
	if err != nil {
		t.Fail()
	}

	fmt.Println("Hasil pembagian lagi adalah", res)
}

/*
	$ go test -v
	=== RUN   TestPembagian
	--- FAIL: TestPembagian (0.00s) // FailNow(), tidak dijalankan blok program selanjutnya, tapi unit test tetep dilanjutkan
	=== RUN   TestPembagianLagi
	Hasil pembagian lagi adalah 0
	--- FAIL: TestPembagianLagi (0.00s)
	=== RUN   TestHelloWorld
	--- PASS: TestHelloWorld (0.00s)
	FAIL
	exit status 1
	FAIL    learn-go-unit-test/helper       0.032s
*/

func TestPembagianFatal(t *testing.T) {
	res, err := Pembagian(53, 0)

	// jika error, gagalkan pakai t.FailNow()
	if err != nil {
		t.Fatal("`Fatal()` | Tidak boleh pembagian dengan nol")
	}

	fmt.Println("`Fatal()` Hasil pembagian lagi adalah", res)
}

func TestPembagianError(t *testing.T) {
	res, err := Pembagian(53, 0)

	// jika error, gagalkan pakai t.FailNow()
	if err != nil {
		t.Error("`Error()` | Tidak boleh pembagian dengan nol")
	}

	fmt.Println("`Error()` Hasil pembagian lagi adalah", res)
}

func TestPembagianAkhir(t *testing.T) {
	fmt.Println("Ini akhir test pembagian")
}

/*
	$ go test -v
	=== RUN   TestPembagian
	--- FAIL: TestPembagian (0.00s)
	=== RUN   TestPembagianLagi
	Hasil pembagian lagi adalah 0
	--- FAIL: TestPembagianLagi (0.00s)
	=== RUN   TestPembagianFatal
			1-testing-fail_test.go:49: `Fatal()` | Tidak boleh pembagian dengan nol
	--- FAIL: TestPembagianFatal (0.00s)
	=== RUN   TestPembagianError
			1-testing-fail_test.go:60: `Error()` | Tidak boleh pembagian dengan nol
	`Error()` Hasil pembagian lagi adalah 0
	--- FAIL: TestPembagianError (0.00s)
	=== RUN   TestPembagianAkhir
	Ini akhir test pembagian
	--- PASS: TestPembagianAkhir (0.00s)
	=== RUN   TestHelloWorld
	--- PASS: TestHelloWorld (0.00s)
	FAIL
	exit status 1
	FAIL    learn-go-unit-test/helper       0.039s
*/
