package helper

import (
	"fmt"
	"testing"
)

func TestCheckOS(t *testing.T) {
	currentOS := CheckOS()
	if currentOS == "windows" {
		t.Skip("Skip test untuk operasi sistem Windows")
	}
	fmt.Println("Selesai menjalankan test CheckOS")
}

/*
$ go test -v -run TestCheckOS
=== RUN   TestCheckOS
windows
    4-testing-skip_test.go:11: Skip test untuk operasi sistem Windows
--- SKIP: TestCheckOS (0.00s)
PASS
ok      learn-go-unit-test/helper       0.821s
*/
