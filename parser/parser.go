package parser

import "fmt"

// NodeType represents the type of a JSON node.
type NodeType int

const (
	NodeObject NodeType = iota
	NodeArray
	NodeString
	NodeNumber
	NodeBool
	NodeNull
)

// Node represents a node in the JSON tree.
type Node struct {
	Type     NodeType
	Key      string
	Children []Node
}

// Parser represents a JSON parser.
type Parser struct {
	tokenizer    *Tokenizer
	currentToken Token
}

// NewParser creates a new parser with the given tokenizer.
func NewParser(tokenizer *Tokenizer) *Parser {
	return &Parser{tokenizer: tokenizer}
}

// Parse parses the JSON input and returns the resulting data structure.
func (p *Parser) Parse() (Node, error) {
	p.nextToken()
	return p.parseValue()
}

// nextToken advances to the next token.
func (p *Parser) nextToken() {
	p.currentToken = p.tokenizer.NextToken()
}

// parseValue parses a JSON value.
func (p *Parser) parseValue() (Node, error) {
	switch p.currentToken.Type {
	case TokenLeftBrace:
		return p.parseObject()
	case TokenString:
		return Node{Type: NodeString}, nil
	case TokenEOF:
		return Node{}, fmt.Errorf("Not a valid JSON input")
	default:
		return Node{}, fmt.Errorf("unexpected token: %v", p.currentToken)
	}
}

// parseObject parses a JSON object.
func (p *Parser) parseObject() (Node, error) {
	node := Node{Type: NodeObject}
	// Consume the '{' token
	p.nextToken()

	for p.currentToken.Type != TokenRightBrace {
		if p.currentToken.Type != TokenString {
			return node, fmt.Errorf("expected string key, got: %v", p.currentToken)
		}
		key := p.currentToken.Value
		node.Key = key

		// Consume the key token
		p.nextToken()

		if p.currentToken.Type != TokenColon {
			return node, fmt.Errorf("expected colon, got: %v", p.currentToken)
		}

		// Consume the ':' token
		p.nextToken()

		value, err := p.parseValue()
		if err != nil {
			return node, err
		}

		node.Children = append(node.Children, value)

		p.nextToken()
		if p.currentToken.Type == TokenComma {
			p.nextToken()
		} else if p.currentToken.Type != TokenRightBrace {
			return node, fmt.Errorf("expected comma or right brace, got: %v", p.currentToken)
		}
	}

	// Consume the '}' token
	p.nextToken()

	return node, nil
}
