package fan_in

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
