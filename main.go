package main

// HOMEWORK:
// universal Hash function verstehen und einbauen
// include generics=> initiate hashtable for different type combinations ?

// WICHTIG: MAKE DIFFERENT GO ROUTINES WORK ON THE SAME HASH TABLE AND SYNCHRONIZE THE METHODS

// Mutex, Waitgroup, Sync, go Channels to synchronize
// instantiate hashTable function || initialize?
// QUESTION: Linked-List -> Am I doing this correctly?
// how does removing from a linked list work?
// slicing something out of an array seems costly and not to be the correct way

import (
	"fmt"
)

type key string

type value struct {
	firstName string
	lastName  string
}

type hashTable []bucket

// MAKE BUCKET A STRUCT WITH ONLY A POINTER TO THE FIRST BUCKET NODE
// WHEN ADDING A NODE WE JUST CHANGE THE HEAD POINTER AND PUT THE OLD ONE IN THE NEW "NEXT"_VALUE
//
type bucket []*bucketNode

type bucketNode struct {
	key      key
	val      value
	end      bool
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
func (h hashTable) print() {
	for index, val := range h {
		fmt.Println(index, ": \t", val)
		for i, e := range val {
			fmt.Print(i, ":\t ")
			fmt.Print(*e, "\n ")
		}
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

func (ht hashTable) lookup(k key) (bool, value) {
	i := hash(k)
	bn := ht[i].lookup(k)
	if bn != nil {
		return true, bn.val
	}
	return false, value{}
}

// HASHTABLE delete

func (ht hashTable) delete(k key) {
	i := hash(k)
	ht[i].delete(k)
}

// BUCKET

// BUCKET insert
func (b *bucket) insert(k key, v value) {
	// check if the key already exists in the bucketNode
	bn := b.lookup(k)
	// // if it exists update the node with new value
	if bn != nil {
		bn.val = v
	} else {
		newNode := bucketNode{
			key: k,
			val: v,
		}
		// prepend the node to the beginning of the bucket if the bucket is not empty
		if len(*b) != 0 {
			bucketValue := *b
			newNode.nextNode = bucketValue[0]
			*b = append(bucket{&newNode}, *b...)
		} else {
			// append the node to the bucket if it will be the first node
			newNode.end = true
			newNode.nextNode = nil
			*b = append(*b, &newNode)
		}
	}
}

// BUCKET lookup

func (b *bucket) lookup(k key) *bucketNode {
	for _, node := range *b {
		if node.key == k {
			return node
		}
	}
	return nil
}

// BUCKET delete
func (b *bucket) delete(k key) {
	// TO-DO: IMPORTANT!

}
func main() {
	ht := instantiateHashTable()
	k1 := key("P((")
	v1 := value{
		firstName: "P((annes",
		lastName:  "Polte",
	}
	k2 := key("PP")
	v2 := value{
		firstName: "PPPatrick",
		lastName:  "Star",
	}
	k3 := key("(P(")
	v3 := value{
		firstName: "(P(",
		lastName:  "another link",
	}

	k4 := key("((P")
	v4 := value{
		firstName: "((P",
		lastName:  "yet another link",
	}

	ht.insert(k1, v1)
	ht.insert(k2, v2)
	ht.insert(k3, v3)
	ht.insert(k4, v4)
	if ok, person1 := ht.lookup("PP"); ok {
		fmt.Println("found person \t\t1 \t", person1)
	}
	if ok, person2 := ht.lookup("((P"); ok {
		fmt.Println("found person \t\t2: \t", person2)
	}
	if ok, person3 := ht.lookup("HANNES"); ok {
		fmt.Println("found person \t\t3: \t", person3)
	} else {
		fmt.Println("did not find person \t3: \t", person3)
	}

	// v1 = value{
	// 	firstName: "TEST TEST TEST",
	// 	lastName:  "POLTE",
	// }

	// ht.insert(k1, v1)
	// if ok, person1 := ht.lookup(k1); ok {
	// 	fmt.Println("found person \t\t1 \t", person1)
	// }
	ht.print()

}
