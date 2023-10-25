package main

import "fmt"

// Element接口表示数据结构中的元素
type Element interface {
	Accept(Visitor)
}

// ConcreteElementA是Element的具体实现
type ConcreteElementA struct {
	Name string
}

func (e *ConcreteElementA) Accept(v Visitor) {
	v.VisitConcreteElementA(e)
}

// ConcreteElementB是Element的具体实现
type ConcreteElementB struct {
	Age int
}

func (e *ConcreteElementB) Accept(v Visitor) {
	v.VisitConcreteElementB(e)
}

type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

// ConcreteVisitor是Visitor接口的具体实现
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Printf("ConcreteVisitor visits ConcreteElementA with name: %s\n", e.Name)
}

func (v *ConcreteVisitor) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Printf("ConcreteVisitor visits ConcreteElementB with age: %d\n", e.Age)
}

type ObjectStructure struct {
	elements []Element
}

func (o *ObjectStructure) Attach(e Element) {
	o.elements = append(o.elements, e)
}

func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range o.elements {
		element.Accept(visitor)
	}
}

func main() {
	objectStructure := &ObjectStructure{}
	visitor := &ConcreteVisitor{}

	objectStructure.Attach(&ConcreteElementA{Name: "Amy"})
	objectStructure.Attach(&ConcreteElementB{Age: 18})

	objectStructure.Accept(visitor)
}
