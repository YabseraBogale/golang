package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
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
