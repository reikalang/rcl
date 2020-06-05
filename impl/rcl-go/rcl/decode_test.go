package rcl_test

import (
	"encoding/json"
	"testing"

	"github.com/dyweb/gommon/util/testutil"
	"github.com/reikalang/rcl/impl/rcl-go/rcl"
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

type container struct {
	Name        string
	Tag         string
	HostNetwork bool
	Port        int
	Envs        []string
}

func TestDecoder_VisitObject(t *testing.T) {
	s := `
{
	"Name": "nginx",
	"Tag": "1.6",
	"HostNetwork": true,
	"Port": 10086,
	"Envs": [
		"DEBUG=1",
		"GC=off",
	]
}
`
	var c container
	err := rcl.UnmarshalRCL([]byte(s), &c)
	if err != nil {
		t.Errorf("%s", err)
	}
	testutil.DumpAsJson(t, c)
}
