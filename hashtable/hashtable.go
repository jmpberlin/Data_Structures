package hashtable

import "fmt"

type Key string

type Value struct {
	FirstName string
	LastName  string
}

type hashTable []bucket

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key      Key
	val      Value
	nextNode *bucketNode
}

func InstantiateHashTable() hashTable {
	ht := hashTable{}
	for i := 0; i < 100; i++ {
		ht = append(ht, bucket{})
	}
	return ht
}

// Print the HashTable somewhat graphically
func (h hashTable) printTable() {
	for index, val := range h {
		fmt.Println(index, ": \t", val)

	}
}

// Print a specific Bucket from a hash
func (ht hashTable) printBucket(s string) {
	i := hash(Key(s))
	bucket := ht[i]
	bucketNode := bucket.head
	for bucketNode != nil {
		fmt.Println(bucketNode)
		bucketNode = bucketNode.nextNode
	}
}

// Hash function
func hash(k Key) int {
	sum := 0
	for _, val := range k {
		sum += int(val)
	}
	return sum % 100
}

// HASHTABLE

// HASHTABLE insert
func (ht hashTable) Insert(k Key, v Value) {
	i := hash(k)
	ht[i].insert(k, v)
}

// HASHTABLE lookup

func (ht hashTable) Lookup(input string) (bool, Value) {
	k := Key(input)
	i := hash(k)
	node := ht[i].lookup(k)
	if node != nil {
		return true, node.val
	}
	return false, Value{}
}

// HASHTABLE delete

func (ht hashTable) Delete(input string) bool {
	k := Key(input)
	i := hash(k)
	return ht[i].delete(k)
}

// BUCKET

// BUCKET insert
func (b *bucket) insert(k Key, v Value) {
	nodeToUpdate := b.lookup(k)

	if len(nodeToUpdate.key) == 0 {
		currentFirstNode := b.head
		nodeToBeInserted := bucketNode{
			key:      k,
			val:      v,
			nextNode: currentFirstNode,
		}
		b.head = &nodeToBeInserted
	} else {
		nodeToUpdate.val = v
	}
}

// BUCKET lookup

func (b *bucket) lookup(k Key) *bucketNode {
	keyNotFound := &bucketNode{}
	if b.head == nil {
		return keyNotFound
	}
	currentNode := *b.head
	for {
		if currentNode.key == k {
			return &currentNode
		} else if currentNode.nextNode == nil {
			return keyNotFound
		} else {
			currentNode = *currentNode.nextNode
		}
	}
}

// BUCKET delete

func (b *bucket) delete(k Key) bool {
	noDeleteableKeyFound := false
	keyDeleted := true
	if b.head == nil {
		return noDeleteableKeyFound
	}
	if b.head.key == k {
		b.head = b.head.nextNode
		return keyDeleted
	}
	if b.head.nextNode == nil {
		return noDeleteableKeyFound
	}
	currentNode := b.head
	nextNode := b.head.nextNode
	for {
		if nextNode.key == k {
			currentNode.nextNode = nextNode.nextNode
			return keyDeleted
		}
		if nextNode.nextNode == nil {
			return noDeleteableKeyFound
		}
		currentNode = nextNode
		nextNode = nextNode.nextNode
	}
}
