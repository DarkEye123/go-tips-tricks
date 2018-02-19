package main

import "fmt"

type Node interface {
	GetValue() int
	SetValue(int)
}

type BasicNode struct {
	value int
}

func (n *BasicNode) GetValue() int {
	return n.value
}

func (n *BasicNode) SetValue(val int) {
	n.value = val
}

func NewBasicNode() *BasicNode {
	return new(BasicNode)
}

type PowerNode struct {
	value int
}

func (n *PowerNode) GetValue() int {
	return n.value
}

func (n *PowerNode) SetValue(val int) {
	n.value = val * val
}

func NewPowerNode() *PowerNode {
	return new(PowerNode)
}

func main() {
	var n Node
	n = NewBasicNode()
	n.SetValue(10)
	fmt.Println("Node value is:", n.GetValue())

	n = NewPowerNode()
	n.SetValue(10)
	fmt.Println("Node value is:", n.GetValue())

	if _, ok := n.(*PowerNode); ok {
		fmt.Println("n is of PowerNode type")
	}

}
