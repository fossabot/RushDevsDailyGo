package Daily21

import (
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/Math"
	"github.com/Synertry/GoSysUtils/Slice"
	"github.com/google/go-cmp/cmp"
)

var sliceOfString []string

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []string
	}{
		"intro":     {input: []string{"tech", "lead", "techlead", "cat", "cats", "dog", "catsdog"}, want: []string{"techlead", "catsdog"}},
		"letters":   {input: []string{"a", "b", "ab", "abd"}, want: []string{"ab"}},
		"noConcats": {input: []string{"a", "ab", "c"}, want: []string{}},
		"empty":     {input: []string{}, want: []string{}},
		"leetcode":  {input: []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}, want: []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := FindAllConcatenatedWordsInADict(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
				t.Fatalf(diff)
			}
		})
	}
}

func BenchmarkFindAllConcatenatedWordsInADict(b *testing.B) {
	type benchmark struct {
		name string
		len  int
	}

	maxExpDictLen := 4
	benchmarks := make([]benchmark, maxExpDictLen+2) // + 2 for empty floor(10^-1) and single 10^0 -> 1

	for i := -1; i <= maxExpDictLen; i++ { // -1 as start, because substraction is more costly than addition
		dictLen := Math.IntPow(10, i)
		benchmarks[i+1] = benchmark{name: "DictLen10^" + strconv.Itoa(i), len: dictLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			input, result := Slice.GenRandomStrings(bm.len), make([]string, bm.len)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = FindAllConcatenatedWordsInADict(input)
			}
			sliceOfString = result
		})
	}
}
