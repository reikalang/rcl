package foo

// Comment for struct
type Foo struct {
	// comment before
	a string // comment at same line
	// comment after
	b []string
	// comment after without fields following
}

func newFoo() Foo {
	return Foo{
		// comment before
		a: "a", // comment at same line
		// comment after
		b: []string{
			// comment before
			"a", // comment at same line
			// comment after
		},
	}
}
