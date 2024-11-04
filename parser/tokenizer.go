package parser

import (
	"strings"
	"unicode"
)

// TokenType represents the type of a token.
type TokenType int

// Token types
const (
	TokenLeftBrace TokenType = iota
	TokenRightBrace
	TokenEOF
	TokenString
	TokenColon
	TokenComma
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	input string
	pos   int
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{input: cleanInput(input)}
}

func (t *Tokenizer) NextToken() Token {

	for t.pos < len(t.input) {

		ch := t.input[t.pos]
		t.pos++

		switch ch {
		case '{':
			return Token{Type: TokenLeftBrace, Value: string(ch)}
		case '}':
			return Token{Type: TokenRightBrace, Value: string(ch)}
		case ':':
			return Token{Type: TokenColon, Value: string(ch)}
		case ',':
			return Token{Type: TokenComma, Value: string(ch)}
		case '"':
			start := t.pos
			for t.pos < len(t.input) && t.input[t.pos] != '"' {
				t.pos++
			}
			if t.pos < len(t.input) {
				value := t.input[start:t.pos]
				t.pos++ // Skip closing quote
				return Token{Type: TokenString, Value: value}
			}
			return Token{Type: TokenEOF, Value: ""}
		default:
			if unicode.IsSpace(rune(ch)) {
				continue
			}
			// For now, we'll just return an EOF token for any unrecognized character.
			return Token{Type: TokenEOF, Value: ""}
		}
	}
	return Token{Type: TokenEOF, Value: ""}
}

func cleanInput(input string) string {
	cleaned := strings.ReplaceAll(input, "\n", "")
	cleaned = strings.ReplaceAll(cleaned, "\r", "")
	return strings.TrimSpace(cleaned)
}
