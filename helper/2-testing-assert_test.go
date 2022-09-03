package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloRasyidev(t *testing.T) {
	res := HelloRasyidev("Taeri")

	assert.Equal(t, "Halo Rasyidev", res, "Harus `Halo Rasyidev` looh!")
}

/*
	Hasilnya pake Assert: Mantap guys, lengkap

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
	=== RUN   TestHelloRasyidev
			2-testing-assert_test.go:12:
									Error Trace:    ..\longpathlahpokonya\2-testing-assert_test.go:12
									Error:          Not equal:
																	expected: "Halo Rasyidev"
																	actual  : "Halo Taeri"

																	Diff:
																	--- Expected
																	+++ Actual
																	@@ -1 +1 @@
																	-Halo Rasyidev
																	+Halo Taeri
									Test:           TestHelloRasyidev
									Messages:       Harus `Halo Rasyidev` looh!
	--- FAIL: TestHelloRasyidev (0.00s)
	=== RUN   TestHelloWorld
	--- PASS: TestHelloWorld (0.00s)
	FAIL
	exit status 1
	FAIL    learn-go-unit-test/helper       0.035s

*/
