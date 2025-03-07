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

func (p *Parser) next_char() rune {
	if p.eof() {
		return 0
	}
	r, _ := utf8.DecodeLastRuneInString(p.input[p.pos:])
	return r
}

func (p *Parser) starts_with(s string) bool {
	return len(p.input[p.pos:]) >= len(s) && p.input[p.pos:p.pos+len(s)] == s
}

func (p *Parser) expect(s string) {
	if p.starts_with(s) {
		p.pos += len(s)
	} else {
		msg := fmt.Sprintf("Expected %s at byte %d but it was not found", s, p.pos)
		panic(msg)
	}
}

func (p *Parser) eof() bool {
	return p.pos >= len(p.input)
}

func (p *Parser) consume_char() rune {
	c := p.next_char()
	p.pos += utf8.RuneLen(c)
	return c
}

func (p *Parser) consume_while(test func(rune) bool) string {
	var result strings.Builder
	for !p.eof() && test(p.next_char()) {
		result.WriteRune(p.consume_char())
	}
	return result.String()
}

func (p *Parser) consume_whitespace() {
	p.consume_while(unicode.IsSpace)
}

func (p *Parser) parse_name() string {
	return p.consume_while(func(c rune) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
	})
}

func (p *Parser) parse_attr_value() string {
	open_quote := p.consume_char()
	assert(open_quote == '"' || open_quote == '\'', "failed in parse_attr_value")
	value := p.consume_while(func(r rune) bool {
		return r != open_quote
	})
	close_quote := p.consume_char()
	assert(close_quote == open_quote, "failed in parse_attr_value")
	return value
}

func (p *Parser) parse_attr() (name string, value string) {
	name = p.parse_name()
	p.expect("=")
	value = p.parse_attr_value()
	return name, value

}

func (p *Parser) parse_attributes() dom.AttrMap {
	attributes := map[string]string{}
	for {
		p.consume_whitespace()
		if p.next_char() == '<' {
			break
		}
		name, value := p.parse_attr()
		attributes[name] = value
	}
	return attributes
}

func (p *Parser) parse_element() dom.Node {
	p.expect("<")
	tag_name := p.parse_name()
	attrs := p.parse_attributes()
	p.expect(">")
	childen := p.parse_nodes()
	p.expect("</")
	p.expect(tag_name)
	p.expect(">")
	return dom.Elem(tag_name, attrs, childen)
}

func (p *Parser) parse_text() dom.Node {
	return dom.Text(p.consume_while(func(r rune) bool {
		return r != '<'
	}))
}

func (p *Parser) parse_node() dom.Node {
	if p.starts_with("<") {
		return p.parse_element()
	} else {
		return p.parse_text()
	}
}

func (p *Parser) parse_nodes() []dom.Node {
	nodes := []dom.Node{}
	for {
		p.consume_whitespace()
		if p.eof() || p.starts_with("</") {
			break
		}
		nodes = append(nodes, p.parse_node())
	}
	return nodes
}

func assert(condition bool, message string) {
	if !condition {
		msg := fmt.Sprintf("Assertion failed: %s", message)
		panic(msg)
	}
}
