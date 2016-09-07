package escape

import (
	"fmt"
	"regexp"
	"strings"
)

var ident = regexp.MustCompile(`(?i)^[a-z_][a-z0-9_$]*$`)
var params = regexp.MustCompile(`%([sIL])`)

// Escape the given `query` with positional `args`.
func Escape(query string, args ...string) (string, error) {
	matches := params.FindAllStringSubmatch(query, -1)

	length := len(matches)
	argc := len(args)

	if argc > length {
		return "", fmt.Errorf("too many arguments for escaped query")
	}

	if argc < length {
		return "", fmt.Errorf("too few arguments for escaped query")
	}

	for i, match := range matches {
		arg := args[i]
		switch match[1] {
		case "L":
			query = strings.Replace(query, "%L", Literal(arg), 1)
		case "I":
			query = strings.Replace(query, "%I", Ident(arg), 1)
		case "s":
			query = strings.Replace(query, "%s", arg, 1)
		}
	}

	return query, nil
}

// Literal escape the given string.
func Literal(s string) string {
	p := ""

	if strings.Contains(s, `\`) {
		p = "E"
	}

	s = strings.Replace(s, `'`, `''`, -1)
	s = strings.Replace(s, `\`, `\\`, -1)
	return p + `'` + s + `'`
}

// Ident quotes the given identifier if required.
func Ident(s string) string {
	if IdentNeedsQuotes(s) {
		return QuoteIdent(s)
	} else {
		return s
	}
}

// IdentNeedsQuotes checks if the given identifier requires quoting
func IdentNeedsQuotes(s string) bool {
	if Reserved[strings.ToUpper(s)] {
		return true
	}

	return !ident.MatchString(s)
}

// QuoteIdent quotes the given identifier string.
func QuoteIdent(s string) string {
	s = strings.Replace(s, `"`, `""`, -1)
	return `"` + s + `"`
}
