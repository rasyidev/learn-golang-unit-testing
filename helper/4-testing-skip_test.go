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
