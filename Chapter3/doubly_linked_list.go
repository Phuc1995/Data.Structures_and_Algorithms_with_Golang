package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
	previousNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

//NodeBetweenValues method of LinkedList

func (linkedList *LinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode{
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty{
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	if linkedList.headNode != nil{
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}
	linkedList.headNode = node
}

//LastNode method of LinkedList
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode{
		if node.nextNode == nil{
			lastNode = node
		}
	}
	return lastNode
}

func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}
//NodeWithValue method of LinkedList
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

//AddAfter method of LinkedList
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var nodeWith *Node

	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		//fmt.Println(node.property)
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}

}
//IterateList method of LinkedList
func (linkedList *LinkedList) IterateList() {

	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {

		fmt.Println(node.property)
	}
}


func main()  {

	var linkedList LinkedList

	linkedList = LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
}