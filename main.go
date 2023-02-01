package main

// HOMEWORK:
// universal Hash function verstehen und einbauen
// include generics=> initiate hashtable for different type combinations ?

// WICHTIG: MAKE DIFFERENT GO ROUTINES WORK ON THE SAME HASH TABLE AND SYNCHRONIZE THE METHODS

// Mutex, Waitgroup, Sync, go Channels to synchronize
// instantiate hashTable function || initialize?

// idiom from example:

// waitgroup
// routines
// loop over the listening function
// one routine closes the channel
// ends the function

import (
	hashtablePkg "Data_Structures/hashtable"
	triePkg "Data_Structures/trie"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ht := hashtablePkg.InstantiateHashTable()
	k1 := hashtablePkg.Key("P((")
	v1 := hashtablePkg.Value{
		FirstName: "Pannes",
		LastName:  "Polte",
	}

	ht.Insert(k1, v1)

	if ok, person1 := ht.Lookup(string(k1)); ok {
		fmt.Printf("found person %+v under key: %s \n", person1, k1)
	}
	fmt.Println("delete P((:", ht.Delete("P(("))

	trie := triePkg.New()
	trie.Insert("ABBA")
	trie.Insert("ABBAESQUE")
	trie.Insert("ABBAESQUELALO")
	fmt.Println("delete 'ABBA':\t", trie.Delete("ABBA"))
	fmt.Println("lookup ABBA", trie.Lookup("ABBA"))
	fmt.Println("lookup ABBAESQUE", trie.Lookup("ABBAESQUE"))
	fmt.Println("lookup ABBAESQUELALO", trie.Lookup("ABBAESQUELALO"))
	fmt.Println("delete 'ABBAESQUE':\t", trie.Delete("ABBAESQUE"))
	fmt.Println("delete 'ABBAESQUELALO':\t", trie.Delete("ABBAESQUELALO"))
	fmt.Println("lookup ABBAESQUE", trie.Lookup("ABBAESQUE"))
	fmt.Println("lookup ABBAESQUELALO", trie.Lookup("ABBAESQUELALO"))

	c := make(chan bool)
	counter := 0
	var wg sync.WaitGroup
	wg.Add(200)
	go func() {
		// fmt.Println("I will wait! ")
		// defer fmt.Println("I'm done waiting! ")
		wg.Wait()
		defer close(c)
	}()
	go hashTableInsertBot(&ht, 100, c, &wg)
	go hashTableInsertBot(&ht, 100, c, &wg)
	for success := range c {
		if success {
			counter++
		}
	}
	fmt.Println("this is the counter: ", counter)
	fmt.Println("this is the hashtable: ", ht)
}

func hashTableInsertBot(ht *hashtablePkg.HashTable, elementsToBeInserted int, c chan bool, wg *sync.WaitGroup) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < elementsToBeInserted; i++ {
		randomWord1 := ""
		randomWord2 := ""
		randomWord3 := ""
		for j := 0; j < 5; j++ {
			randomWord1 = randomWord1 + string(r1.Intn(26)+'a')
			randomWord2 = randomWord2 + string(r1.Intn(26)+'a')
			randomWord3 = randomWord3 + string(r1.Intn(26)+'a')
		}
		val := hashtablePkg.Value{FirstName: randomWord1, LastName: randomWord1}
		key := hashtablePkg.Key(randomWord3)
		ht.Insert(key, val)
		c <- true
		wg.Done()
	}
}

// Write test functions that will input random 10.000 values into the hashtable
// and then retrieve a random set of values again
// gather metrics on execution time, etc.
