package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Rule represents a CSS rule with a selector and declarations.
type Rule struct {
	Selector     string
	Declarations map[string]string
}

// ParseCSS parses a simple CSS string and returns a slice of rules.
func ParseCSS(css string) []Rule {
	var rules []Rule
	css = strings.TrimSpace(css)

	// Regular expression to match CSS rules
	ruleRegex := regexp.MustCompile(`(?s)([^{}]+)\s*{\s*([^}]*)\s*}`)
	matches := ruleRegex.FindAllStringSubmatch(css, -1)

	for _, match := range matches {
		selector := strings.TrimSpace(match[1])
		declarations := parseDeclarations(match[2])

		rules = append(rules, Rule{Selector: selector, Declarations: declarations})
	}

	return rules
}

// parseDeclarations parses CSS declarations and returns them as a map.
func parseDeclarations(decl string) map[string]string {
	declarations := make(map[string]string)
	lines := strings.Split(decl, ";")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			declarations[key] = value
		}
	}

	return declarations
}

func main() {
	css := `
		body {
			background-color: #fff;
			color: #333;
		}
		p {
			font-size: 16px;
			margin: 10px;
		}
	`

	rules := ParseCSS(css)
	for _, rule := range rules {
		fmt.Println("Selector:", rule.Selector)
		for prop, value := range rule.Declarations {
			fmt.Printf("  %s: %s\n", prop, value)
		}
	}
}
