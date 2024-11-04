package parser

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			input: "{}",
			expected: []Token{
				{Type: TokenLeftBrace, Value: "{"},
				{Type: TokenRightBrace, Value: "}"},
				{Type: TokenEOF},
			},
		},
		{
			input: " { } ",
			expected: []Token{
				{Type: TokenLeftBrace, Value: "{"},
				{Type: TokenRightBrace, Value: "}"},
				{Type: TokenEOF},
			},
		},
		{
			input: "{ }",
			expected: []Token{
				{Type: TokenLeftBrace, Value: "{"},
				{Type: TokenRightBrace, Value: "}"},
				{Type: TokenEOF},
			},
		},
		{
			input: "{ \"key\": \"value\" }",
			expected: []Token{
				{Type: TokenLeftBrace, Value: "{"},
				{Type: TokenString, Value: "key"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenString, Value: "value"},
				{Type: TokenRightBrace, Value: "}"},
				{Type: TokenEOF},
			},
		},
		{
			input: "{ \"key\": \"value\", }",
			expected: []Token{
				{Type: TokenLeftBrace, Value: "{"},
				{Type: TokenString, Value: "key"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenString, Value: "value"},
				{Type: TokenComma, Value: ","},
				{Type: TokenRightBrace, Value: "}"},
				{Type: TokenEOF},
			},
		},
	}

	for _, test := range tests {
		tokenizer := NewTokenizer(test.input)
		for i, expectedToken := range test.expected {
			token := tokenizer.NextToken()

			if token.Type != expectedToken.Type || token.Value != expectedToken.Value {
				t.Errorf("Test failed for input '%s'. Expected token %d to be %v, got %v", test.input, i, expectedToken, token)
			}
		}
	}
}
