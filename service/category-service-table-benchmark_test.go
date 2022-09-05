package service

import (
	"learn-go-unit-test/helper"
	"testing"
)

func BenchmarkTableHelloRasyidev(b *testing.B) {
	benchmarks := []struct {
		name, request string
	}{
		{name: "Rasyidev", request: "Rasyidev"},
		{name: "Taeri", request: "Kim Taeri"},
		{name: "Suzy", request: "Bae Suzy"},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			helper.HelloRasyidev(benchmark.request)
		})
	}
}

/*
$ go test -v -run AplikasiNone -bench BenchmarkTableHelloRasyidev
goos: windows
goarch: amd64
pkg: learn-go-unit-test/service
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkTableHelloRasyidev
BenchmarkTableHelloRasyidev/Rasyidev
BenchmarkTableHelloRasyidev/Rasyidev-8          1000000000
BenchmarkTableHelloRasyidev/Taeri
BenchmarkTableHelloRasyidev/Taeri-8             1000000000
BenchmarkTableHelloRasyidev/Suzy
BenchmarkTableHelloRasyidev/Suzy-8              1000000000
PASS
ok      learn-go-unit-test/service      0.064s
*/
