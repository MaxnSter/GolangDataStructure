package basic

//1.冒泡排序
func bubbleSort(a []int) {
	for i := len(a) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}

		//flag := false
		//for j := 1; j <= i; j++ {
		//	if a[j-1] > a[j] {
		//		a[j-1], a[j] = a[j], a[j-1]
		//
		//		flag = true
		//	}
		//}
		//if !flag {
		//	break
		//}
	}
}

//2.插入排序
func insertSort(a []int) {
	var insert, j int
	for i := 1; i < len(a); i++ {
		insert = a[i]
		for j = i; j > 0 && a[j-1] > insert; j-- {
			//大于插入点的元素都往右移动
			a[j] = a[j-1]
		}

		a[j] = insert
	}
}

//3.选择排序
func selectSort(a []int) {
	var minIdx int
	for i := 0; i < len(a); i++ {
		minIdx = i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}

		a[i], a[minIdx] = a[minIdx], a[i]
	}
}

//4.希尔排序
func shellSort(a []int) {

	var inc, j int

	for inc = len(a) / 2; inc > 0; inc /= 2 {
		for i := inc; i < len(a); i++ {
			insert := a[i]

			for j = i; j >= inc && a[j-inc] > insert; j -= inc {
				a[j] = a[j-inc]
			}

			a[j] = insert
		}
	}
}

//5.堆排序
func heapSort(a []int) {
	N := len(a)

	//构建最大堆
	for i := (N / 2) - 1; i >= 0; i-- {
		percDown(a, i, N)
	}

	for i := N - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0] //删除最大元素
		percDown(a, 0, i)       //重新构建
	}

}

func percDown(a []int, i int, N int) {
	//
	//            1                        3
	//          /   \         ->         /   \
	//        /      \                 /      \
	//       2        3               2        1

	var child, tmp int

	for tmp = a[i]; i*2+1 < N; i = child {
		child = i*2 + 1
		if child+1 < N && a[child+1] > a[child] {
			child++
		}

		if a[child] > tmp {
			a[i] = a[child]
		} else {
			break
		}
	}

	a[i] = tmp
}

//6.归并排序
func mergeSort(a []int) {
	if len(a) <= 1 {
		return
	}

	mid := (len(a) - 1) / 2
	mergeSort(a[:mid+1]) // [start,mid]
	mergeSort(a[mid+1:]) // [mid+1,end]
	merge(a, mid)
}

func merge(a []int, mid int) {
	lPos, lEnd := 0, mid
	rPos, rEnd := mid+1, len(a)-1
	tmp, tPos := make([]int, len(a)), lPos

	for lPos <= lEnd && rPos <= rEnd {
		if a[lPos] < a[rPos] {
			tmp[tPos] = a[lPos]
			lPos++
		} else {
			tmp[tPos] = a[rPos]
			rPos++
		}
		tPos++
	}

	for lPos <= lEnd {
		tmp[tPos] = a[lPos]
		lPos++
		tPos++
	}

	for rPos <= rEnd {
		tmp[tPos] = a[rPos]
		rPos++
		tPos++
	}

	copy(a, tmp)
}

//7.快速排序
func quickSort(a []int) {
	if len(a) <= 3 {
		insertSort(a)
		return
	}

	left, right := 0, len(a)-1
	pivot := median3(a, left, right)
	i, j := left, right-1

	for {
		// for a[++i] < pivot {}
		for i++; a[i] < pivot; {
			i++
		}

		// for a[--j] > pivot {}
		for j--; a[j] > pivot; {
			j--
		}

		if i < j {
			//交换后,a[i]<pivot,a[j]>pivot
			a[i], a[j] = a[j], a[i]
		} else {
			break
		}
	}
	//此时i的位置正好是第一个大于pivot的位置,与pivot交换
	//得到两个区域, a[i]==pivot
	//[left,i-1]区域都小于pivot
	//[i+1,right]区域都大于pivot
	a[i], a[right-1] = a[right-1], a[i]
	quickSort(a[:i])   //[left,i-1]
	quickSort(a[i+1:]) //[i+1,right]
}

func median3(a []int, left, right int) int {
	mid := left + (right-left)/2

	//a[left] < <= a[mid] <= a[right]
	if a[left] > a[mid] {
		a[left], a[mid] = a[mid], a[left]
	}

	if a[left] > a[right] {
		a[left], a[right] = a[right], a[left]
	}

	if a[mid] > a[right] {
		a[mid], a[right] = a[right], a[mid]
	}

	//hide pivot
	a[mid], a[right-1] = a[right-1], a[mid]
	return a[right-1]
}
