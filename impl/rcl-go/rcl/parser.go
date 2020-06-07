package rcl

import (
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
	pos   int
	width int // width of last character, used for backup
	line  int // 1+number of newlines seen
}

func NewParser(input string) *Parser {
	return &Parser{input: input, line: 1}
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
	if strings.HasPrefix(p.input[p.pos:], "null") {
		p.pos += 4
		return nil, nil
	}
	return &Null{}, errors.New("invalid null")
}

func (p *Parser) parseBool() (*Bool, error) {
	if strings.HasPrefix(p.input[p.pos:], "true") {
		p.pos += 4
		return &Bool{Val: true}, nil
	}
	if strings.HasPrefix(p.input[p.pos:], "false") {
		p.pos += 4
		return &Bool{Val: false}, nil
	}
	return nil, errors.New("invalid bool")
}

func (p *Parser) parseString() (*String, error) {
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
	return &String{Val: sb.String()}, nil
}

func (p *Parser) parseNumber() (*Number, error) {
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

	if float {
		n, err := strconv.ParseFloat(string(v), 64)
		// TODO: wrap error and contains position information
		if err != nil {
			return nil, err
		}
		return &Number{
			Val:  int64(math.Float64bits(n)),
			Type: NumberTypeDouble,
		}, nil
	} else {
		n, err := strconv.ParseInt(string(v), 10, 64)
		// TODO: wrap error and contains position information
		if err != nil {
			return nil, err
		}
		return &Number{
			Val:  n,
			Type: NumberTypeInt,
		}, nil
	}
}

func (p *Parser) parseArray() (*Array, error) {
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
	return &Array{Values: values}, nil
}

func (p *Parser) parseObject() (*Object, error) {
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
	return &Object{Keys: keys, Values: values}, nil
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
