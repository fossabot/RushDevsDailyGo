package Daily16

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/Synertry/GoSysUtils/Math"
	"github.com/google/go-cmp/cmp"
)

type benchmark struct {
	name string
	len  int
}

const maxExpArrLen = 1

var (
	resultSliceOfInts2D [][]int
	random              = rand.New(rand.NewSource(time.Now().UnixNano()))
	benchmarks          = make([]benchmark, maxExpArrLen+3) // do not use maps! Order will be randomized; + 3 for empty (10^-1), 10^0 and ArrLen20
)

func init() {
	// generate benchmark data
	for i := -1; i <= maxExpArrLen; i++ {
		arrLen := Math.IntPow(10, i)
		benchmarks[i+1] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}
	benchmarks[maxExpArrLen+2] = benchmark{name: "ArrLen20", len: 20}
}

func TestGenerateAllSubsets(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  [][]int
	}{
		"intro":  {input: []int{1, 2, 3}, want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		"single": {input: []int{1}, want: [][]int{{}, {1}}},
		"empty":  {input: []int{}, want: [][]int{{}}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := GenerateAllSubsets(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
				t.Log(diff)
				t.Logf("input: %#v\n", tc.input)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func BenchmarkGenerateAllSubsets(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			random.Seed(time.Now().UnixNano())
			input := random.Perm(bm.len)
			var result [][]int
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = GenerateAllSubsets(input)
			}
			resultSliceOfInts2D = result
		})
	}
}