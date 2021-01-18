package example

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer int
}

func TestSumTable(t *testing.T) {
	tests := []test{
		{[]int{1, 2, 3}, 6},
		{[]int{10, 11, 12}, 33},
		{[]int{5, 6}, 11},
	}

	for _, v := range tests {
		x := Sum(v.data...)
		if x != v.answer {
			t.Error("expected ", v.answer, ", got:", x)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 2, 3))
	//Output: 6
}

func TestSum(t *testing.T) {
	x := Sum(1, 2, 3)

	if x != 6 {
		t.Error("expected 6, got:", x)
	}
}

func TestSumFail(t *testing.T) {
	x := Sum(1, 2, 3, 4)

	if x == 6 {
		t.Error("expected 10, got:", x)
	}
}

func BenchmarkSum(b *testing.B) {
	// tests := []test{
	// 	{[]int{1, 2, 3}, 6},
	// 	{[]int{10, 11, 12}, 33},
	// 	{[]int{5, 6}, 11},
	// }

	// for i := 0; i < b.N; i++ {
	// 	for _, v := range tests {
	// 		Sum(v.data...)
	// 	}
	// }

	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}
