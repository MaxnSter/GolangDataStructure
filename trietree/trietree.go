package trietree

import (
	"bufio"
	"errors"
	"io"
)

type trieNode struct {
	isEnd      bool
	childNodes map[byte]*trieNode
}

type trieTree struct {
	root *trieNode
}

func NewTrieTree() *trieTree {
	return &trieTree{
		root: &trieNode{
			isEnd:      false,
			childNodes: make(map[byte]*trieNode),
		}}
}

func BuildTrieTree(rd io.Reader) *trieTree {
	t := NewTrieTree()
	scan := bufio.NewScanner(rd)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		t.AddWord(scan.Text())
	}

	return t
}

func (t *trieTree) AddWord(word string) error {
	if len(word) == 0 {
		return errors.New("empty word")
	}

	tmp := t.root
	for i := 0; i < len(word); i++ {
		b := word[i]

		if _, ok := tmp.childNodes[b]; !ok {
			tmp.childNodes[b] = &trieNode{
				isEnd:      false,
				childNodes: make(map[byte]*trieNode),
			}
		}

		tmp = tmp.childNodes[b]
		if i == len(word)-1 {
			tmp.isEnd = true
		}
	}
	return nil

}

func (t *trieTree) Chk(line string) bool {
	if len(line) == 0 {
		return false
	}

	tmp := t.root
	begin, pos := 0, 0

	for pos < len(line) {
		b := line[pos]

		if node, ok := tmp.childNodes[b]; !ok {
			begin++
			pos = begin
			tmp = t.root
		} else if node.isEnd {
			return true
		} else {
			pos++
			tmp = node
		}
	}

	return false
}

func (t *trieTree) Filter(line string, replace string) (string, error) {
	if len(line) == 0 {
		return line, errors.New("empty line")
	}

	tmp := t.root
	begin, pos := 0, 0
	result := make([]byte, 0, len(line))

	for pos < len(line) {
		b := line[pos]

		if node, ok := tmp.childNodes[b]; !ok {
			result = append(result, line[begin])
			begin++
			pos = begin
			tmp = t.root
		} else if node.isEnd {
			result = append(result, replace...)
			pos++
			begin = pos
			tmp = t.root
		} else {
			pos++
			tmp = node
		}
	}

	result = append(result, line[begin:]...)
	return string(result), nil
}
