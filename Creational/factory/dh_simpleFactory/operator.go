package main

import (
	"errors"
	"fmt"
)

type Operator interface {
	GetResult() float32
}

const (
	Add = iota
	Sub
	Mul
	Div
)

func GetOperator(m, num1, num2 int) (Operator, error) {
	switch m {
	case Add:
		return &AddOp{Num1: num1, Num2: num2}, nil
	case Sub:
		return &SubOp{Num1: num1, Num2: num2}, nil
	case Mul:
		return &MulOp{Num1: num1, Num2: num2}, nil
	case Div:
		return &DivOp{Num1: num1, Num2: num2}, nil
	default:
		return nil, errors.New("invalid operator")
	}
}

type AddOp struct {
	Num1, Num2 int
}

func (op *AddOp) GetResult() float32 {
	return float32(op.Num1 + op.Num2)
}

type SubOp struct {
	Num1, Num2 int
}

func (op *SubOp) GetResult() float32 {
	return float32(op.Num1 - op.Num2)
}

type MulOp struct {
	Num1, Num2 int
}

func (op *MulOp) GetResult() float32 {
	return float32(op.Num1 * op.Num2)
}

type DivOp struct {
	Num1, Num2 int
}

func (op *DivOp) GetResult() float32 {
	if op.Num2 != 0 {
		return float32(op.Num1 / op.Num2)
	} else {
		return 0.0
	}
}

func main() {
	Num1 := 10
	Num2 := 5

	// 创建加法运算符对象
	addOp, err := GetOperator(Add, Num1, Num2)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := addOp.GetResult()
	fmt.Println("加法运算结果:", result)

	// 创建减法运算符对象
	subOp, err := GetOperator(Sub, Num1, Num2)
	if err != nil {
		fmt.Println(err)
		return
	}
	result = subOp.GetResult()
	fmt.Println("减法运算结果:", result)

	// 创建乘法运算符对象
	mulOp, err := GetOperator(Mul, Num1, Num2)
	if err != nil {
		fmt.Println(err)
		return
	}
	result = mulOp.GetResult()
	fmt.Println("乘法运算结果:", result)

	// 创建除法运算符对象
	divOp, err := GetOperator(Div, Num1, Num2)
	if err != nil {
		fmt.Println(err)
		return
	}
	result = divOp.GetResult()
	fmt.Println("除法运算结果:", result)
}
