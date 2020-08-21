package rcl

import (
	"reflect"

	"github.com/dyweb/gommon/errors"
)

// decode.go traverse AST and decode it into a go struct
// ref: https://github.com/at15/papers-i-read/blob/master/lang/go/go-walkthrough_encoding-json.md

// TODO: maybe return (Node, error)? but we are already giving pointer, one can *ptr = newVal
type Visitor interface {
	VisitNull(n *Null) error
	VisitBool(b *Bool) error
	VisitNumber(n *Number) error
	VisitString(s *String) error
	VisitArray(arr *Array) error
	VisitObject(obj *Object) error
}

// TODO: add a printer as visitor implementation, it serves as formatter

type Decoder struct {
	rv reflect.Value
	n  Node
}

func (d *Decoder) VisitNull(n *Null) error {
	return errors.New("visit null not implemented")
}

func (d *Decoder) VisitBool(b *Bool) error {
	if d.rv.Kind() != reflect.Bool {
		// TODO: typed error
		return errors.Errorf("invalid type for string got %s", d.rv.Kind())
	}
	d.rv.SetBool(b.Val)
	return nil
}

func (d *Decoder) VisitNumber(n *Number) error {
	switch n.Type {
	case NumberTypeUnknown:
		return errors.New("number is of type unknown")
	case NumberTypeInt:
		switch d.rv.Kind() {
		// TODO: add more int types ...
		case reflect.Int, reflect.Int64:
			d.rv.SetInt(n.Int())
		case reflect.Uint, reflect.Uint64:
			if n.Val < 0 {
				return errors.Errorf("can't decode negative value into unsigned int got %d", n.Val)
			}
			d.rv.SetUint(uint64(n.Int()))
		default:
			return errors.Errorf("unexpected type for int, got %s", d.rv.Kind())
		}
	case NumberTypeDouble:
		switch d.rv.Kind() {
		case reflect.Float32, reflect.Float64:
			d.rv.SetFloat(n.Double())
		default:
			return errors.Errorf("unexpected type for double, got %s", d.rv.Kind())
		}
	default:
		return errors.New("parser impl error: unhandled number type")
	}
	return nil
}

func (d *Decoder) VisitString(s *String) error {
	if d.rv.Kind() != reflect.String {
		// TODO: typed error
		return errors.Errorf("invalid type for string got %s", d.rv.Kind())
	}
	d.rv.SetString(s.Val)
	return nil
}

func (d *Decoder) VisitArray(arr *Array) error {
	if d.rv.Kind() != reflect.Slice {
		// TODO: typed error
		return errors.Errorf("invalid type for array got %s", d.rv.Kind())
	}
	l := len(arr.Values)
	newv := reflect.MakeSlice(d.rv.Type(), l, l)
	d.rv.Set(newv)
	for i := 0; i < l; i++ {
		ev := newv.Index(i)
		d.rv = ev
		if err := arr.Values[i].Accept(d); err != nil {
			// TODO: typed error
			return errors.Wrapf(err, "error decode array element %d", i)
		}
	}
	return nil
}

func (d *Decoder) VisitObject(obj *Object) error {
	// TODO: will this breaks field of type struct
	if d.rv.Kind() != reflect.Ptr {
		return errors.Errorf("reflect value need to be pointer got %s", d.rv.Kind())
	}
	v := d.rv.Elem()
	if v.Kind() != reflect.Struct {
		return errors.Errorf("object decodes to struct but got %s", v.Kind())
	}
	for i, key := range obj.Keys {
		k := key.Val
		fv := v.FieldByName(k)
		if !fv.IsValid() {
			// TODO: typed error
			return errors.Errorf("object field not found in struct %s field %s", v.Type().Name(), k)
		}
		//log.Infof("fv is %s", fv.String())
		d.rv = fv
		if err := obj.Values[i].Accept(d); err != nil {
			return errors.Wrapf(err, "error decode field %s", k)
		}
		//log.Infof("fv is %s", fv.String())
	}
	return nil
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
	return n.Accept(dec)
}
