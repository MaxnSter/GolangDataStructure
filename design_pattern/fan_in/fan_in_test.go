package fan_in

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFanInWithPicker(t *testing.T) {
	// 1324...1324(100 times)12
	var expect []int
	for i := 0; i < 100; i++ {
		expect = append(expect, 1)
		expect = append(expect, 3)
		expect = append(expect, 2)
		expect = append(expect, 4)
	}
	expect = append(expect, []int{1, 2}...)
	sendF := func(ch chan interface{}, val interface{}) {
		num := 100
		if val.(int) == 1 || val.(int) == 2 {
			num += 1
		}
		for i := 0; i < num; i++ {
			ch <- val
		}
		close(ch)
	}

	ch1 := make(chan interface{})
	go sendF(ch1, 1)

	ch2 := make(chan interface{})
	go sendF(ch2, 2)

	ch3 := make(chan interface{})
	go sendF(ch3, 3)

	ch4 := make(chan interface{})
	go sendF(ch4, 4)

	pickF := func() OnPick {
		return func(idx []int) []int {
			var a0 []int
			var a1 []int
			var a2 []int
			var a3 []int

			for _, i := range idx {
				switch i {
				case 0:
					a0 = append(a0, i)
				case 1:
					a1 = append(a1, i)
				case 2:
					a2 = append(a2, i)
				case 3:
					a3 = append(a3, i)
				}
			}

			// 1324
			return append(append(append(a0, a2...), a1...), a3...)
		}
	}

	var actual []int
	out := FanInWithPicker(pickF(), ch1, ch2, ch3, ch4)
	for data := range out {
		actual = append(actual, data.(int))
	}

	assert.EqualValues(t, expect, actual)
}

func TestFanInWithPicker2(t *testing.T) {
	// goog....goog
	var expect []string
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			expect = append(expect, "g")
			expect = append(expect, "o")
		} else {
			expect = append(expect, "o")
			expect = append(expect, "g")
		}
	}
	sendF := func(ch chan interface{}, val interface{}) {
		num := 100
		for i := 0; i < num; i++ {
			ch <- val
		}
		close(ch)
	}

	ch1 := make(chan interface{})
	go sendF(ch1, "g")

	ch2 := make(chan interface{})
	go sendF(ch2, "o")

	i := 0
	picker := func() OnPick {
		return func(available []int) []int {
			defer func() {
				i++
			}()

			if i%2 == 0 {
				return []int{0, 1}
			} else {
				return []int{1, 0}
			}
		}
	}

	var actual []string
	out := FanInWithPicker(picker(), ch1, ch2)
	for data := range out {
		actual = append(actual, data.(string))
	}

	assert.EqualValues(t, expect, actual)
}

func TestFanInPriority(t *testing.T) {
	var expect []int
	for i := 0; i < 100; i++ {
		expect = append(expect, 4)
		expect = append(expect, 3)
		expect = append(expect, 2)
		expect = append(expect, 1)
	}
	sendF := func(ch chan interface{}, val interface{}) {
		num := 100
		for i := 0; i < num; i++ {
			ch <- val
		}
		close(ch)
	}

	ch1 := make(chan interface{})
	p1 := PriorityCh{
		Ch:       ch1,
		Priority: 1,
	}
	go sendF(ch1, 1)

	ch2 := make(chan interface{})
	p2 := PriorityCh{
		Ch:       ch2,
		Priority: 2,
	}
	go sendF(ch2, 2)

	ch3 := make(chan interface{})
	p3 := PriorityCh{
		Ch:       ch3,
		Priority: 3,
	}
	go sendF(ch3, 3)

	ch4 := make(chan interface{})
	p4 := PriorityCh{
		Ch:       ch4,
		Priority: 4,
	}
	go sendF(ch4, 4)

	var actual []int
	out := FanInWithPriority(p1, p2, p3, p4)
	for data := range out {
		actual = append(actual, data.(int))
	}

	assert.EqualValues(t, expect, actual)
}
