package pipe_filter

import "testing"

func TestStraightPipeLine(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sumfilter := NewSumFilter()

	sp := NewStraightPipeLine("p1", spliter, converter, sumfilter)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("The expected is 6,but the actual is %d", ret)
	}
	t.Log(ret)
}
