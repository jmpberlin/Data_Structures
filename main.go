package main

// HOMEWORK:
// universal Hash function verstehen und einbauen
// include generics=> initiate hashtable for different type combinations ?

// WICHTIG: MAKE DIFFERENT GO ROUTINES WORK ON THE SAME HASH TABLE AND SYNCHRONIZE THE METHODS

// Mutex, Waitgroup, Sync, go Channels to synchronize
// instantiate hashTable function || initialize?

import (
	"fmt"
	trie "hash_table_diy/trie"
)

type key string

type value struct {
	firstName string
	lastName  string
}

type hashTable []bucket

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key      key
	val      value
	nextNode *bucketNode
}

func instantiateHashTable() hashTable {
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
	i := hash(key(s))
	bucket := ht[i]
	bucketNode := bucket.head
	for bucketNode != nil {
		fmt.Println(bucketNode)
		bucketNode = bucketNode.nextNode
	}
}

// Hash function
func hash(k key) int {
	sum := 0
	for _, val := range k {
		sum += int(val)
	}
	return sum % 100
}

// HASHTABLE

// HASHTABLE insert
func (ht hashTable) insert(k key, v value) {
	i := hash(k)
	ht[i].insert(k, v)
}

// HASHTABLE lookup

func (ht hashTable) lookup(input string) (bool, value) {
	k := key(input)
	i := hash(k)
	node := ht[i].lookup(k)
	if node != nil {
		return true, node.val
	}
	return false, value{}
}

// HASHTABLE delete

func (ht hashTable) delete(input string) bool {
	k := key(input)
	i := hash(k)
	return ht[i].delete(k)
}

// BUCKET

// BUCKET insert
func (b *bucket) insert(k key, v value) {
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

func (b *bucket) lookup(k key) *bucketNode {
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

func (b *bucket) delete(k key) bool {
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
func main() {
	// ht := instantiateHashTable()
	// k1 := key("P((")
	// v1 := value{
	// 	firstName: "Pannes",
	// 	lastName:  "Polte",
	// }
	// k2 := key("PP")
	// v2 := value{
	// 	firstName: "Harry",
	// 	lastName:  "Harold",
	// }
	// k3 := key("(P(")
	// v3 := value{
	// 	firstName: "Gustav",
	// 	lastName:  "Gans",
	// }

	// k4 := key("((P")
	// v4 := value{
	// 	firstName: "Kevin",
	// 	lastName:  "Kobold",
	// }
	// k5 := key("Johannes")
	// v5 := value{
	// 	firstName: "Joe",
	// 	lastName:  "von P",
	// }

	// //  All the same hash code: 60
	// ht.insert(k1, v1)

	// // different Hash code
	// ht.insert(k5, v5)

	// if ok, person1 := ht.lookup(string(k5)); ok {
	// 	fmt.Printf("found person %+v under key: %s \n", person1, k1)
	// }
	// ht.printTable()
	// ht.printBucket("((P")
	trie := trie.New()
	trie.Insert("ABBA")
	trie.Insert("ABBA")
	trie.Insert("ABBA")
	trie.Insert("ABBA")
	fmt.Println(trie.Lookup("ABB"))
}

// Write test functions that will input random 10.000 values into the hashtable
// and then retrieve a random set of values again
// gather metrics on execution time, etc.
