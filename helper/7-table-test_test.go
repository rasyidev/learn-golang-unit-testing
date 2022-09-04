package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRasyidevTabletest(t *testing.T) {
	testCases := []struct {
		name, request, expected string
	}{
		{
			name:     "Rasyidev",
			request:  "Rasyidev",
			expected: "Halo Rasyidev",
		},
		{
			name:     "Taeri",
			request:  "Kim Teari",
			expected: "Halo Kim Taeri",
		},
		{
			name:     "Suzy",
			request:  "Bae Suzy",
			expected: "Halo Bae Suzy",
		},
		{
			name:     "Yoona",
			request:  "Im Yoona",
			expected: "Halo Im Yoona",
		},
		{
			name:     "Youngwoo",
			request:  "Woo Young Woo",
			expected: "Halo Woo Young Woo",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := HelloRasyidev(tc.request)
			require.Equal(t, tc.expected, result, "Seharusnya "+tc.expected)
		})
	}
}

/*
$ go test -v -run TestRasyidevTabletest
Before Testing
=== RUN   TestRasyidevTabletest
=== RUN   TestRasyidevTabletest/Rasyidev
=== RUN   TestRasyidevTabletest/Taeri
    7-table-test_test.go:43:
                Error Trace:    C:\Docs\App Repositories\learn-go-unit-test\helper\7-table-test_test.go:43
                Error:          Not equal:
                                expected: "Halo Kim Taeri"
                                actual  : "Halo Kim Teari"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -Halo Kim Taeri
                                +Halo Kim Teari
                Test:           TestRasyidevTabletest/Taeri
                Messages:       Seharusnya Halo Kim Taeri
=== RUN   TestRasyidevTabletest/Suzy
=== RUN   TestRasyidevTabletest/Yoona
=== RUN   TestRasyidevTabletest/Youngwoo
--- FAIL: TestRasyidevTabletest (0.00s)
    --- PASS: TestRasyidevTabletest/Rasyidev (0.00s)
    --- FAIL: TestRasyidevTabletest/Taeri (0.00s)
    --- PASS: TestRasyidevTabletest/Suzy (0.00s)
    --- PASS: TestRasyidevTabletest/Yoona (0.00s)
    --- PASS: TestRasyidevTabletest/Youngwoo (0.00s)
FAIL
After Testing
exit status 1
FAIL    learn-go-unit-test/helper       0.816s
*/
