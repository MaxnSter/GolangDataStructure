package trietree

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewTrieTree(t *testing.T) {
	tree := NewTrieTree()
	if tree == nil {
		t.Fail()
	}
}

func TestBuildTrieTree(t *testing.T) {
	input := []byte("ab edf 贪玩蓝月")
	rd := bytes.NewReader(input)
	tree := BuildTrieTree(rd)
	if tree == nil {
		t.Fail()
	}
}

func TestTrieTree_Chk(t *testing.T) {
	input := []byte("ab edf 贪玩蓝月")
	rd := bytes.NewReader(input)
	tree := BuildTrieTree(rd)

	lines := []string{
		"我代言了贪玩蓝月",
		"abcd",
		"edf",
		"acccccccab",
	}

	for _, line := range lines {
		if !tree.Chk(line) {
			t.Log("chk error at %s", line)
			t.Fail()
		}
	}
}

func TestTrieTree_Chk2(t *testing.T) {

	input := []byte("ab edf 贪玩蓝月")
	rd := bytes.NewReader(input)
	tree := BuildTrieTree(rd)

	lines := []string{
		"我代言了贪玩",
		"helllllo",
		"deef",
		"acb",
	}

	for _, line := range lines {
		if tree.Chk(line) {
			t.Log("chk error at ", line)
			t.Fail()
		}
	}
}

func TestTrieTree_Filter(t *testing.T) {

	input := []byte("ab edf 贪玩蓝月")
	rd := bytes.NewReader(input)
	tree := BuildTrieTree(rd)

	lines := []string{
		"我代言了贪玩蓝月",
		"abcd",
		"def",
		"acccccccab",
	}

	correct := []string{
		"我代言了*",
		"*cd",
		"def",
		"accccccc*",
	}

	for i, line := range lines {
		result, _ := tree.Filter(line, "*")
		if strings.Compare(result, correct[i]) != 0 {
			t.Logf("filter error at %s, should be %s", result, correct[i])
			t.Fail()
		}
	}
}
