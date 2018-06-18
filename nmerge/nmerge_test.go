package nmerge

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

type source struct {
	data   []int
	curIdx int
}

func (s *source) Less(oth MergeSource) bool {
	return s.curVal() < oth.(*source).curVal()
}

func (s *source) Next() bool {
	if s.curIdx == len(s.data)-1 {
		return false
	}

	//更新状态,很重要!
	s.curIdx++
	return true
}

func (s *source) curVal() int {
	return s.data[s.curIdx]
}

func (s *source) OutputTo(to interface{}) {
	output := to.(*[]int)
	*output = append(*output, s.curVal())
}

func TestMergeMulti(t *testing.T) {
	ss := make([]MergeSource, 0)
	ss = append(ss, &source{
		data:   []int{1, 5, 12},
		curIdx: 0,
	})

	ss = append(ss, &source{
		data:   []int{2, 6, 11},
		curIdx: 0,
	})

	ss = append(ss, &source{
		data:   []int{3, 7, 10},
		curIdx: 0,
	})

	ss = append(ss, &source{
		data:   []int{4, 8, 9},
		curIdx: 0,
	})

	output := make([]int, 0)
	Merge(ss, &output)
	assert.Equal(t, len(output), 12)
	assert.Equal(t, output, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
}

func TestMergeSingle(t *testing.T) {
	ss := make([]MergeSource, 0)
	ss = append(ss, &source{
		data:   []int{1, 5, 12},
		curIdx: 0,
	})

	output := make([]int, 0)
	Merge(ss, &output)
	assert.Equal(t, len(output), 3)
	assert.Equal(t, output, []int{1, 5, 12})
}

func TestMergeEmptySource(t *testing.T) {
	ss := make([]MergeSource, 0)
	output := make([]int, 0)
	Merge(ss, &output)
	assert.Equal(t, len(output), 0)
	assert.Equal(t, output, []int{})
}
