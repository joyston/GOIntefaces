package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type NilInterface interface {
	lUndeclaredMethod()
}

func main() {
	var a Abser //Interface
	b := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4} //struct
	a = b
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())

	var i I = Dummy{"Test Interface"}
	i.M()
	describe(i)

	var j I

	j = MyFloat(math.Pi)
	j.M()
	describe(j)

	var lI interface{} //empty interface
	describeEmptyInterface(lI)
	//describe(lI)

	lI = 7
	describeEmptyInterface(lI)

	lI = "String interface value"
	describeEmptyInterface(lI)

	var lIntf NilInterface
	describe(lIntf)
}

func describeEmptyInterface(i interface{}) {
	fmt.Printf("Value: %v and Type: %T\n", i, i)
}

type Dummy struct {
	s string
}

func (x Dummy) M() {
	fmt.Println(x.s)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) M() {
	fmt.Println(f)
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func describe(i I) {
	fmt.Printf("Value: %v & Type: %T\n", i, i)
}
