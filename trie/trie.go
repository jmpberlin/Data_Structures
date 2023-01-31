package trie

import (
	"strings"
)

const alphabetSize = 26

type Trie struct {
	Root *Node
}

func New() *Trie {
	return &Trie{Root: &Node{}}
}

type Node struct {
	children [alphabetSize]*Node
	isEnd    bool
}

func (t *Trie) Insert(key string) {
	currentNode := t.Root
	for _, e := range strings.ToLower(key) {
		alphabetIndex := e - 'a'
		if currentNode.children[alphabetIndex] == nil {
			currentNode.children[alphabetIndex] = &Node{}
		}
		currentNode = currentNode.children[alphabetIndex]
	}
	currentNode.isEnd = true
}
func (t *Trie) Lookup(key string) bool {
	currentNode := t.Root
	for _, e := range strings.ToLower(key) {
		alphabetIndex := e - 'a'
		if currentNode.children[alphabetIndex] == nil {
			return false
		}
		currentNode = currentNode.children[alphabetIndex]
	}
	return currentNode.isEnd
}

// sanitize the input.
// get rid of everything, that is not an alphabet character
