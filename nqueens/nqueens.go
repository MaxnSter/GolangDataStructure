package nqueens

import (
	"bufio"
	"fmt"
	"os"
)

// 用一个一维数组表示,数组下标表示行,数组元素值表示列
// 于是,使用i代表下标Index,x代表对应元素值,x=-1代表为未存放元素,可以有:
// 行冲突:x!=-1,即已经该行已经存放元素了
// 列冲突:数组中所有元素的值不重复,也就是x唯一
// 对角线冲突:斜线上冲突的皇后的位置都有规律即它们所在的行列互减的绝对值相等,也就是abs(x[i] - i)唯一
// 解法1:非递归回溯法
func NQueens(n int) (solutions int) {
	//https://zh.wikipedia.org/wiki/%E5%85%AB%E7%9A%87%E5%90%8E%E9%97%AE%E9%A2%98
	//FIXME maxN
	if !(n == 1 || n >= 4) {
		return 0
	}

	N := make([]int, n)
	reset(N)
	row, col := 0, 0
	for row < n {
		//每一行,对它的每一列探测,看是否可以放置皇后
		for col < n {
			if valid(N, row, col) {
				//成功放置一行
				N[row] = col
				//下一行的探测当然也从0开始
				col = 0
				break
			} else {
				//继续探测下一列
				col++
			}
		}

		//该行探测所有列之后还没有找到放置的为止
		if N[row] == -1 {
			//回溯
			if row == 0 {
				//回溯到了第一行并且第一行无解,说明程序已经找到了所有的解答
				break
			} else {
				//回到上一行
				row--
				//上一行皇后的位置右移一列
				col = N[row] + 1
				//右移一列后,上一行重新探测
				N[row] = -1
				continue
			}
		}

		if row == n-1 {
			//最后一行探测成功,说明我们已经找到了一个解答
			solutions++
			//fmt.Printf("answer %d : \n", solutions)
			//printResult(N)

			col = N[row] + 1 //从最后一行放置皇后列数的下一列继续探测
			N[row] = -1      //最后一行重新探测,由此一直回溯直到break
			continue
		}

		//该行探测成功,但是不是最后一行,
		//对下一行进行探测,探测列同样从0开始
		row++
	}

	return solutions
}

// 打印结果
func printResult(N []int) {
	wr := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(N); i++ {
		for j := 0; j < len(N); j++ {
			if N[i] != j {
				//上面没有棋子
				wr.WriteByte('.')
				wr.WriteByte(' ')
			} else {
				//上面有棋子
				wr.WriteByte('#')
			}
		}
		wr.WriteByte('\n')
	}

	for _, v := range N {
		wr.WriteString(fmt.Sprintf("%d ", v))
	}
	wr.WriteByte('\n')

	wr.WriteString("--------------------------\n")
	wr.Flush()
}

func reset(N []int) {
	for i := 0; i < len(N); i++ {
		N[i] = -1
	}
}

// row:要存放在第几行, col:要存放在第几列
func valid(N []int, row int, col int) bool {
	//只需要和前row行(已经放置了皇后的行)比较
	for i := 0; i < row; i++ {
		iRow, iCol := i, N[i]
		//列冲突或对角冲突
		if iCol == col || absInt(col-iCol) == absInt(row-iRow) {
			return false
		}
	}
	return true
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func NQueensBits(n int) int {
	if !(n == 1 || n >= 4) {
		return 0
	}

	sum := 0
	upperLimit := uint((1 << uint(n)) - 1)
	backTrace(0, 0, 0, upperLimit, &sum)
	return sum
}

func backTrace(row uint, ld uint, rd uint, upperLimit uint, sum *int) {
	if row != upperLimit {
		pos := upperLimit & ^(row | ld | rd)
		for pos != 0 {
			//拷贝pos最右边为1的bit,其余bit置为0
			//也就是取得可以放皇后的最后的列
			p := pos & (-pos)

			//将pos最右边为1的bit清零
			//也就是为获取下一次最右边可使用列做准备
			//程序将来会回溯到这个为止继续试探
			pos -= p

			//row+p,将当前列置为1,表示记录这次皇后放置的列
			//(ld + p) << 1,标志当前皇后左边相邻的列不允许下一个皇后放置
			//(rd + p) >> 1,标志当前皇后右边相邻的列不允许下一个皇后放置
			backTrace(row+p, (ld+p)<<1, (rd+p)>>1, upperLimit, sum)
		}
	} else {
		*sum++
	}
}
