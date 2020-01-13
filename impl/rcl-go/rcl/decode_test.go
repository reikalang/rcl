package rcl_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// decode_test.go tests decoding into struct

type TestUser struct {
	Name string
	Age  int
}

// just used for jumping into JSON decoding code
func TestJSON(t *testing.T) {
	s := `{"name": "xiao ming", "age": 10}`
	var u TestUser
	err := json.Unmarshal([]byte(s), &u)
	assert.Nil(t, err)
}
