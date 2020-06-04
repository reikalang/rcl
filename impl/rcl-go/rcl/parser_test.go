package rcl_test

import (
	"strconv"
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
			n := mustParse(t, tc.s)
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
			n := mustParse(t, tc.s)
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
			n := mustParse(t, tc.s)
			assert.Equal(t, tc.e, n.(*rcl.String).Val)
		}
	})

	t.Run("int", func(t *testing.T) {
		cases := []struct {
			n int64
			s string
		}{
			{0, "0"},
			{-1, "-1"},
			{100, "100"},
			{10086, "10086"},
			{100_000, "100_000"},
		}
		for _, tc := range cases {
			n := mustParse(t, tc.s)
			assert.Equal(t, tc.n, n.(*rcl.Number).Val)
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
			n := mustParse(t, tc)
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
		n := mustParse(t, `{"a": "b"}`)
		obj, ok := n.(*rcl.Object)
		require.True(t, ok)
		assert.Equal(t, 1, len(obj.Keys))
		assert.Equal(t, "a", obj.Keys[0].Val)
		assert.Equal(t, "b", obj.Values[0].(*rcl.String).Val)
	})
}

// ParseInt does not support _, but we want that in RCL.
func TestStd_ParseInt(t *testing.T) {
	i := "100_100"
	_, err := strconv.ParseInt(i, 10, 64)
	require.NotNil(t, err)
}

func mustParse(t *testing.T, s string) rcl.Node {
	p := rcl.NewParser(s)
	n, err := p.Parse()
	require.Nil(t, err, s)
	return n
}
