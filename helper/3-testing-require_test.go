package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloRasyidevRequire(t *testing.T) {
	res := HelloRasyidev("Taeri")
	require.Equal(t, "Halo Rasyidev", res, "Harusnya `Halo Rasyidev`")
	fmt.Println("Test menggunakan require selesai. Tapi harusnya pesan ini gak muncul karena langsung manggil `FailNow()`")

}

/*
$ go test -run TestHelloRasyidevRequire
--- FAIL: TestHelloRasyidevRequire (0.00s)
    3-testing-require_test.go:12:
                Error Trace:    learn-go-unit-test\helper\3-testing-require_test.go:12
                Error:          Not equal:
                                expected: "Halo Rasyidev"
                                actual  : "Halo Taeri"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -Halo Rasyidev
                                +Halo Taeri
                Test:           TestHelloRasyidevRequire
                Messages:       Harusnya `Halo Rasyidev`
FAIL
exit status 1
FAIL    learn-go-unit-test/helper       0.782s
*/
