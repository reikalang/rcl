package rcl

import (
	"math"
	"sort"
)

// Pos is offset in bytes in current file
type Pos int

type Position struct {
	File   string
	Offset int // offset in bytes, pos from 0
	Line   int // line number, pos from 1
	Column int // column number
}

type File struct {
	name  string
	lines []int // offset of the first character for each line
}

func (f *File) Position(p Pos) Position {
	offset := int(p)
	// TODO: check this logic ...
	i := sort.Search(len(f.lines), func(i int) bool {
		return f.lines[i] > offset
	})
	line := i + 1
	column := offset - f.lines[i] + 1
	return Position{
		File:   f.name,
		Offset: offset,
		Line:   line,
		Column: column,
	}
}

type Node interface {
	Pos() Pos                     // Pos is start position
	End() Pos                     // End is end position. The actual range is [Pos, End), value of range is +1 of last character in token.
	Accept(visitor Visitor) error // Accept should call corresponding visitor method based on node type
}

type baseNode struct {
	pos int
	end int
}

func (b *baseNode) Pos() Pos {
	return Pos(b.pos)
}

func (b *baseNode) End() Pos {
	return Pos(b.end)
}

type Null struct {
	baseNode
}

func (n *Null) Accept(visitor Visitor) error {
	return visitor.VisitNull(n)
}

type Bool struct {
	baseNode
	Val bool
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
	baseNode
	Val  int64
	Type NumberType
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
	baseNode
	Val string
}

func (s *String) Accept(visitor Visitor) error {
	return visitor.VisitString(s)
}

type Array struct {
	baseNode
	Values []Node
}

func (a *Array) Accept(visitor Visitor) error {
	return visitor.VisitArray(a)
}

type Object struct {
	baseNode
	Keys   []*String
	Values []Node
}

func (o *Object) Accept(visitor Visitor) error {
	return visitor.VisitObject(o)
}
