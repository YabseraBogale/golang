package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/YabseraBogale/golang/browser_engine/dom"
)

type Parser struct {
	pos   int
	input string
}

func (p *Parser) Next_Char() rune {
	if p.Eof() {
		return 0
	}
	r, _ := utf8.DecodeLastRuneInString(p.input[p.pos:])
	return r
}

func (p *Parser) Starts_With(s string) bool {
	return len(p.input[p.pos:]) >= len(s) && p.input[p.pos:p.pos+len(s)] == s
}

func (p *Parser) Expect(s string) {
	if p.Starts_With(s) {
		p.pos += len(s)
	} else {
		msg := fmt.Sprintf("Expected %s at byte %d but it was not found", s, p.pos)
		panic(msg)
	}
}

func (p *Parser) Eof() bool {
	return p.pos >= len(p.input)
}

func (p *Parser) Consume_Char() rune {
	c := p.Next_Char()
	p.pos += utf8.RuneLen(c)
	return c
}

func (p *Parser) Consume_While(test func(rune) bool) string {
	var result strings.Builder
	for !p.Eof() && test(p.Next_Char()) {
		result.WriteRune(p.Consume_Char())
	}
	return result.String()
}

func (p *Parser) Consume_Whitespace() {
	p.Consume_While(unicode.IsSpace)
}

func (p *Parser) Parse_Name() string {
	return p.Consume_While(func(c rune) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
	})
}

func (p *Parser) Parse_Attr_Value() string {
	open_quote := p.Consume_Char()
	assert(open_quote == '"' || open_quote == '\'', "failed in parse_attr_value")
	value := p.Consume_While(func(r rune) bool {
		return r != open_quote
	})
	close_quote := p.Consume_Char()
	assert(close_quote == open_quote, "failed in parse_attr_value")
	return value
}

func (p *Parser) Parse_Attr() (name string, value string) {
	name = p.Parse_Name()
	p.Expect("=")
	value = p.Parse_Attr_Value()
	return name, value

}

func (p *Parser) Parse_Attributes() dom.AttrMap {
	attributes := map[string]string{}
	for {
		p.Consume_Whitespace()
		if p.Next_Char() == '<' {
			break
		}
		name, value := p.Parse_Attr()
		attributes[name] = value
	}
	return attributes
}

func (p *Parser) Parse_Element() dom.Node {
	p.Expect("<")
	tag_name := p.Parse_Name()
	attrs := p.Parse_Attributes()
	p.Expect(">")
	childen := p.Parse_Nodes()
	p.Expect("</")
	p.Expect(tag_name)
	p.Expect(">")
	return dom.Elem(tag_name, attrs, childen)
}

func (p *Parser) Parse_Text() dom.Node {
	return dom.Text(p.Consume_While(func(r rune) bool {
		return r != '<'
	}))
}

func (p *Parser) Parse_Node() dom.Node {
	if p.Starts_With("<") {
		return p.Parse_Element()
	} else {
		return p.Parse_Text()
	}
}

func (p *Parser) Parse_Nodes() []dom.Node {
	nodes := []dom.Node{}
	for {
		p.Consume_Whitespace()
		if p.Eof() || p.Starts_With("</") {
			break
		}
		nodes = append(nodes, p.Parse_Node())
	}
	return nodes
}

func assert(condition bool, message string) {
	if !condition {
		msg := fmt.Sprintf("Assertion failed: %s", message)
		panic(msg)
	}
}
