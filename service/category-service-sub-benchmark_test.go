package service

import (
	"learn-go-unit-test/helper"
	"testing"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Rasyidev", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = helper.HelloRasyidev("Rasyidev")
		}
	})

	b.Run("Taeri", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = helper.HelloRasyidev("Kim Taery Beauty")
		}
	})

}

/*
$ go test -v -run AplikasiTidakAda -bench BenchmarkSub/Taeri
goos: windows
goarch: amd64
pkg: learn-go-unit-test/service
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkSub
BenchmarkSub/Taeri
BenchmarkSub/Taeri-8            77722220                13.03 ns/op
PASS
ok      learn-go-unit-test/service      1.072s
*/
