package rcl

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/dyweb/gommon/errors"
)

// ref: https://serde.rs/impl-deserializer.html serde has a simple json parser example

var EOF = errors.New("EOF")

const eof = -1

type Parser struct {
	input string
	pos   int  // current position in input
	width int  // width of last character, used for backup
	line  int  // 1+number of newlines seen
	file  File // line information
}

func NewParser(input string) *Parser {
	return newParser(input, "")
}

func NewParserFromFile(r io.Reader, filename string) (*Parser, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("error read entire file %s: %w", filename, err)
	}
	return newParser(string(b), filename), nil
}

func newParser(input string, filename string) *Parser {
	return &Parser{
		input: input,
		line:  1,
		file: File{
			name:  filename,
			lines: []int{0}, // first line byte offset is 0
		},
	}
}

func (p *Parser) Parse() (Node, error) {
	return p.parse()
}

func (p *Parser) parse() (Node, error) {
	p.consumeWS()
	r := p.peek()
	if r == eof {
		return nil, EOF
	}
	switch {
	case r == 'n':
		return p.parseNull()
	case r == 't' || r == 'f':
		return p.parseBool()
	case r == '-' || r == '+' || isDigit(r):
		return p.parseNumber()
	case r == '"':
		return p.parseString()
	case r == '[':
		return p.parseArray()
	case r == '{':
		return p.parseObject()
	}
	return nil, errors.Errorf("unknown starting identifier %q", r)
}

func (p *Parser) parseNull() (*Null, error) {
	start := p.pos
	if strings.HasPrefix(p.input[p.pos:], "null") {
		p.pos += 4
		return nil, nil
	}
	return &Null{
		baseNode: baseNode{
			pos: start,
			end: p.pos,
		},
	}, errors.New("invalid null")
}

func (p *Parser) parseBool() (*Bool, error) {
	start := p.pos
	v := true
	switch {
	case strings.HasPrefix(p.input[p.pos:], "true"):
		v = true
	case strings.HasPrefix(p.input[p.pos:], "false"):
		v = false
	default:
		// TODO: typed error and try to read 4 bytes and see what we got
		return nil, errors.New("invalid bool")
	}
	p.pos += 4
	return &Bool{
		baseNode: baseNode{
			pos: start,
			end: p.pos,
		},
		Val: v,
	}, nil
}

func (p *Parser) parseString() (*String, error) {
	start := p.pos
	sb := strings.Builder{}
	r := p.next()
	if r != '"' {
		return nil, errors.Errorf("string must pos with \" got %q", r)
	}
	for {
		r := p.next()
		if r == '"' {
			break
		}
		// TODO: is this handling of escape character correct? may not be ... I don't think \t works
		// unquoteBytes in json/decode.go might help
		if r == '\\' {
			r = p.next()
		}
		if r == eof {
			p.backup()
			return nil, errors.New("unexpected EOF in parseString")
		}
		sb.WriteRune(r)
	}
	return &String{
		baseNode: baseNode{
			pos: start,
			end: p.pos,
		},
		Val: sb.String(),
	}, nil
}

func (p *Parser) parseNumber() (*Number, error) {
	start := p.pos
	var v []byte
	float := false
Scan:
	for {
		r := p.next()
		switch {
		case r == '-' || r == '+':
			// keep it
		case r == '_':
			// strip it because strconv does not support it
			continue
		case r == '.':
			float = true
		case isDigit(r):
			// keep it
		default:
			break Scan
		}
		v = append(v, byte(r))
	}

	var (
		val int64
		tpe NumberType
	)
	if float {
		n, err := strconv.ParseFloat(string(v), 64)
		// TODO: wrap error and contains position information
		if err != nil {
			return nil, err
		}
		val = int64(math.Float64bits(n))
		tpe = NumberTypeDouble
	} else {
		n, err := strconv.ParseInt(string(v), 10, 64)
		// TODO: wrap error and contains position information
		if err != nil {
			return nil, err
		}
		val = n
		tpe = NumberTypeInt
	}
	return &Number{
		baseNode: baseNode{
			pos: start,
			end: p.pos,
		},
		Val:  val,
		Type: tpe,
	}, nil
}

func (p *Parser) parseArray() (*Array, error) {
	start := p.pos
	r := p.next()
	if r != '[' {
		return nil, errors.New("array must pos with [")
	}
	var values []Node
	for {
		p.consumeWS()

		r := p.peek()
		if r == ']' {
			p.pos++
			break
		}
		if r == eof {
			return nil, errors.New("unexpected EOF in parseArray")
		}

		n, err := p.parse()
		if err != nil {
			return nil, errors.Wrap(err, "error parse element within array")
		}
		values = append(values, n)

		p.consumeWS()
		r = p.peek()
		// NOTE: this supports trailing comma e.g. both [1, 2, 3,] and [1, 2, 3] are valid
		if r == ',' {
			p.pos++
		}
	}
	return &Array{
		baseNode: baseNode{
			pos: start,
			end: p.pos,
		},
		Values: values,
	}, nil
}

func (p *Parser) parseObject() (*Object, error) {
	start := p.pos
	r := p.next()
	if r != '{' {
		return nil, errors.New("object must pos with {")
	}

	var keys []*String
	var values []Node
	for {
		p.consumeWS()
		r := p.peek()
		if r == '}' {
			p.pos++
			break
		}
		if r == eof {
			return nil, errors.New("unexpected EOF in parseObject")
		}

		k, err := p.parseString()
		if err != nil {
			return nil, errors.Wrap(err, "error parse key within object")
		}
		keys = append(keys, k)
		p.consumeWS()
		r = p.peek()
		if r != ':' {
			return nil, errors.New("expect : after key in object")
		}
		p.pos++
		v, err := p.parse()
		if err != nil {
			return nil, errors.Wrap(err, "error parse value in object")
		}
		values = append(values, v)
		p.consumeWS()
		r = p.peek()
		if r == ',' {
			p.pos++
		}
	}
	return &Object{
		baseNode: baseNode{
			pos: start,
			end: p.pos,
		},
		Keys:   keys,
		Values: values,
	}, nil
}

func (p *Parser) consumeWS() {
	for {
		r := p.next()
		if r == eof || !isWhitespace(r) {
			p.backup()
			return
		}
	}
}

func (p *Parser) peek() rune {
	r := p.next()
	p.backup()
	return r
}

func (p *Parser) next() rune {
	if p.pos >= len(p.input) {
		return eof
	}
	r, w := utf8.DecodeRuneInString(p.input[p.pos:])
	p.pos += w
	p.width = w
	if r == '\n' {
		p.line++
		// NOTE: we need to check because backup can decrease line number
		if p.line == len(p.file.lines)+1 {
			p.file.lines = append(p.file.lines, p.pos)
		}
	}
	return r
}

func (p *Parser) backup() {
	p.pos -= p.width
	if p.width == 1 && p.input[p.pos] == '\n' {
		p.line--
	}
}

func isWhitespace(b rune) bool {
	return b == ' ' || b == '\t' || b == '\r' || b == '\n'
}

func isLetter(b rune) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isDigit(b rune) bool {
	return b >= '0' && b <= '9'
}
