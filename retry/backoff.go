package retry

import "time"

func ConstantBackoff(n int, amount time.Duration) []time.Duration {
	ts := make([]time.Duration, n)
	for i := 0; i < n; i++ {
		ts[i] = amount
	}
	return ts
}

func ExponentialBackoff(n int, initialAmount time.Duration) []time.Duration {
	return ExponentialLimitBackoff(n, initialAmount, -1)
}

func ExponentialLimitBackoff(n int, initialAmount, maxAmount time.Duration) []time.Duration {
	ts := make([]time.Duration, n)
	next := initialAmount
	for i := 0; i < n; i++ {
		ts[i] = next
		next *= 2
		if maxAmount > 0 && next > maxAmount {
			next = maxAmount
		}
	}
	return ts
}
