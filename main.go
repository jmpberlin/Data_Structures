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
	fmt.Println(ht.Delete("P(("))

	trie := triePkg.New()
	trie.Insert("ABBA")
	fmt.Println("delete 'ABBA':\t", trie.Delete("ABBA"))
	fmt.Println("lookup ABBA", trie.Lookup("ABBA"))

}

// Write test functions that will input random 10.000 values into the hashtable
// and then retrieve a random set of values again
// gather metrics on execution time, etc.
