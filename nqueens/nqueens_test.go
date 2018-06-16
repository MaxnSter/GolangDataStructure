package nqueens

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

/*
	https://zh.wikipedia.org/wiki/%E5%85%AB%E7%9A%87%E5%90%8E%E9%97%AE%E9%A2%98
	n	1	2	3	4	5	6	7	8	9	10	11		12		13		14	..
	U	1	0	0	1	2	1	6	12	46	92	341		1,787	9,233	45,752	..
	D	1	0	0	2	10	4	40	92	352	724	2,680	14,200	73,712	365,596	..
 */

func TestNQueens(t *testing.T) {
	assert.Equal(t, NQueens(1), 1)
	assert.Equal(t, NQueens(2), 0)
	assert.Equal(t, NQueens(3), 0)
	assert.Equal(t, NQueens(4), 2)
	assert.Equal(t, NQueens(5), 10)
	assert.Equal(t, NQueens(6), 4)
	assert.Equal(t, NQueens(7), 40)
	assert.Equal(t, NQueens(8), 92)
	assert.Equal(t, NQueens(9), 352)
	assert.Equal(t, NQueens(10), 724)
}

func TestNQueensBitsBits(t *testing.T) {
	assert.Equal(t, NQueensBits(1), 1)
	assert.Equal(t, NQueensBits(2), 0)
	assert.Equal(t, NQueensBits(3), 0)
	assert.Equal(t, NQueensBits(4), 2)
	assert.Equal(t, NQueensBits(5), 10)
	assert.Equal(t, NQueensBits(6), 4)
	assert.Equal(t, NQueensBits(7), 40)
	assert.Equal(t, NQueensBits(8), 92)
	assert.Equal(t, NQueensBits(9), 352)
	assert.Equal(t, NQueensBits(10), 724)
}

func BenchmarkNQueens(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NQueens(8)
	}
}

func BenchmarkNQueensBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NQueensBits(8)
	}
}
