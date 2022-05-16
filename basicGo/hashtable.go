package main

import "fmt"

//Hash Table --> Insert,Search, Delete

//Bucket structure--> Insert,Search,Delete
const arsize = 10

type hashtable struct {
	arr [arsize]*bucketStru
}
type bucketStru struct {
	head *bucketNode
}
type bucketNode struct {
	key  string
	next *bucketNode
}

//Init function will create a bucket in each slot

func Init() *hashtable {
	res := &hashtable{}
	for i := range res.arr {
		res.arr[i] = &bucketStru{}
	}
	return res
}

//Insert the key into hashTable
func (h *hashtable) Insert(key string) {
	index := hash(key)
	h.arr[index].insert(key)
}

//Seraching the from the hashTable
func (h *hashtable) Search(key string) bool {
	index := hash(key)
	return h.arr[index].search(key)
}

//Delete the key from hashTable
func (h *hashtable) delete(key string) {
	index := hash(key)
	h.arr[index].delete(key)
}

//Insert
func (b *bucketStru) insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

//Search
func (b *bucketStru) search(k string) bool {
	currNode := b.head
	for currNode != nil {
		if currNode.key == k {
			return true
		}
		currNode = currNode.next
	}
	return false
}

//Delete

func (b *bucketStru) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

func hash(key string) int {
	sum := 0
	for _, k := range key {
		sum += int(k)
	}
	return sum % arsize
}

func main() {
	hashTable := Init()
	list := []string{
		"RANDY",
		"ERIC",
		"JHON",
		"MIT",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}

	Bucket := &bucketStru{}
	Bucket.insert("RANDY")
	fmt.Println(Bucket.search("RANDY"))
	fmt.Println(Bucket.search("MIT"))
}
