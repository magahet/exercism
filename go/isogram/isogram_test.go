package isogram

import "testing"

func TestIsIsogramNested(t *testing.T) {
	for _, c := range testCases {
		if IsIsogramNested(c.input) != c.expected {
			t.Fatalf("FAIL: %s\nWord %q, expected %t, got %t", c.description, c.input, c.expected, !c.expected)
		}

		t.Logf("PASS: Word %q", c.input)
	}
}

func BenchmarkIsIsogramNested(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, c := range testCases {
			IsIsogramNested(c.input)
		}

	}
}

func BenchmarkIsIsogramSet(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, c := range testCases {
			IsIsogramSet(c.input)
		}

	}
}
