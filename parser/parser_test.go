package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input    string
		expected error
	}{
		{
			input:    "{}",
			expected: nil,
		},
		{
			input:    `{"key": "value"}`,
			expected: nil,
		},
		{
			input:    `{"key": "value",}`,
			expected: fmt.Errorf("unexpected end of input"),
		},
		{
			input:    `{"key": "value", "key2": "value2"}`,
			expected: nil,
		},
		{
			input:    `{"key": "value", key2: "value2"}`,
			expected: fmt.Errorf("unexpected token: k at position: 17"),
		},
		{
			input:    `{"key": true, "key2": false, "key3": null, "key4": 123}`,
			expected: nil,
		},
		{
			input:    `{"key": true, "key2": False, "key3": null, "key4": 123}`,
			expected: fmt.Errorf("unexpected token: F at position: 22"),
		},
		{
			input:    `{"key": "value", "key-n": 101, "key-0": {}, "key-l": []}`,
			expected: nil,
		},
		{
			input:    `{"key": "value", "key-n": 101, "key-o": { "inner key": "inner value"}, "key-l": ["list value"]}`,
			expected: nil,
		},
		{
			input:    `{"key": "value", "key-n": 101, "key-o": { "inner key": "inner value"}, "key-l": ['list value']}`,
			expected: fmt.Errorf("unexpected token: ' at position: 81"),
		},
	}

	for _, test := range tests {
		tokenizer := NewTokenizer(test.input)
		parser := NewParser(tokenizer)
		err := parser.Parse()

		if err != nil && test.expected == nil {
			t.Errorf("Test failed for input '%s'. Unexpected error: %v", test.input, err)
		}
		if err == nil && test.expected != nil {
			t.Errorf("Test failed for input '%s'. Expected error: %v", test.input, test.expected)
		}
		if err != nil && test.expected != nil && err.Error() != test.expected.Error() {
			t.Errorf("Test failed for input '%s'. Expected error: %v, got: %v", test.input, test.expected, err)
		}

	}
}
