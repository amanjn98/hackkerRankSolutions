package Tree

import (
	"strconv"
	"strings"
)

type TrieNode struct {
	Char     string
	Children [10]*TrieNode
}

type Trie struct {
	RootNode *TrieNode
}

func NewNode(char string) *TrieNode {
	return &TrieNode{Char: char}
}

func NewTrie() *Trie {
	root := NewNode("") // Use an empty string for the root
	return &Trie{RootNode: root}
}

func (t *Trie) Insert(word string) {
	current := t.RootNode
	strippedWords := strings.ReplaceAll(word, " ", "") // Remove spaces
	for _, char := range strippedWords {
		index := char - '0'
		if current.Children[index] == nil {
			current.Children[index] = NewNode(string(char))
		}
		current = current.Children[index]
	}
}

func dfs(sb *string, ans *[]int, root *TrieNode) {
	if root == nil {
		return
	}
	for i := 0; i < 10; i++ {
		if root.Children[i] != nil {
			*sb += root.Children[i].Char
			num, _ := strconv.Atoi(*sb)
			*ans = append(*ans, num)
			dfs(sb, ans, root.Children[i])
			shortenedStr := *sb
			*sb = shortenedStr[:len(shortenedStr)-1]
		}
	}
}

func lexicalOrder(n int) []int {
	t := NewTrie()
	for i := 1; i <= n; i++ {
		key := strconv.Itoa(i)
		t.Insert(key)
	}
	var ans []int
	var sb string
	dfs(&sb, &ans, t.RootNode)
	return ans
}
