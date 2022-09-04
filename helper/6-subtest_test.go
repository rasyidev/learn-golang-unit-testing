package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtest(t *testing.T) {
	// subtest 1
	t.Run("Taeri", func(t *testing.T) {
		res := HelloRasyidev("Kim Taeri")
		assert.Equal(t, "Halo Kim Teari", res, "Seharusnya `Halo Kim Taeri`")
	})

	// subtest 2
	t.Run("Rasyidev", func(t *testing.T) {
		res := HelloRasyidev("Rasyidev")
		assert.Equal(t, "Halo Rasyidev", res, "Seharusnya `Halo Rasyidev`")
	})
}

/*
$ go test -v -run TestSubtest
Before Testing
=== RUN   TestSubtest
=== RUN   TestSubtest/Taeri
    6-subtest_test.go:13:
                Error Trace:    learn-go-unit-test\helper\6-subtest_test.go:13
                Error:          Not equal:
                                expected: "Halo Kim Teari"
                                actual  : "Halo Kim Taeri"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -Halo Kim Teari
                                +Halo Kim Taeri
                Test:           TestSubtest/Taeri
                Messages:       Seharusnya `Halo Kim Taeri`
=== RUN   TestSubtest/Rasyidev
--- FAIL: TestSubtest (0.00s)
    --- FAIL: TestSubtest/Taeri (0.00s)
    --- PASS: TestSubtest/Rasyidev (0.00s)
FAIL
After Testing
exit status 1
FAIL    learn-go-unit-test/helper       0.668s
*/

/*
$ go test -v -run TestSubtest/Rasyidev
Before Testing
=== RUN   TestSubtest
=== RUN   TestSubtest/Rasyidev
--- PASS: TestSubtest (0.00s)
    --- PASS: TestSubtest/Rasyidev (0.00s)
PASS
After Testing
ok      learn-go-unit-test/helper       0.679s
*/

/*
$ go test -v -run TestSubtest/Taeri
Before Testing
=== RUN   TestSubtest
=== RUN   TestSubtest/Taeri
    6-subtest_test.go:13:
                Error Trace:    learn-go-unit-test\helper\6-subtest_test.go:13
                Error:          Not equal:
                                expected: "Halo Kim Teari"
                                actual  : "Halo Kim Taeri"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -Halo Kim Teari
                                +Halo Kim Taeri
                Test:           TestSubtest/Taeri
                Messages:       Seharusnya `Halo Kim Taeri`
--- FAIL: TestSubtest (0.00s)
    --- FAIL: TestSubtest/Taeri (0.00s)
FAIL
After Testing
exit status 1
FAIL    learn-go-unit-test/helper       0.726s
*/
