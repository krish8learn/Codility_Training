package main

import "testing"

func TestMissinteger(t *testing.T) {
	type cases struct {
		data   []int
		answer int
	}

	test := []cases{
		cases{[]int{1, 3, 6, 4, 1, 2}, 5},
		cases{[]int{-1, -3}, 1},
		cases{[]int{1, 2, 3}, 4},
		cases{[]int{1},2},
		cases{[]int{4,2,5,3},1},
		cases{[]int{2},1},
		cases{[]int{23, 21},1},
		cases{[]int{1,23,22},2},
		cases{[]int{1,2,3,23,22},4},
		cases{[]int{0},1},
		cases{[]int{4,1,5,6,2},3},
		cases{[]int{-2, 4, 1, -9, 3},2},
		cases{[]int{-1000000, 1000000},1},
		cases{[]int{-3,4},1},
	}

	for _, val := range test {
		result := Missinteger(val.data)
		if val.answer != result {
			t.Error("For:", val.data, "Expected:", val.answer, "got:", result)
		}
	}
}
