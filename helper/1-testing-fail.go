package helper

import "errors"

func Pembagian(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Tidak bisa membagi dengan angka nol")
	}

	return a / b, nil
}
