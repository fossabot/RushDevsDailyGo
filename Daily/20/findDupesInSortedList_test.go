package Daily20

import (
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/Synertry/GoSysUtils/Math"
	"github.com/google/go-cmp/cmp"
)

var (
	sliceOfInts []int
	random      = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func TestFindDupesInSortedList(t *testing.T) {
	tests := map[string]struct {
		input  []int
		target int
		want   []int
	}{
		"intro1":   {input: []int{1, 1, 3, 5, 7}, target: 1, want: []int{0, 1}},
		"intro2":   {input: []int{1, 2, 3, 4}, target: 5, want: []int{-1, -1}},
		"single 0": {input: []int{0}, target: 0, want: []int{0, -1}},
		"empty":    {input: []int{}, target: 1, want: []int{-1, -1}},
	}

	for inputLen := 3; inputLen < 5; inputLen++ {
		random.Seed(time.Now().UnixNano()) // ensure pseudo-randomness
		input, target, want := make([]int, inputLen), -1, []int{-1, -1}

		targetIdx := random.Intn(inputLen - 2) // inputLen = 3 -> 1 -> rangeIdx [0, 1]
		logn := int(math.Log2(float64(inputLen)))
		for i, idxV := 0, 0; i < inputLen; i++ {
			input[i] = idxV
			if i == targetIdx { // add duplicate here
				want = []int{i, i + 1}
				i++
				input[i] = idxV
				target = idxV
			}
			idxV = idxV + logn // increase by logn to simulate missing values
		}

		tests["random"+strconv.Itoa(inputLen)] = struct {
			input  []int
			target int
			want   []int
		}{input: input, target: target, want: want}
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := FindDupesInSortedList(tc.input, tc.target)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("target, %d, expected: %#v, got: %#v", tc.target, tc.want, got)
				t.Log(diff)
				t.Logf("input: %#v\n", tc.input)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func BenchmarkFindDupesInSortedList(b *testing.B) {
	maxExpArrLen := 4
	type benchmark struct {
		name string
		len  int
	}

	benchmarks := make([]benchmark, maxExpArrLen+1)    // do not use maps! Order will be randomized; + 1 for 2^0
	benchmarks[0] = benchmark{name: "ArrLen3", len: 3} // start case

	for i := 1; i <= maxExpArrLen; i++ {
		arrLen := Math.IntPow(10, i)
		benchmarks[i] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			random.Seed(time.Now().UnixNano())
			input, target, result := random.Perm(bm.len), random.Intn(bm.len-2), make([]int, 2)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = FindDupesInSortedList(input, target)
			}
			sliceOfInts = result
		})
	}
}
