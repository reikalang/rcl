package rcl_test

import (
	"testing"

	"github.com/reikalang/rcl/impl/rcl-go/rcl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_Parse(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		cases := []struct {
			s string
		}{
			{"null"},
			{" null"},
			{"	null	"},
		}
		for _, tc := range cases {
			p := rcl.NewParser(tc.s)
			n, err := p.Parse()
			require.Nil(t, err)
			_, ok := n.(*rcl.Null)
			assert.True(t, ok)
		}
	})

	t.Run("bool", func(t *testing.T) {
		cases := []struct {
			b bool
			s string
		}{
			{true, "true"},
			{true, "	true"},
			{false, "false"},
			{false, " false"},
		}
		for _, tc := range cases {
			p := rcl.NewParser(tc.s)
			n, err := p.Parse()
			require.Nil(t, err)
			assert.Equal(t, tc.b, n.(*rcl.Bool).Val)
		}
	})

	t.Run("string", func(t *testing.T) {
		cases := []struct {
			e string
			s string
		}{
			{"foo", `"foo"`},
			{"foo", `	"foo"`},
			{"foo", `  "foo"`},
			{`foo"`, `"foo\""`},
			{"foo ", `"foo "`},
			{"支持中文", `"支持中文"`},
		}
		for _, tc := range cases {
			p := rcl.NewParser(tc.s)
			n, err := p.Parse()
			require.Nil(t, err)
			assert.Equal(t, tc.e, n.(*rcl.String).Val)
		}
	})

	t.Run("array", func(t *testing.T) {
		cases := []string{
			`[true, null, "bar"]`,
			`[   true, null , "bar"]`,
			`[	true, null, "bar" ]`,
			`[true, null, "bar",]`,
		}
		for _, tc := range cases {
			p := rcl.NewParser(tc)
			n, err := p.Parse()
			require.Nil(t, err)
			arr, ok := n.(*rcl.Array)
			require.True(t, ok)
			values := arr.Values
			require.Equal(t, 3, len(values))
			assert.Equal(t, true, values[0].(*rcl.Bool).Val)
			_, ok = values[1].(*rcl.Null)
			assert.True(t, ok)
			assert.Equal(t, "bar", values[2].(*rcl.String).Val)
		}
	})

	t.Run("object", func(t *testing.T) {
		p := rcl.NewParser(`{"a": "b"}`)
		n, err := p.Parse()
		require.Nil(t, err)
		obj, ok := n.(*rcl.Object)
		require.True(t, ok)
		assert.Equal(t, 1, len(obj.Keys))
		assert.Equal(t, "a", obj.Keys[0].Val)
		assert.Equal(t, "b", obj.Values[0].(*rcl.String).Val)
	})
}
