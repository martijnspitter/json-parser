package parser

import "fmt"

// NodeType represents the type of a JSON node.
type NodeType string

const (
	NodeObject = "NodeObject"
	NodeArray  = "NodeArray"
	NodeString = "NodeString"
	NodeNumber = "NodeNumber"
	NodeBool   = "NodeBool"
	NodeNull   = "NodeNull"
)

// Node represents a node in the JSON tree.
type Node struct {
	Type     NodeType
	Children []*Node
}

// Parser represents a JSON parser.
type Parser struct {
	tokenizer    *Tokenizer
	expectingVal bool
	stack        []*Node
	currentNode  *Node
}

// NewParser creates a new parser with the given tokenizer.
func NewParser(tokenizer *Tokenizer) *Parser {
	return &Parser{
		tokenizer:    tokenizer,
		expectingVal: false,
		stack:        make([]*Node, 0),
	}
}

// Parse parses the JSON input and returns the resulting data structure.
func (p *Parser) Parse() error {
	for _, token := range p.tokenizer.Tokenize() {
		err := p.parseValue(token)
		if err != nil {
			return err
		}
	}

	if p.expectingVal {
		return fmt.Errorf("unexpected end of input")
	}
	return nil
}

// parseValue parses a JSON value.
func (p *Parser) parseValue(token Token) error {
	switch token.Type {
	case TokenLeftBrace:
		// Start of a new JSON object
		// Create new Node
		newNode := getNewNode(NodeObject)
		if p.currentNode != nil {
			// Add new Node to the current Node's children
			p.currentNode.Children = append(p.currentNode.Children, newNode)
		}
		// Push the current Node to the stack
		p.stack = append(p.stack, p.currentNode)
		// Set the new Node as the current Node
		p.currentNode = newNode
	case TokenRightBrace:
		if len(p.stack) == 0 {
			return fmt.Errorf("unexpected token: %v at position: %d", token.Value, token.Pos)
		}
		// End of the current JSON object
		// Set current Node to the last Node in the stack (parent Node)
		p.currentNode = p.stack[len(p.stack)-1]
		// Pop the last Node from the stack
		p.stack = p.stack[:len(p.stack)-1]
	case TokenLeftBracket:
		// Start of a new JSON array
		// Create new Node
		newNode := getNewNode(NodeArray)
		// Add new Node to the current Node's children
		p.currentNode.Children = append(p.currentNode.Children, newNode)
		// Push the current Node to the stack
		p.stack = append(p.stack, p.currentNode)
		// Set the new Node as the current Node
		p.currentNode = newNode
	case TokenRightBracket:
		if len(p.stack) == 0 {
			return fmt.Errorf("unexpected token: %v at position: %d", token.Value, token.Pos)
		}
		// End of the current JSON array
		// Set current Node to the last Node in the stack (parent Node)
		p.currentNode = p.stack[len(p.stack)-1]
		// Pop the last Node from the stack
		p.stack = p.stack[:len(p.stack)-1]
		p.expectingVal = false
	case TokenColon:
		// Expecting a value after a key
		p.expectingVal = true
	case TokenComma:
		// Expecting a key after a value
		p.expectingVal = true
	case TokenString, TokenNumber, TokenBool, TokenNull:
		// Handle string value
		// Check if the Current Node has Children and the Last Child is an Object
		if len(p.currentNode.Children) > 0 && p.currentNode.Children[len(p.currentNode.Children)-1].Type == NodeObject {
			// Add the Token as a Child of the Last Object Node
			p.currentNode.Children[len(p.currentNode.Children)-1].Children = append(p.currentNode.Children[len(p.currentNode.Children)-1].Children, getNewNode(NodeString))
		} else {
			// Add the Token as a Child of the Current Node
			p.currentNode.Children = append(p.currentNode.Children, getNewNode(NodeString))
		}
		p.expectingVal = false
	case TokenEOF:
		return fmt.Errorf("Not a valid JSON input")
	default:
		return fmt.Errorf("unexpected token: %v at position: %d", token.Value, token.Pos)
	}

	return nil
}

func getNewNode(nodeType NodeType) *Node {
	return &Node{
		Type:     nodeType,
		Children: make([]*Node, 0),
	}
}
