package retry

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConstantBackoff(t *testing.T) {
	initialAmount := time.Second
	n := 10
	expect := make([]time.Duration, n)
	for i := 0; i < n; i++ {
		expect[i] = initialAmount
	}

	actual := ConstantBackoff(n, initialAmount)
	assert.EqualValues(t, expect, actual)
}

func TestExponentialBackoff(t *testing.T) {
	initialAmount := time.Second
	n := 10
	expect := make([]time.Duration, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			expect[i] = initialAmount
			continue
		}

		expect[i] = expect[i - 1] * 2
	}

	actual := ExponentialBackoff(n, initialAmount)
	assert.EqualValues(t, expect, actual)
}