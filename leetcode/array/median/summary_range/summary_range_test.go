package summary_range

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, []string{"0->2", "4->5", "7"}, summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	assert.Equal(t, []string{"0", "2->4", "6", "8->9"}, summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
