package algorithm

import (
	"testing"
)

func TestWaysMoving(t *testing.T) {
	n := WaysMoving(0, 10)
	t.Logf("WaysMoving : %v", n)
}

//func TestWaysMovingRC(t *testing.T) {
//	n := WaysMovingRC(10)
//	t.Logf("WaysMovingRC : %v", n)
//}

func TestWaysEvenCount(t *testing.T) {
	n := WaysMovingByEvenCount(10)
	t.Logf("WaysMovingByEvenCount : %v", n)
}

func BenchmarkWaysMoving(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WaysMoving(0, 10)
	}
}

func BenchmarkWaysMovingByEvenCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WaysMovingByEvenCount(10)
	}
}

func TestMaxAscSubSeq(t *testing.T) {
	m := MaxAscSubSeq([]int{1, 7, 3, 6, 9, 8})
	if m != 4 {
		t.Errorf("MaxAscSubSeq Not Correct")
	} else {
		t.Logf("MaxAscSubSeq: %v, Success", m)
	}
}

//func TestShuffleReverse(t *testing.T) {
//	n := ShuffleReverse([]int8{1, 2, 3, 4})
//	if n != 4 {
//		t.Errorf("ShuffleReverse Not Correct")
//	} else {
//		t.Logf("ShuffleReverse: %v, Success", n)
//	}
//}

func TestReverse(t *testing.T) {
	list := ConstructList([]int{1, 2, 3, 4, 5, 6, 7})
	rList := Reverse(list)
	t.Log(FormatList(rList))
}

func TestReorderList(t *testing.T) {
	list := ConstructList([]int{1, 2, 3, 4, 5})
	t.Log(FormatList(list))
	rList := ReorderList(list)
	t.Log(FormatList(rList))
}
