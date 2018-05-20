# HCL

- https://github.com/hashicorp/hcl
- https://github.com/hashicorp/hil  a lightweight embedded language used primarily for configuration interpolation.
- https://github.com/hashicorp/hcl2
  - HCL + HIL

## HCL 1

- https://github.com/hashicorp/hcl

````go
// token/position.go
// Pos describes an arbitrary source position， including the file, line, and column location.
// A Position is valid if the line number is > 0.
type Pos struct {
	Filename string // filename, if any
	Offset   int    // offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (character count)
}

// ast/ast.go
// Node is an element in the abstract syntax tree.
type Node interface {
	node()
	Pos() token.Pos
}

type File struct {
	Node     Node            // usually a *ObjectList
	Comments []*CommentGroup // list of all comments in the source
}

// An HCL file itself is an ObjectList.
type ObjectList struct {
	Items []*ObjectItem
}

// Comment node represents a single //, # style or /*- style commment
type Comment struct {
	Start token.Pos // position of / or #
	Text  string
}

func (c *Comment) Pos() token.Pos {
	return c.Start
}
````

## HIL

- https://github.com/hashicorp/hil
- it's a interpreter, it support functions
- variables are defined externally (like context for template)

````go
// ast/ast.go
type Node interface {
	// Accept is called to dispatch to the visitors. It must return the
	// resulting Node (which might be different in an AST transform).
	Accept(Visitor) Node
	// Pos returns the position of this node in some source.
	Pos() Pos
	// Type returns the type of this node for the given context.
	Type(Scope) (Type, error)
}

// Pos is the starting position of an AST node
type Pos struct {
	Column, Line int    // Column/Line number, starting at 1
	Filename     string // Optional source filename, if known
}

// Note that this isn't a true implementation of the visitor pattern, which
// generally requires proper type dispatch on the function. However,
// implementing this basic visitor pattern style is still very useful even
// if you have to type switch.
type Visitor func(Node) Node

// ast/scope.go
// Scope is the interface used to look up variables and functions while
// evaluating. How these functions/variables are defined are up to the caller.
type Scope interface {
	LookupFunc(string) (Function, bool)
	LookupVar(string) (Variable, bool)
}

// ast/stack.go
// Stack is a stack of Node.
type Stack struct {
	stack []Node
}

// eval.go
type evalVisitor struct {
	Scope ast.Scope
	Stack ast.Stack

	err  error
	lock sync.Mutex
}
````

## HCL2

much larger codebase than hcl + hil, and still under heavy construction

- `gohcl` unmarshal into go struct using reflection
- `hcl/hclsyntax` contains parser, ast etc.
  - the node and visitor are similar to hil instead of hcl
  - ast nodes are defined in `expression.go`, `expression_*.go`, `structure.go`
  - `generate.go` calls all the generators
  
````go
// hcl/hclsyntax/node.go

// This is a closed interface, so it cannot be implemented from outside of
// this package.
type Node interface {
	// This is the mechanism by which the public-facing walk functions
	// are implemented. Implementations should call the given function
	// for each child node and then replace that node with its return value.
	// The return value might just be the same node, for non-transforming
	// walks.
	walkChildNodes(w internalWalkFunc)
	Range() hcl.Range
}

type internalWalkFunc func(Node) Node

// hcl/pos.go
// Pos represents a single position in a source file, by addressing the
// start byte of a unicode character encoded in UTF-8.
//
// Pos is generally used only in the context of a Range, which then defines
// which source file the position is within.
type Pos struct {
	// Line is the source code line where this position points. Lines are
	// counted starting at 1 and incremented for each newline character
	// encountered.
	Line int

	// Column is the source code column where this position points, in
	// unicode characters, with counting starting at 1.
	//
	// Column counts characters as they appear visually, so for example a
	// latin letter with a combining diacritic mark counts as one character.
	// This is intended for rendering visual markers against source code in
	// contexts where these diacritics would be rendered in a single character
	// cell. Technically speaking, Column is counting grapheme clusters as
	// used in unicode normalization.
	Column int

	// Byte is the byte offset into the file where the indicated character
	// begins. This is a zero-based offset to the first byte of the first
	// UTF-8 codepoint sequence in the character, and thus gives a position
	// that can be resolved _without_ awareness of Unicode characters.
	Byte int
}

// Range represents a span of characters between two positions in a source
// file.
//
// This struct is usually used by value in types that represent AST nodes,
// but by pointer in types that refer to the positions of other objects,
// such as in diagnostics.
type Range struct {
	Filename string
	Start, End Pos
}
````
