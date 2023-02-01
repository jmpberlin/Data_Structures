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
	Children [alphabetSize]*Node
	isEnd    bool
}

func (t *Trie) Insert(key string) {
	currentNode := t.Root
	for _, e := range strings.ToLower(key) {
		alphabetIndex := e - 'a'
		if currentNode.Children[alphabetIndex] == nil {
			currentNode.Children[alphabetIndex] = &Node{}
		}
		currentNode = currentNode.Children[alphabetIndex]
	}
	currentNode.isEnd = true
}
func (t *Trie) Lookup(key string) bool {
	currentNode := t.Root
	for _, e := range strings.ToLower(key) {
		alphabetIndex := e - 'a'
		if currentNode.Children[alphabetIndex] == nil {
			return false
		}
		currentNode = currentNode.Children[alphabetIndex]
	}
	return currentNode.isEnd
}

func (t *Trie) Delete(key string) bool {
	lowerCaseKey := strings.ToLower(key)
	return DeleteAndRemoveGarbage(t.Root, nil, lowerCaseKey, 0)
}
func DeleteAndRemoveGarbage(currentNode, previousNode *Node, key string, i int) bool {
	if currentNode == nil {
		return false
	}

	if i == len(key) {
		if !currentNode.isEnd {
			return false
		} else {
			currentNode.isEnd = false
			indexOfRemovableNode := int(key[i-1] - 'a')
			removeNodeIfEmpty(currentNode, previousNode, indexOfRemovableNode)
			return true
		}
	}

	index := key[i] - 'a'
	i++
	deletionSuccessfull := DeleteAndRemoveGarbage(currentNode.Children[index], currentNode, key, i)

	i--

	if previousNode == nil {

		return deletionSuccessfull
	}
	index = key[i-1] - 'a'
	removeNodeIfEmpty(currentNode, previousNode, int(index))
	return deletionSuccessfull
}
func removeNodeIfEmpty(currentNode *Node, previousNode *Node, indexOfRemovableNode int) {
	if currentNode.isEnd {
		return
	}
	for _, e := range currentNode.Children {
		if e != nil {
			return
		}
	}
	if previousNode == nil {
		return
	}
	previousNode.Children[indexOfRemovableNode] = nil

}
