package parser

import (
	"strings"
	"unicode"
)

// TokenType represents the type of a token.
type TokenType string

// Token types
const (
	TokenLeftBrace    = "TokenLeftBrace"
	TokenRightBrace   = "TokenRightBrace"
	TokenLeftBracket  = "TokenLeftBracket"
	TokenRightBracket = "TokenRightBracket"
	TokenEOF          = "TokenEOF"
	TokenString       = "TokenString"
	TokenColon        = "TokenColon"
	TokenComma        = "TokenComma"
	TokenNumber       = "TokenNumber"
	TokenBool         = "TokenBool"
	TokenNull         = "TokenNull"
)

type Token struct {
	Type  TokenType
	Value string
	Pos   int
}

type Tokenizer struct {
	input string
	pos   int
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{input: cleanInput(input)}
}

func (t *Tokenizer) Tokenize() []Token {
	var tokens []Token
	for token := t.NextToken(); token.Type != TokenEOF; token = t.NextToken() {
		tokens = append(tokens, token)
	}
	return tokens
}

func (t *Tokenizer) NextToken() Token {

	for t.pos < len(t.input) {

		ch := t.input[t.pos]
		tokenPos := t.pos

		t.pos++

		switch ch {
		case '{':
			return Token{Type: TokenLeftBrace, Value: string(ch), Pos: tokenPos}
		case '}':
			return Token{Type: TokenRightBrace, Value: string(ch), Pos: tokenPos}
		case '[':
			return Token{Type: TokenLeftBracket, Value: string(ch), Pos: tokenPos}
		case ']':
			return Token{Type: TokenRightBracket, Value: string(ch), Pos: tokenPos}
		case ':':
			return Token{Type: TokenColon, Value: string(ch), Pos: tokenPos}
		case ',':
			return Token{Type: TokenComma, Value: string(ch), Pos: tokenPos}
		case '"':
			start := t.pos
			for t.pos < len(t.input) && t.input[t.pos] != '"' {
				t.pos++
			}
			if t.pos < len(t.input) {
				value := t.input[start:t.pos]
				t.pos++ // Skip closing quote
				return Token{Type: TokenString, Value: value, Pos: tokenPos}
			}
			return Token{Type: TokenEOF, Value: ""}
		case 't', 'f':
			start := t.pos - 1
			for t.pos < len(t.input) && unicode.IsLetter(rune(t.input[t.pos])) {
				t.pos++
			}
			value := t.input[start:t.pos]
			if value == "true" || value == "false" {
				return Token{Type: TokenBool, Value: value, Pos: tokenPos}
			}
		case 'n':
			start := t.pos - 1
			for t.pos < len(t.input) && unicode.IsLetter(rune(t.input[t.pos])) {
				t.pos++
			}
			value := t.input[start:t.pos]
			if value == "null" {
				return Token{Type: TokenNull, Value: value, Pos: tokenPos}
			}
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			start := t.pos - 1
			for t.pos < len(t.input) && unicode.IsDigit(rune(t.input[t.pos])) {
				t.pos++
			}
			value := t.input[start:t.pos]
			return Token{Type: TokenNumber, Value: value, Pos: tokenPos}
		default:
			if unicode.IsSpace(rune(ch)) {
				continue
			}
			// For now, we'll just return an EOF token for any unrecognized character.
			return Token{Type: TokenEOF, Value: "", Pos: tokenPos}
		}
	}
	return Token{Type: TokenEOF, Value: "", Pos: t.pos}
}

func cleanInput(input string) string {
	cleaned := strings.ReplaceAll(input, "\n", "")
	cleaned = strings.ReplaceAll(cleaned, "\r", "")
	return strings.TrimSpace(cleaned)
}

func (t *Tokenizer) Pos() int {
	return t.pos
}
