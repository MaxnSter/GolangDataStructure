package sort_colors

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	a := []int{2,0,1,1,1,1,1,1,2,0,0,0,0,2,2,2}
	sortColors(a)
	fmt.Println(a)
}
