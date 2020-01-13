package rcl

import (
	"github.com/dyweb/gommon/errors"
	"reflect"
)

// decode.go traverse AST and decode it into a go struct
// ref: https://github.com/at15/papers-i-read/blob/master/lang/go/go-walkthrough_encoding-json.md

type Visitor interface {
	VisitNull(n *Null)
	VisitBool(b *Bool)
	VisitNumber(n *Number)
	VisitString(s *String)
	VisitArray(arr *Array)
	VisitObject(obj *Object)
}

// TODO: add a printer as visitor implementation, it serves as formatter

type Decoder struct {
	rv reflect.Value
	n  Node
}

func (d *Decoder) VisitNull(n *Null) {
	panic("implement me")
}

func (d *Decoder) VisitBool(b *Bool) {
	panic("implement me")
}

func (d *Decoder) VisitNumber(n *Number) {
	panic("implement me")
}

func (d *Decoder) VisitString(s *String) {
	panic("implement me")
}

func (d *Decoder) VisitArray(arr *Array) {
	panic("implement me")
}

func (d *Decoder) VisitObject(obj *Object) {
	panic("implement me")
}

func (d *Decoder) Error() error {
	return nil
}

func UnmarshalRCL(b []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("must decode to a nil pointer")
	}
	n, err := NewParser(string(b)).Parse()
	if err != nil {
		return errors.Wrap(err, "error parse before decode")
	}
	dec := &Decoder{
		rv: rv,
	}
	n.Accept(dec)
	return dec.Error()
}
