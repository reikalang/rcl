package rcl

import "math"

type Node interface {
	Pos() Position
	Accept(visitor Visitor) error
}

type Position struct {
	Line int
	Col  int
}

func (p *Position) Pos() Position {
	return *p
}

type Null struct {
	Position
}

func (n *Null) Accept(visitor Visitor) error {
	return visitor.VisitNull(n)
}

type Bool struct {
	Val bool
	Position
}

func (b *Bool) Accept(visitor Visitor) error {
	return visitor.VisitBool(b)
}

type NumberType byte

const (
	NumberTypeUnknown NumberType = iota
	NumberTypeInt
	NumberTypeDouble
)

func (t NumberType) String() string {
	return []string{
		"unknown",
		"int",
		"double",
	}[t]
}

type Number struct {
	Val  int64
	Type NumberType
	Position
}

func (n *Number) Accept(visitor Visitor) error {
	return visitor.VisitNumber(n)
}

func (n *Number) Int() int64 {
	if n.Type != NumberTypeInt {
		panic("number type is not int but " + n.Type.String())
	}
	return n.Val
}

func (n *Number) Double() float64 {
	if n.Type != NumberTypeDouble {
		panic("number type is not double but " + n.Type.String())
	}
	return math.Float64frombits(uint64(n.Val))
}

type String struct {
	Val string
	Position
}

func (s *String) Accept(visitor Visitor) error {
	return visitor.VisitString(s)
}

type Array struct {
	Values []Node
	Position
}

func (a *Array) Accept(visitor Visitor) error {
	return visitor.VisitArray(a)
}

type Object struct {
	Keys   []*String
	Values []Node
	Position
}

func (o *Object) Accept(visitor Visitor) error {
	return visitor.VisitObject(o)
}
