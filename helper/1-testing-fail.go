package helper

import (
	"errors"
	"fmt"
	"runtime"
)

func Pembagian(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa membagi dengan angka nol")
	}

	return a / b, nil
}

func HelloRasyidev(name string) string {
	return "Halo " + name
}

func CheckOS() string {
	fmt.Println(runtime.GOOS)
	return runtime.GOOS
}
