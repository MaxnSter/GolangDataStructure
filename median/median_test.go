package median

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMedianFinder_AddNum(t *testing.T) {
	m := MedianFinder{}
	m.AddNum(1)
	m.AddNum(2)
	m.AddNum(3)
	m.AddNum(4)

	assert.Equal(t, m.FindMedian(), (2+3)/float64(2))
}

func TestMedianFinder_AddNum2(t *testing.T) {

	m := MedianFinder{}
	m.AddNum(1)
	m.AddNum(2)
	m.AddNum(3)
	m.AddNum(4)
	m.AddNum(5)

	assert.Equal(t, m.FindMedian(), float64(3))
}
