package merge_sort

import "sync"

func MergeSort(a []int) {
	if len(a) <= 1 {
		return
	}

	MergeSort(a[:len(a)/2])
	MergeSort(a[len(a)/2:])
	merge(a, len(a)/2)
}

func merge(a []int, mid int) {
	tmp, i := make([]int, len(a)), 0
	lb, le := 0, mid-1
	rb, re := mid, len(a)-1

	for lb < le && rb < re {
		if a[lb] < a[rb] {
			tmp[i] = a[lb]
			lb++
		} else {
			tmp[i] = a[rb]
			rb++
		}

		i++
	}

	for lb < le {
		tmp[i] = a[lb]
		lb++
		i++
	}

	for rb < re {
		tmp[i] = a[rb]
		rb++
		i++
	}

	copy(a, tmp)
}

const (
	max = 1 << 11
)

func MergeSortParallel(a []int) {
	if len(a) <= 1 {
		return
	}

	if len(a) <= max {
		MergeSort(a)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		MergeSortParallel(a[:len(a)/2])
	}()

	MergeSortParallel(a[len(a)/2:])

	wg.Wait()
	merge(a, len(a)/2)
}
