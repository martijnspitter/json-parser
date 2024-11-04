package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input    string
		expected Node
	}{
		{
			input: "{}",
			expected: Node{
				Type:     NodeObject,
				Key:      "",
				Children: []Node{},
			},
		},
		{
			input: `{"key": "value"}`,
			expected: Node{
				Type: NodeObject,
				Key:  "",
				Children: []Node{
					{
						Type: NodeString,
						Key:  "key",
						Children: []Node{
							{Type: NodeString, Key: "", Children: nil},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		tokenizer := NewTokenizer(test.input)
		parser := NewParser(tokenizer)
		result, err := parser.Parse()

		if err != nil {
			t.Errorf("Test failed for input '%s'. Unexpected error: %v", test.input, err)
		}

		if !compareNodes(result, test.expected) {
			t.Errorf("Test failed for input '%s'. Expected %v, got %v", test.input, test.expected, result)
		}
	}
}

func compareNodes(a, b Node) bool {
	if a.Type != b.Type || a.Key != b.Key || len(a.Children) != len(b.Children) {
		return false
	}

	for i := range a.Children {
		if !compareNodes(a.Children[i], b.Children[i]) {
			return false
		}
	}

	return true
}
