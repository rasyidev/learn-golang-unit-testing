package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rasyidev")
	if result != "Halo Rasyidev" {
		// unit test gagal
		panic("Result is not Rasyidev")
	}
}
