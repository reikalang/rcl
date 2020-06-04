package rcl

type Node interface {
	Pos() Position
	Accept(visitor Visitor)
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

func (n *Null) Accept(visitor Visitor) {
	visitor.VisitNull(n)
}

type Bool struct {
	Val bool
	Position
}

func (b *Bool) Accept(visitor Visitor) {
	visitor.VisitBool(b)
}

type NumberType byte

const (
	NumberTypeUnknown NumberType = iota
	NumberTypeInt
	NumberTypeDouble
)

type Number struct {
	Val  int64
	Type NumberType
	Position
}

func (n *Number) Accept(visitor Visitor) {
	visitor.VisitNumber(n)
}

type String struct {
	Val string
	Position
}

func (s *String) Accept(visitor Visitor) {
	visitor.VisitString(s)
}

type Array struct {
	Values []Node
	Position
}

func (a *Array) Accept(visitor Visitor) {
	visitor.VisitArray(a)
}

type Object struct {
	Keys   []*String
	Values []Node
	Position
}

func (o *Object) Accept(visitor Visitor) {
	visitor.VisitObject(o)
}
