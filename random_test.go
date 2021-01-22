package random_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/frozentech/random"
)

func ExampleString() {
	fmt.Println(regexp.MatchString(`[0-9a-zA-Z]{50}`, random.String(50, []rune(random.RuneAlNumCS))))
	// Output:
	// true <nil>
}

func BenchmarkRandomString(b *testing.B) {
	for n := 0; n < 100000; n++ {
		random.String(50, []rune(random.RuneAlNumCS))
	}
}
