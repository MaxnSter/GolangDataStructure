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
	ts := make([]time.Duration, n)
	next := initialAmount
	for i := 0; i < n; i++ {
		ts[i] = next
		next *= 2
	}
	return ts
}
