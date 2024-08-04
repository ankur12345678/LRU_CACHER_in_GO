package main

import (
	"bufio"
	"fmt"
	"os"
)

// for storing the node pointers for efficient access
var hash = make(map[int]*Node)

// Node structure of a Doubly Linked list
type Node struct {
	data string
	next *Node
	prev *Node
}

// Doubly Linked List structure that contains the head which points to the starting node
type DList struct {
	head *Node
	back *Node
}

// initializing the list
var list DList

// initializing the length count
var dllMaxLength int

// current list length
var listLength int = 0

func main() {
	//taaking user input for length
	fmt.Print("Enter the Length of the LRU cache: ")
	fmt.Scan(&dllMaxLength)

	//input variable for PUT and GET functionality
	//PUT : P
	//GET : G
	//View cache : V
	//Quit : Q
	//initially an arbitrary value "H"
	var inputChar string = "H"

	for 1 == 1 {
		//Giving user the options for PUT and GET
		fmt.Println("--------XXXXX--------")
		fmt.Print("For GET, type G\n")
		fmt.Print("For PUT, type P\n")
		fmt.Print("For exit, type Q\n")
		fmt.Print("For view Cache, type V\n")
		fmt.Printf("Current space used: %v\n", listLength)
		fmt.Scan(&inputChar)

		if inputChar == "Q" {
			break
		} else if inputChar == "P" {
			var key int
			var value string

			fmt.Println("Enter the KEY to store: ")
			fmt.Scan(&key)
			fmt.Println("Enter the VALUE associated to the KEY entered above: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			err := scanner.Err()
			if err != nil {
				fmt.Print("Error in reading the value\n")
			}
			value = scanner.Text()

			lruPut(key, value)
		} else if inputChar == "G" {
			var key int
			fmt.Print("Enter the KEY : ")
			fmt.Scan(&key)

			lruGet(key)
		} else if inputChar == "V" {
			lruView()
		} else {
			fmt.Print("Please enter some valid input !!\n")
		}

	}

}

func lruView() {
	temp := list.head
	fmt.Print("list: ")
	for temp != nil {

		fmt.Printf("%v <--> ", temp.data)

		temp = temp.next
	}
	fmt.Print("nil \n")
}

func lruGet(key int) {
	//Takes input a KEY searches the hashmap
	// if KEY present in hashmap then returns the value associated to the KEY
	// if KEY not present then return -1

	if hash[key] == nil {
		fmt.Print("-1\n")
		return
	}

	//get associated node to that value
	var nodeOfKey *Node = hash[key]
	fmt.Println(nodeOfKey.data)

	fmt.Print("-------")
	fmt.Println(nodeOfKey.prev)
	fmt.Println(nodeOfKey.next)
	fmt.Print("-------")

	lruUpadte(nodeOfKey)
}

func lruUpadte(nodeOfKey *Node) {
	//update the priority of node (bring it to the head to DLL)
	temp := list.head

	if nodeOfKey.prev != nil && nodeOfKey.next != nil {

		// it is in somewhere mid
		prevNode := nodeOfKey.prev
		nextNode := nodeOfKey.next

		prevNode.next = nextNode
		nextNode.prev = prevNode

		nodeOfKey.prev = nil
		nodeOfKey.next = temp

		list.head = nodeOfKey

		return
	}

	if nodeOfKey.next == nil && nodeOfKey.prev != nil {

		//it is the last node
		prevNode := nodeOfKey.prev
		prevNode.next = nil

		nodeOfKey.prev = nil
		nodeOfKey.next = temp
		list.head.prev = nodeOfKey
		list.head = nodeOfKey

		//updaing back pointer
		list.back = prevNode
		return
	}

	return
}

func lruDetachLastNode(node *Node) {
	//remove its link to its prev
	prevNode := node.prev
	prevNode.next = nil

	node.prev = nil

	//remove it from hash map
	for k, v := range hash {
		if v == node {
			delete(hash, k)
			break
		}
	}
	listLength -= 1
}

func lruAddNode(key int, value string) {
	//create a node with given values and add
	var newNode = new(Node)
	newNode.prev = nil
	newNode.next = nil
	newNode.data = value

	if listLength == 0 {
		list.head = newNode
	} else {
		list.head.prev = newNode
		newNode.next = list.head
		list.head = newNode
	}

	listLength += 1

	//also update the hashmap
	hash[key] = newNode

}

func lruPut(key int, value string) {
	//check if the key exists or not
	// if key exists then update the value of the node with the current value entered

	if hash[key] != nil {
		//find the associated node with the key
		nodeOfKey := hash[key]
		nodeOfKey.data = value

		lruUpadte(nodeOfKey)
		return
	}

	//if node doesnt exist then
	//check if chache has space or not
	if listLength == dllMaxLength {
		//delete the last node and add the node at the beginning
		//the length of list wont change
		lruDetachLastNode(list.back)
		lruAddNode(key, value)
		return
	} else {
		//create a node and add at beginning(head)
		lruAddNode(key, value)
		return
	}

}
