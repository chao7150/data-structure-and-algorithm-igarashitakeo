package main

import (
	"testing"
)

// [
// 	[0, 3, 1],
// 	[Inf, 0, Inf],
// 	[Inf, 1, 0]
// ]

func TestDijkstraShortestPath(t *testing.T) {
	m := Matrix{{0, 3, 1}, {Inf, 0, Inf}, {Inf, 1, 0}}
	costTable := DijkstraShortestPath(m, 0)
	res := []int{0, 2, 1}
	for i, v := range costTable {
		if v.distance != res[i] {
			t.Fail()
		}
	}
	m = Matrix{{0, 2, 9, Inf, Inf, Inf}, {Inf, 0, Inf, 3, Inf, Inf}, {Inf, 7, 0, Inf, 2, Inf}, {Inf, Inf, 3, 0, 8, 9}, {Inf, Inf, Inf, Inf, 0, 1}, {Inf, Inf, Inf, Inf, Inf, 0}}
	costTable = DijkstraShortestPath(m, 0)
	res = []int{0, 2, 8, 5, 10, 11}
	for i, v := range costTable {
		if v.distance != res[i] {
			t.Fail()
		}
	}
}

func TestAllFixed(t *testing.T) {
	m := DistanceMap{{0, 0, true}, {0, 0, true}, {0, 0, true}}
	res := AllFixed(m)
	if !res {
		t.Fail()
	}
	m = DistanceMap{{0, 0, true}, {0, 0, true}, {0, 0, false}}
	res = AllFixed(m)
	if res {
		t.Fail()
	}
}

func TestGetNearestItem(t *testing.T) {
	res := getNearestUnfixedNode(DistanceMap{{2, 0, false}, {3, 1, false}, {1, 2, false}})
	if res != 2 {
		t.Fail()
	}
	res = getNearestUnfixedNode(DistanceMap{{2, 0, true}, {3, 1, false}, {1, 2, true}})
	if res != 1 {
		t.Fail()
	}
}